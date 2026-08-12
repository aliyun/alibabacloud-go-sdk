// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecuritySuggestionNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSecuritySuggestionNumberResponseBody
	GetCode() *string
	SetData(v int64) *GetSecuritySuggestionNumberResponseBody
	GetData() *int64
	SetMessage(v string) *GetSecuritySuggestionNumberResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSecuritySuggestionNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSecuritySuggestionNumberResponseBody
	GetSuccess() *bool
}

type GetSecuritySuggestionNumberResponseBody struct {
	// The status code.
	//
	// - **200**: Successful.
	//
	// - **Others (400, 500)**: Failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of pending items.
	//
	// example:
	//
	// 5
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2E130B0F-9E69-52FA-84FC-187FE1BA9489
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSecuritySuggestionNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecuritySuggestionNumberResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecuritySuggestionNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSecuritySuggestionNumberResponseBody) GetData() *int64 {
	return s.Data
}

func (s *GetSecuritySuggestionNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSecuritySuggestionNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecuritySuggestionNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSecuritySuggestionNumberResponseBody) SetCode(v string) *GetSecuritySuggestionNumberResponseBody {
	s.Code = &v
	return s
}

func (s *GetSecuritySuggestionNumberResponseBody) SetData(v int64) *GetSecuritySuggestionNumberResponseBody {
	s.Data = &v
	return s
}

func (s *GetSecuritySuggestionNumberResponseBody) SetMessage(v string) *GetSecuritySuggestionNumberResponseBody {
	s.Message = &v
	return s
}

func (s *GetSecuritySuggestionNumberResponseBody) SetRequestId(v string) *GetSecuritySuggestionNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecuritySuggestionNumberResponseBody) SetSuccess(v bool) *GetSecuritySuggestionNumberResponseBody {
	s.Success = &v
	return s
}

func (s *GetSecuritySuggestionNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
