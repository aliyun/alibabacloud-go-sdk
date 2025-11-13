// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteDBClusterResponseBodyData) *DeleteDBClusterResponseBody
	GetData() *DeleteDBClusterResponseBodyData
	SetRequestId(v string) *DeleteDBClusterResponseBody
	GetRequestId() *string
}

type DeleteDBClusterResponseBody struct {
	Data *DeleteDBClusterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// F203FA74-3041-589F-BE66-E570793A0C91
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterResponseBody) GetData() *DeleteDBClusterResponseBodyData {
	return s.Data
}

func (s *DeleteDBClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBClusterResponseBody) SetData(v *DeleteDBClusterResponseBodyData) *DeleteDBClusterResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDBClusterResponseBody) SetRequestId(v string) *DeleteDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBClusterResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteDBClusterResponseBodyData struct {
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// 220088764060782
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s DeleteDBClusterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBClusterResponseBodyData) GetOrderId() *string {
	return s.OrderId
}

func (s *DeleteDBClusterResponseBodyData) SetDBInstanceId(v string) *DeleteDBClusterResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBClusterResponseBodyData) SetOrderId(v string) *DeleteDBClusterResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *DeleteDBClusterResponseBodyData) Validate() error {
	return dara.Validate(s)
}
