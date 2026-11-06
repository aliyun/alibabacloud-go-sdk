// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDBClusterCTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableDBClusterCTableResponseBody
	GetRequestId() *string
}

type DisableDBClusterCTableResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableDBClusterCTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableDBClusterCTableResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDBClusterCTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableDBClusterCTableResponseBody) SetRequestId(v string) *DisableDBClusterCTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableDBClusterCTableResponseBody) Validate() error {
	return dara.Validate(s)
}
