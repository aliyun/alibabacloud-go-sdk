// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMajorProtectionBlackIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyMajorProtectionBlackIpResponseBody
	GetRequestId() *string
}

type ModifyMajorProtectionBlackIpResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 66A98669-CC6E-4F3E-80A6-3014697B11AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMajorProtectionBlackIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMajorProtectionBlackIpResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMajorProtectionBlackIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMajorProtectionBlackIpResponseBody) SetRequestId(v string) *ModifyMajorProtectionBlackIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMajorProtectionBlackIpResponseBody) Validate() error {
	return dara.Validate(s)
}
