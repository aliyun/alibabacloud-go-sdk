// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdnsSwitchRuleCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SdnsSwitchRuleCategoryResponseBody
	GetRequestId() *string
}

type SdnsSwitchRuleCategoryResponseBody struct {
	// example:
	//
	// 50F9C40E-188D-4208-BE2C-74271337****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SdnsSwitchRuleCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SdnsSwitchRuleCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *SdnsSwitchRuleCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SdnsSwitchRuleCategoryResponseBody) SetRequestId(v string) *SdnsSwitchRuleCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *SdnsSwitchRuleCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
