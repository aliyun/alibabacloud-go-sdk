// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseUserIntentionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CloseUserIntentionResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *CloseUserIntentionResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *CloseUserIntentionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CloseUserIntentionResponseBody
	GetSuccess() *bool
}

type CloseUserIntentionResponseBody struct {
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// DD5639FF-1240-51DE-9BA8-2075670A1EAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloseUserIntentionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseUserIntentionResponseBody) GoString() string {
	return s.String()
}

func (s *CloseUserIntentionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CloseUserIntentionResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CloseUserIntentionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseUserIntentionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CloseUserIntentionResponseBody) SetErrorCode(v string) *CloseUserIntentionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CloseUserIntentionResponseBody) SetErrorMsg(v string) *CloseUserIntentionResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CloseUserIntentionResponseBody) SetRequestId(v string) *CloseUserIntentionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseUserIntentionResponseBody) SetSuccess(v bool) *CloseUserIntentionResponseBody {
	s.Success = &v
	return s
}

func (s *CloseUserIntentionResponseBody) Validate() error {
	return dara.Validate(s)
}
