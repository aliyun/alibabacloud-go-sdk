// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOutboundCallRestrictionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteOutboundCallRestrictionResponseBody
	GetCode() *string
	SetData(v string) *DeleteOutboundCallRestrictionResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteOutboundCallRestrictionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteOutboundCallRestrictionResponseBody
	GetMessage() *string
	SetParams(v []*string) *DeleteOutboundCallRestrictionResponseBody
	GetParams() []*string
	SetRequestId(v string) *DeleteOutboundCallRestrictionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteOutboundCallRestrictionResponseBody
	GetSuccess() *bool
}

type DeleteOutboundCallRestrictionResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	//
	// example:
	//
	// xxx
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Instance does not exist. Instance=xxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteOutboundCallRestrictionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOutboundCallRestrictionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOutboundCallRestrictionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteOutboundCallRestrictionResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteOutboundCallRestrictionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteOutboundCallRestrictionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteOutboundCallRestrictionResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DeleteOutboundCallRestrictionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOutboundCallRestrictionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteOutboundCallRestrictionResponseBody) SetCode(v string) *DeleteOutboundCallRestrictionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteOutboundCallRestrictionResponseBody) SetData(v string) *DeleteOutboundCallRestrictionResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteOutboundCallRestrictionResponseBody) SetHttpStatusCode(v int32) *DeleteOutboundCallRestrictionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteOutboundCallRestrictionResponseBody) SetMessage(v string) *DeleteOutboundCallRestrictionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteOutboundCallRestrictionResponseBody) SetParams(v []*string) *DeleteOutboundCallRestrictionResponseBody {
	s.Params = v
	return s
}

func (s *DeleteOutboundCallRestrictionResponseBody) SetRequestId(v string) *DeleteOutboundCallRestrictionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOutboundCallRestrictionResponseBody) SetSuccess(v bool) *DeleteOutboundCallRestrictionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteOutboundCallRestrictionResponseBody) Validate() error {
	return dara.Validate(s)
}
