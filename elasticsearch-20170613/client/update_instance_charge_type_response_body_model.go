// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceChargeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateInstanceChargeTypeResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateInstanceChargeTypeResponseBody
	GetResult() *bool
}

type UpdateInstanceChargeTypeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return results:
	//
	// 	- true: conversion successful
	//
	// 	- false: conversion failed
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateInstanceChargeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceChargeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceChargeTypeResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateInstanceChargeTypeResponseBody) SetRequestId(v string) *UpdateInstanceChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceChargeTypeResponseBody) SetResult(v bool) *UpdateInstanceChargeTypeResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateInstanceChargeTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
