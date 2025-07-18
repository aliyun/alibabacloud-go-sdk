// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeExtensionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeExtensionsResponseBody
	GetRequestId() *string
}

type UpgradeExtensionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeExtensionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeExtensionsResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeExtensionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeExtensionsResponseBody) SetRequestId(v string) *UpgradeExtensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeExtensionsResponseBody) Validate() error {
	return dara.Validate(s)
}
