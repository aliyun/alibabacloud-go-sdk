// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteDBInstanceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DeleteDBInstanceResponseBody
	GetRequestId() *string
}

type DeleteDBInstanceResponseBody struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBInstanceResponseBody) SetRegionId(v string) *DeleteDBInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *DeleteDBInstanceResponseBody) SetRequestId(v string) *DeleteDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
