// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdnsDeleteRuleCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SdnsDeleteRuleCategoryResponseBody
	GetRequestId() *string
}

type SdnsDeleteRuleCategoryResponseBody struct {
	// The request ID, which is used to troubleshoot API call issues. This is not the policy ID.
	//
	// example:
	//
	// 50F9C40E-188D-4208-BE2C-74271337****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SdnsDeleteRuleCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SdnsDeleteRuleCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *SdnsDeleteRuleCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SdnsDeleteRuleCategoryResponseBody) SetRequestId(v string) *SdnsDeleteRuleCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *SdnsDeleteRuleCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
