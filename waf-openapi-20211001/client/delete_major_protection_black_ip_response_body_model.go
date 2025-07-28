// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMajorProtectionBlackIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMajorProtectionBlackIpResponseBody
	GetRequestId() *string
}

type DeleteMajorProtectionBlackIpResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 66A98669-CC6E-4F3E-80A6-3014697B11AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMajorProtectionBlackIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMajorProtectionBlackIpResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMajorProtectionBlackIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMajorProtectionBlackIpResponseBody) SetRequestId(v string) *DeleteMajorProtectionBlackIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMajorProtectionBlackIpResponseBody) Validate() error {
	return dara.Validate(s)
}
