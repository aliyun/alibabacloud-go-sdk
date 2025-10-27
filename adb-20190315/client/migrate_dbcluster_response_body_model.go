// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateDBClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MigrateDBClusterResponseBody
	GetRequestId() *string
}

type MigrateDBClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F0983B43-B2EC-536A-8791-142B5CF1E9B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MigrateDBClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MigrateDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateDBClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MigrateDBClusterResponseBody) SetRequestId(v string) *MigrateDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateDBClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
