// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package context defines a type that carries context specific data such as the logger.
// Inspired by Google's http://godoc.org/golang.org/x/net/context
package context

import (
	"github.com/aws/amazon-ssm-agent/agent/appconfig"
	"github.com/aws/amazon-ssm-agent/agent/log"
)

// T transfers context specific data across different execution boundaries.
// Instead of adding the context to specific structs, we pass Context as the first
// parameter to the methods themselves.
type T interface {
	Log() log.T
	AppConfig() appconfig.SsmagentConfig
	With(context string) T
}

// Default returns an empty context that use the default logger and appconfig.
func Default(logger log.T, appconfig appconfig.SsmagentConfig) T {
	ctx := &defaultContext{log: logger, appconfig: appconfig}
	return ctx
}

type defaultContext struct {
	context   []string
	log       log.T
	appconfig appconfig.SsmagentConfig
}

func (c *defaultContext) With(logContext string) T {
	contextSlice := append(c.context, logContext)
	newContext := &defaultContext{
		context:   contextSlice,
		log:       log.WithContext(contextSlice...),
		appconfig: c.appconfig,
	}
	return newContext
}

func (c *defaultContext) Log() log.T {
	return c.log
}

func (c *defaultContext) AppConfig() appconfig.SsmagentConfig {
	return c.appconfig
}
