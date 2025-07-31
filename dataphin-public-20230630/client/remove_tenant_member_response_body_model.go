// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTenantMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveTenantMemberResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RemoveTenantMemberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RemoveTenantMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveTenantMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveTenantMemberResponseBody
	GetSuccess() *bool
}

type RemoveTenantMemberResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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

func (s RemoveTenantMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveTenantMemberResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTenantMemberResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveTenantMemberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemoveTenantMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveTenantMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveTenantMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveTenantMemberResponseBody) SetCode(v string) *RemoveTenantMemberResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveTenantMemberResponseBody) SetHttpStatusCode(v int32) *RemoveTenantMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveTenantMemberResponseBody) SetMessage(v string) *RemoveTenantMemberResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveTenantMemberResponseBody) SetRequestId(v string) *RemoveTenantMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveTenantMemberResponseBody) SetSuccess(v bool) *RemoveTenantMemberResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveTenantMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
