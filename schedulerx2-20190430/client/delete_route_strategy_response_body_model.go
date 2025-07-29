// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteRouteStrategyResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteRouteStrategyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteRouteStrategyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteRouteStrategyResponseBody
	GetSuccess() *bool
}

type DeleteRouteStrategyResponseBody struct {
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
	// strategy is already deleted.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 71BCC0E3-64B2-4B63-A870-AFB64EBCB5A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRouteStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRouteStrategyResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteRouteStrategyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteRouteStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRouteStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRouteStrategyResponseBody) SetCode(v int32) *DeleteRouteStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRouteStrategyResponseBody) SetMessage(v string) *DeleteRouteStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRouteStrategyResponseBody) SetRequestId(v string) *DeleteRouteStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRouteStrategyResponseBody) SetSuccess(v bool) *DeleteRouteStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRouteStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
