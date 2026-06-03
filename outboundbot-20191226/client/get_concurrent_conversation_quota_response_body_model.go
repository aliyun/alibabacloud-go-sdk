// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConcurrentConversationQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetConcurrentConversationQuotaResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetConcurrentConversationQuotaResponseBody
	GetHttpStatusCode() *int32
	SetMaxConcurrent(v int32) *GetConcurrentConversationQuotaResponseBody
	GetMaxConcurrent() *int32
	SetMessage(v string) *GetConcurrentConversationQuotaResponseBody
	GetMessage() *string
	SetRemainingConcurrent(v int32) *GetConcurrentConversationQuotaResponseBody
	GetRemainingConcurrent() *int32
	SetRequestId(v string) *GetConcurrentConversationQuotaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetConcurrentConversationQuotaResponseBody
	GetSuccess() *bool
}

type GetConcurrentConversationQuotaResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// maxConcurrent
	//
	// example:
	//
	// 10
	MaxConcurrent *int32 `json:"MaxConcurrent,omitempty" xml:"MaxConcurrent,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2
	RemainingConcurrent *int32 `json:"RemainingConcurrent,omitempty" xml:"RemainingConcurrent,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetConcurrentConversationQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConcurrentConversationQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetConcurrentConversationQuotaResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetConcurrentConversationQuotaResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetConcurrentConversationQuotaResponseBody) GetMaxConcurrent() *int32 {
	return s.MaxConcurrent
}

func (s *GetConcurrentConversationQuotaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetConcurrentConversationQuotaResponseBody) GetRemainingConcurrent() *int32 {
	return s.RemainingConcurrent
}

func (s *GetConcurrentConversationQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConcurrentConversationQuotaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetConcurrentConversationQuotaResponseBody) SetCode(v string) *GetConcurrentConversationQuotaResponseBody {
	s.Code = &v
	return s
}

func (s *GetConcurrentConversationQuotaResponseBody) SetHttpStatusCode(v int32) *GetConcurrentConversationQuotaResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetConcurrentConversationQuotaResponseBody) SetMaxConcurrent(v int32) *GetConcurrentConversationQuotaResponseBody {
	s.MaxConcurrent = &v
	return s
}

func (s *GetConcurrentConversationQuotaResponseBody) SetMessage(v string) *GetConcurrentConversationQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *GetConcurrentConversationQuotaResponseBody) SetRemainingConcurrent(v int32) *GetConcurrentConversationQuotaResponseBody {
	s.RemainingConcurrent = &v
	return s
}

func (s *GetConcurrentConversationQuotaResponseBody) SetRequestId(v string) *GetConcurrentConversationQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConcurrentConversationQuotaResponseBody) SetSuccess(v bool) *GetConcurrentConversationQuotaResponseBody {
	s.Success = &v
	return s
}

func (s *GetConcurrentConversationQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
