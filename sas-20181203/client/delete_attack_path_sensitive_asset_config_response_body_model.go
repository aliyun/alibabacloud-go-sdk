// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAttackPathSensitiveAssetConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAttackPathSensitiveAssetConfigResponseBody
	GetRequestId() *string
}

type DeleteAttackPathSensitiveAssetConfigResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D65AADFC-1D20-5A6A-8F6A-9FA53C0D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAttackPathSensitiveAssetConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAttackPathSensitiveAssetConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAttackPathSensitiveAssetConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAttackPathSensitiveAssetConfigResponseBody) SetRequestId(v string) *DeleteAttackPathSensitiveAssetConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAttackPathSensitiveAssetConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
