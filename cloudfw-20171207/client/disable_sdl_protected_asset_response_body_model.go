// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSdlProtectedAssetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableSdlProtectedAssetResponseBody
	GetRequestId() *string
}

type DisableSdlProtectedAssetResponseBody struct {
	// example:
	//
	// F93A490D-9E92-5AA4-BA79-600FFC09****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableSdlProtectedAssetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableSdlProtectedAssetResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSdlProtectedAssetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableSdlProtectedAssetResponseBody) SetRequestId(v string) *DisableSdlProtectedAssetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSdlProtectedAssetResponseBody) Validate() error {
	return dara.Validate(s)
}
