// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteRouteStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *BatchDeleteRouteStrategyResponseBody
	GetCode() *int32
	SetMessage(v string) *BatchDeleteRouteStrategyResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchDeleteRouteStrategyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchDeleteRouteStrategyResponseBody
	GetSuccess() *bool
}

type BatchDeleteRouteStrategyResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// job is not existed, jobId=162837
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 704A2A61-3681-5568-92F7-2DFCC53F33D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// true: The request was successful.
	//
	// false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchDeleteRouteStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteRouteStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteRouteStrategyResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *BatchDeleteRouteStrategyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchDeleteRouteStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteRouteStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchDeleteRouteStrategyResponseBody) SetCode(v int32) *BatchDeleteRouteStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *BatchDeleteRouteStrategyResponseBody) SetMessage(v string) *BatchDeleteRouteStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *BatchDeleteRouteStrategyResponseBody) SetRequestId(v string) *BatchDeleteRouteStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteRouteStrategyResponseBody) SetSuccess(v bool) *BatchDeleteRouteStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *BatchDeleteRouteStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
