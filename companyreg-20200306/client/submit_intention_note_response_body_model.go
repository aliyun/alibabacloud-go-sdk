// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIntentionNoteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SubmitIntentionNoteResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *SubmitIntentionNoteResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *SubmitIntentionNoteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitIntentionNoteResponseBody
	GetSuccess() *bool
}

type SubmitIntentionNoteResponseBody struct {
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 668571EF-1E7E-5435-AA65-4ECFFDDA133F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitIntentionNoteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitIntentionNoteResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitIntentionNoteResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SubmitIntentionNoteResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *SubmitIntentionNoteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitIntentionNoteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitIntentionNoteResponseBody) SetErrorCode(v string) *SubmitIntentionNoteResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SubmitIntentionNoteResponseBody) SetErrorMsg(v string) *SubmitIntentionNoteResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SubmitIntentionNoteResponseBody) SetRequestId(v string) *SubmitIntentionNoteResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitIntentionNoteResponseBody) SetSuccess(v bool) *SubmitIntentionNoteResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitIntentionNoteResponseBody) Validate() error {
	return dara.Validate(s)
}
