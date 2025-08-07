// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBClusterVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeDBClusterVersionResponseBody
	GetRequestId() *string
}

type UpgradeDBClusterVersionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CAE6755F-B79A-4861-B227-801FE8******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeDBClusterVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBClusterVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBClusterVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeDBClusterVersionResponseBody) SetRequestId(v string) *UpgradeDBClusterVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeDBClusterVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
