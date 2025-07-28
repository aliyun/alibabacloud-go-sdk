// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearMajorProtectionBlackIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ClearMajorProtectionBlackIpResponseBody
	GetRequestId() *string
}

type ClearMajorProtectionBlackIpResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 66A98669-CC6E-4F3E-80A6-3014697B11AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClearMajorProtectionBlackIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClearMajorProtectionBlackIpResponseBody) GoString() string {
	return s.String()
}

func (s *ClearMajorProtectionBlackIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClearMajorProtectionBlackIpResponseBody) SetRequestId(v string) *ClearMajorProtectionBlackIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearMajorProtectionBlackIpResponseBody) Validate() error {
	return dara.Validate(s)
}
