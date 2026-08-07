// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOutboundCallRestrictionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateOutboundCallRestrictionResponseBody
	GetCode() *string
	SetData(v string) *CreateOutboundCallRestrictionResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateOutboundCallRestrictionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateOutboundCallRestrictionResponseBody
	GetMessage() *string
	SetParams(v []*string) *CreateOutboundCallRestrictionResponseBody
	GetParams() []*string
	SetRequestId(v string) *CreateOutboundCallRestrictionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateOutboundCallRestrictionResponseBody
	GetSuccess() *bool
}

type CreateOutboundCallRestrictionResponseBody struct {
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
	// Sample value
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
	// Instance does not exist. Instance=xxxxx
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

func (s CreateOutboundCallRestrictionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOutboundCallRestrictionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOutboundCallRestrictionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateOutboundCallRestrictionResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateOutboundCallRestrictionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateOutboundCallRestrictionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateOutboundCallRestrictionResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CreateOutboundCallRestrictionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOutboundCallRestrictionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateOutboundCallRestrictionResponseBody) SetCode(v string) *CreateOutboundCallRestrictionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateOutboundCallRestrictionResponseBody) SetData(v string) *CreateOutboundCallRestrictionResponseBody {
	s.Data = &v
	return s
}

func (s *CreateOutboundCallRestrictionResponseBody) SetHttpStatusCode(v int32) *CreateOutboundCallRestrictionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateOutboundCallRestrictionResponseBody) SetMessage(v string) *CreateOutboundCallRestrictionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOutboundCallRestrictionResponseBody) SetParams(v []*string) *CreateOutboundCallRestrictionResponseBody {
	s.Params = v
	return s
}

func (s *CreateOutboundCallRestrictionResponseBody) SetRequestId(v string) *CreateOutboundCallRestrictionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOutboundCallRestrictionResponseBody) SetSuccess(v bool) *CreateOutboundCallRestrictionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateOutboundCallRestrictionResponseBody) Validate() error {
	return dara.Validate(s)
}
