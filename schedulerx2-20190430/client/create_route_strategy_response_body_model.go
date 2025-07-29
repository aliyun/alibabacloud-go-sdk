// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateRouteStrategyResponseBody
	GetCode() *int32
	SetData(v *CreateRouteStrategyResponseBodyData) *CreateRouteStrategyResponseBody
	GetData() *CreateRouteStrategyResponseBodyData
	SetMessage(v string) *CreateRouteStrategyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateRouteStrategyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateRouteStrategyResponseBody
	GetSuccess() *bool
}

type CreateRouteStrategyResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateRouteStrategyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The additional information, including errors and tips.
	//
	// example:
	//
	// strategy name is null or empty.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRouteStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRouteStrategyResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateRouteStrategyResponseBody) GetData() *CreateRouteStrategyResponseBodyData {
	return s.Data
}

func (s *CreateRouteStrategyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateRouteStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRouteStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateRouteStrategyResponseBody) SetCode(v int32) *CreateRouteStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRouteStrategyResponseBody) SetData(v *CreateRouteStrategyResponseBodyData) *CreateRouteStrategyResponseBody {
	s.Data = v
	return s
}

func (s *CreateRouteStrategyResponseBody) SetMessage(v string) *CreateRouteStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRouteStrategyResponseBody) SetRequestId(v string) *CreateRouteStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRouteStrategyResponseBody) SetSuccess(v bool) *CreateRouteStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRouteStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateRouteStrategyResponseBodyData struct {
}

func (s CreateRouteStrategyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteStrategyResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateRouteStrategyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
