// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCallCenterProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCallCenterProviderResponseBody
	GetCode() *string
	SetData(v string) *CreateCallCenterProviderResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateCallCenterProviderResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateCallCenterProviderResponseBody
	GetMessage() *string
	SetParams(v []*string) *CreateCallCenterProviderResponseBody
	GetParams() []*string
	SetRequestId(v string) *CreateCallCenterProviderResponseBody
	GetRequestId() *string
}

type CreateCallCenterProviderResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the human agent service provider.
	//
	// example:
	//
	// xxxxxxxxx
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
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The dynamic error parameters.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCallCenterProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCallCenterProviderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCallCenterProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCallCenterProviderResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateCallCenterProviderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateCallCenterProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCallCenterProviderResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CreateCallCenterProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCallCenterProviderResponseBody) SetCode(v string) *CreateCallCenterProviderResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCallCenterProviderResponseBody) SetData(v string) *CreateCallCenterProviderResponseBody {
	s.Data = &v
	return s
}

func (s *CreateCallCenterProviderResponseBody) SetHttpStatusCode(v int32) *CreateCallCenterProviderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateCallCenterProviderResponseBody) SetMessage(v string) *CreateCallCenterProviderResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCallCenterProviderResponseBody) SetParams(v []*string) *CreateCallCenterProviderResponseBody {
	s.Params = v
	return s
}

func (s *CreateCallCenterProviderResponseBody) SetRequestId(v string) *CreateCallCenterProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCallCenterProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
