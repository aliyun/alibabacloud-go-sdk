// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSupportAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DisableSupportAccountResponseBody
	GetData() *string
	SetErrorCode(v string) *DisableSupportAccountResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DisableSupportAccountResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *DisableSupportAccountResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *DisableSupportAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableSupportAccountResponseBody
	GetSuccess() *bool
}

type DisableSupportAccountResponseBody struct {
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2A8DEF6E-067E-5DB0-BAE1-2894266E6C6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableSupportAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableSupportAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSupportAccountResponseBody) GetData() *string {
	return s.Data
}

func (s *DisableSupportAccountResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DisableSupportAccountResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DisableSupportAccountResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DisableSupportAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableSupportAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableSupportAccountResponseBody) SetData(v string) *DisableSupportAccountResponseBody {
	s.Data = &v
	return s
}

func (s *DisableSupportAccountResponseBody) SetErrorCode(v string) *DisableSupportAccountResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DisableSupportAccountResponseBody) SetErrorMessage(v string) *DisableSupportAccountResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DisableSupportAccountResponseBody) SetHttpStatusCode(v string) *DisableSupportAccountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisableSupportAccountResponseBody) SetRequestId(v string) *DisableSupportAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSupportAccountResponseBody) SetSuccess(v bool) *DisableSupportAccountResponseBody {
	s.Success = &v
	return s
}

func (s *DisableSupportAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
