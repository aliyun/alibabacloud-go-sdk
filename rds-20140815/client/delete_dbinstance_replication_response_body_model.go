// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceReplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDBInstanceReplicationResponseBody
	GetRequestId() *string
}

type DeleteDBInstanceReplicationResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F2911788-25E8-42E5-A3A3-1B38D263F01E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBInstanceReplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceReplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceReplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBInstanceReplicationResponseBody) SetRequestId(v string) *DeleteDBInstanceReplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBInstanceReplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
