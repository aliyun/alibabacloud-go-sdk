// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTenantMembersBySourceUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddTenantMembersBySourceUserResponseBody
	GetCode() *string
	SetData(v bool) *AddTenantMembersBySourceUserResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *AddTenantMembersBySourceUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddTenantMembersBySourceUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddTenantMembersBySourceUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddTenantMembersBySourceUserResponseBody
	GetSuccess() *bool
}

type AddTenantMembersBySourceUserResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddTenantMembersBySourceUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTenantMembersBySourceUserResponseBody) GoString() string {
	return s.String()
}

func (s *AddTenantMembersBySourceUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddTenantMembersBySourceUserResponseBody) GetData() *bool {
	return s.Data
}

func (s *AddTenantMembersBySourceUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddTenantMembersBySourceUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddTenantMembersBySourceUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTenantMembersBySourceUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddTenantMembersBySourceUserResponseBody) SetCode(v string) *AddTenantMembersBySourceUserResponseBody {
	s.Code = &v
	return s
}

func (s *AddTenantMembersBySourceUserResponseBody) SetData(v bool) *AddTenantMembersBySourceUserResponseBody {
	s.Data = &v
	return s
}

func (s *AddTenantMembersBySourceUserResponseBody) SetHttpStatusCode(v int32) *AddTenantMembersBySourceUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddTenantMembersBySourceUserResponseBody) SetMessage(v string) *AddTenantMembersBySourceUserResponseBody {
	s.Message = &v
	return s
}

func (s *AddTenantMembersBySourceUserResponseBody) SetRequestId(v string) *AddTenantMembersBySourceUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTenantMembersBySourceUserResponseBody) SetSuccess(v bool) *AddTenantMembersBySourceUserResponseBody {
	s.Success = &v
	return s
}

func (s *AddTenantMembersBySourceUserResponseBody) Validate() error {
	return dara.Validate(s)
}
