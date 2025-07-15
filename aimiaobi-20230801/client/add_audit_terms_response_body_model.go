// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAuditTermsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddAuditTermsResponseBody
	GetCode() *string
	SetData(v bool) *AddAuditTermsResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *AddAuditTermsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddAuditTermsResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddAuditTermsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddAuditTermsResponseBody
	GetSuccess() *bool
}

type AddAuditTermsResponseBody struct {
	// example:
	//
	// DataNotExists
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
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F2F366D6-E9FE-1006-BB70-2C650896AAB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddAuditTermsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAuditTermsResponseBody) GoString() string {
	return s.String()
}

func (s *AddAuditTermsResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddAuditTermsResponseBody) GetData() *bool {
	return s.Data
}

func (s *AddAuditTermsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddAuditTermsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddAuditTermsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAuditTermsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddAuditTermsResponseBody) SetCode(v string) *AddAuditTermsResponseBody {
	s.Code = &v
	return s
}

func (s *AddAuditTermsResponseBody) SetData(v bool) *AddAuditTermsResponseBody {
	s.Data = &v
	return s
}

func (s *AddAuditTermsResponseBody) SetHttpStatusCode(v int32) *AddAuditTermsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddAuditTermsResponseBody) SetMessage(v string) *AddAuditTermsResponseBody {
	s.Message = &v
	return s
}

func (s *AddAuditTermsResponseBody) SetRequestId(v string) *AddAuditTermsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAuditTermsResponseBody) SetSuccess(v bool) *AddAuditTermsResponseBody {
	s.Success = &v
	return s
}

func (s *AddAuditTermsResponseBody) Validate() error {
	return dara.Validate(s)
}
