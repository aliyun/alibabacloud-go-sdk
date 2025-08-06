// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseIntentionForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CloseIntentionForPartnerResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *CloseIntentionForPartnerResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *CloseIntentionForPartnerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CloseIntentionForPartnerResponseBody
	GetSuccess() *bool
}

type CloseIntentionForPartnerResponseBody struct {
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 4674B06A-B57F-5922-890C-D23D17C5BD21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloseIntentionForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseIntentionForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *CloseIntentionForPartnerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CloseIntentionForPartnerResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CloseIntentionForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseIntentionForPartnerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CloseIntentionForPartnerResponseBody) SetErrorCode(v string) *CloseIntentionForPartnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CloseIntentionForPartnerResponseBody) SetErrorMsg(v string) *CloseIntentionForPartnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CloseIntentionForPartnerResponseBody) SetRequestId(v string) *CloseIntentionForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseIntentionForPartnerResponseBody) SetSuccess(v bool) *CloseIntentionForPartnerResponseBody {
	s.Success = &v
	return s
}

func (s *CloseIntentionForPartnerResponseBody) Validate() error {
	return dara.Validate(s)
}
