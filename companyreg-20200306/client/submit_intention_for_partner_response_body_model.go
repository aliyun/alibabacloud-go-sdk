// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIntentionForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMsg(v string) *SubmitIntentionForPartnerResponseBody
	GetErrorMsg() *string
	SetIntentionBizId(v string) *SubmitIntentionForPartnerResponseBody
	GetIntentionBizId() *string
	SetRequestId(v string) *SubmitIntentionForPartnerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitIntentionForPartnerResponseBody
	GetSuccess() *bool
}

type SubmitIntentionForPartnerResponseBody struct {
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// I20211223101045000001
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6A603AA0-73BA-52B3-AC7D-0F846ECF7A9D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitIntentionForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitIntentionForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitIntentionForPartnerResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *SubmitIntentionForPartnerResponseBody) GetIntentionBizId() *string {
	return s.IntentionBizId
}

func (s *SubmitIntentionForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitIntentionForPartnerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitIntentionForPartnerResponseBody) SetErrorMsg(v string) *SubmitIntentionForPartnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SubmitIntentionForPartnerResponseBody) SetIntentionBizId(v string) *SubmitIntentionForPartnerResponseBody {
	s.IntentionBizId = &v
	return s
}

func (s *SubmitIntentionForPartnerResponseBody) SetRequestId(v string) *SubmitIntentionForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitIntentionForPartnerResponseBody) SetSuccess(v bool) *SubmitIntentionForPartnerResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitIntentionForPartnerResponseBody) Validate() error {
	return dara.Validate(s)
}
