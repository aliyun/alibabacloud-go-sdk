// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMajorProtectionBlackIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMajorProtectionBlackIpResponseBody
	GetRequestId() *string
}

type CreateMajorProtectionBlackIpResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMajorProtectionBlackIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMajorProtectionBlackIpResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMajorProtectionBlackIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMajorProtectionBlackIpResponseBody) SetRequestId(v string) *CreateMajorProtectionBlackIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMajorProtectionBlackIpResponseBody) Validate() error {
	return dara.Validate(s)
}
