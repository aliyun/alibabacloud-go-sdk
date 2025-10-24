// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMajorProtectionBlackIpV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMajorProtectionBlackIpV2ResponseBody
	GetRequestId() *string
}

type CreateMajorProtectionBlackIpV2ResponseBody struct {
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMajorProtectionBlackIpV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMajorProtectionBlackIpV2ResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMajorProtectionBlackIpV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMajorProtectionBlackIpV2ResponseBody) SetRequestId(v string) *CreateMajorProtectionBlackIpV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMajorProtectionBlackIpV2ResponseBody) Validate() error {
	return dara.Validate(s)
}
