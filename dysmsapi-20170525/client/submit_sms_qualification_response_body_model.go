// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmsQualificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SubmitSmsQualificationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *SubmitSmsQualificationResponseBody
	GetCode() *string
	SetData(v string) *SubmitSmsQualificationResponseBody
	GetData() *string
	SetMessage(v string) *SubmitSmsQualificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitSmsQualificationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitSmsQualificationResponseBody
	GetSuccess() *bool
}

type SubmitSmsQualificationResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1000****
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

func (s SubmitSmsQualificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmsQualificationResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSmsQualificationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SubmitSmsQualificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitSmsQualificationResponseBody) GetData() *string {
	return s.Data
}

func (s *SubmitSmsQualificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitSmsQualificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSmsQualificationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitSmsQualificationResponseBody) SetAccessDeniedDetail(v string) *SubmitSmsQualificationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SubmitSmsQualificationResponseBody) SetCode(v string) *SubmitSmsQualificationResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitSmsQualificationResponseBody) SetData(v string) *SubmitSmsQualificationResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitSmsQualificationResponseBody) SetMessage(v string) *SubmitSmsQualificationResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitSmsQualificationResponseBody) SetRequestId(v string) *SubmitSmsQualificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSmsQualificationResponseBody) SetSuccess(v bool) *SubmitSmsQualificationResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitSmsQualificationResponseBody) Validate() error {
	return dara.Validate(s)
}
