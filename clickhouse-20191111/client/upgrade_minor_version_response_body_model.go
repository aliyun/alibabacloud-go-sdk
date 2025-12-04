// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeMinorVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeMinorVersionResponseBody
	GetRequestId() *string
}

type UpgradeMinorVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FE242962-6DA3-5FC8-9691-37B62A3210F7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeMinorVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeMinorVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeMinorVersionResponseBody) SetRequestId(v string) *UpgradeMinorVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeMinorVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
