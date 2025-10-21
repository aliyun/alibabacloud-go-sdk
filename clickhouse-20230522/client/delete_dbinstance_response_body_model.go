// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteDBInstanceResponseBodyData) *DeleteDBInstanceResponseBody
	GetData() *DeleteDBInstanceResponseBodyData
	SetRequestId(v string) *DeleteDBInstanceResponseBody
	GetRequestId() *string
}

type DeleteDBInstanceResponseBody struct {
	// The data returned.
	Data *DeleteDBInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceResponseBody) GetData() *DeleteDBInstanceResponseBodyData {
	return s.Data
}

func (s *DeleteDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBInstanceResponseBody) SetData(v *DeleteDBInstanceResponseBodyData) *DeleteDBInstanceResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDBInstanceResponseBody) SetRequestId(v string) *DeleteDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteDBInstanceResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DeleteDBInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBInstanceResponseBodyData) SetDBInstanceId(v string) *DeleteDBInstanceResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
