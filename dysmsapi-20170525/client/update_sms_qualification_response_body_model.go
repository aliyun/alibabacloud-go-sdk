// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmsQualificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateSmsQualificationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateSmsQualificationResponseBody
	GetCode() *string
	SetData(v string) *UpdateSmsQualificationResponseBody
	GetData() *string
	SetMessage(v string) *UpdateSmsQualificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSmsQualificationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSmsQualificationResponseBody
	GetSuccess() *bool
}

type UpdateSmsQualificationResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 10000****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 25D5AFDE-8EBC-132E-8909-1FDC071DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSmsQualificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmsQualificationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSmsQualificationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateSmsQualificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSmsQualificationResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateSmsQualificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSmsQualificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSmsQualificationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSmsQualificationResponseBody) SetAccessDeniedDetail(v string) *UpdateSmsQualificationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateSmsQualificationResponseBody) SetCode(v string) *UpdateSmsQualificationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSmsQualificationResponseBody) SetData(v string) *UpdateSmsQualificationResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateSmsQualificationResponseBody) SetMessage(v string) *UpdateSmsQualificationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSmsQualificationResponseBody) SetRequestId(v string) *UpdateSmsQualificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSmsQualificationResponseBody) SetSuccess(v bool) *UpdateSmsQualificationResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSmsQualificationResponseBody) Validate() error {
	return dara.Validate(s)
}
