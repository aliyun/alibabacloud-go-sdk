// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBClusterVersionZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeDBClusterVersionZonalResponseBody
	GetRequestId() *string
}

type UpgradeDBClusterVersionZonalResponseBody struct {
	// example:
	//
	// CAE6755F-B79A-4861-B227-801FE8******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeDBClusterVersionZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBClusterVersionZonalResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBClusterVersionZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeDBClusterVersionZonalResponseBody) SetRequestId(v string) *UpgradeDBClusterVersionZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeDBClusterVersionZonalResponseBody) Validate() error {
	return dara.Validate(s)
}
