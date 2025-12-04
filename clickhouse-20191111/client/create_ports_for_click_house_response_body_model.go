// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePortsForClickHouseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreatePortsForClickHouseResponseBody
	GetRequestId() *string
}

type CreatePortsForClickHouseResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePortsForClickHouseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePortsForClickHouseResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePortsForClickHouseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePortsForClickHouseResponseBody) SetRequestId(v string) *CreatePortsForClickHouseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePortsForClickHouseResponseBody) Validate() error {
	return dara.Validate(s)
}
