// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCallCenterProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateCallCenterProviderResponseBody
	GetCode() *string
	SetData(v string) *UpdateCallCenterProviderResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateCallCenterProviderResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateCallCenterProviderResponseBody
	GetMessage() *string
	SetParams(v []*string) *UpdateCallCenterProviderResponseBody
	GetParams() []*string
	SetRequestId(v string) *UpdateCallCenterProviderResponseBody
	GetRequestId() *string
}

type UpdateCallCenterProviderResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// xxxxxxxxx
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCallCenterProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCallCenterProviderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCallCenterProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateCallCenterProviderResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateCallCenterProviderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateCallCenterProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCallCenterProviderResponseBody) GetParams() []*string {
	return s.Params
}

func (s *UpdateCallCenterProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCallCenterProviderResponseBody) SetCode(v string) *UpdateCallCenterProviderResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCallCenterProviderResponseBody) SetData(v string) *UpdateCallCenterProviderResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateCallCenterProviderResponseBody) SetHttpStatusCode(v int32) *UpdateCallCenterProviderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateCallCenterProviderResponseBody) SetMessage(v string) *UpdateCallCenterProviderResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCallCenterProviderResponseBody) SetParams(v []*string) *UpdateCallCenterProviderResponseBody {
	s.Params = v
	return s
}

func (s *UpdateCallCenterProviderResponseBody) SetRequestId(v string) *UpdateCallCenterProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCallCenterProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
