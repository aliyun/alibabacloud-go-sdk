// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteDBResponseBodyData) *DeleteDBResponseBody
	GetData() *DeleteDBResponseBodyData
	SetRequestId(v string) *DeleteDBResponseBody
	GetRequestId() *string
}

type DeleteDBResponseBody struct {
	// The data returned.
	Data *DeleteDBResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 06798FEE-BEF2-5FAF-A30D-728973BBE97C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBResponseBody) GetData() *DeleteDBResponseBodyData {
	return s.Data
}

func (s *DeleteDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBResponseBody) SetData(v *DeleteDBResponseBodyData) *DeleteDBResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDBResponseBody) SetRequestId(v string) *DeleteDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteDBResponseBodyData struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// testdb001
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
}

func (s DeleteDBResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDBResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBResponseBodyData) GetDBName() *string {
	return s.DBName
}

func (s *DeleteDBResponseBodyData) SetDBInstanceId(v string) *DeleteDBResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBResponseBodyData) SetDBName(v string) *DeleteDBResponseBodyData {
	s.DBName = &v
	return s
}

func (s *DeleteDBResponseBodyData) Validate() error {
	return dara.Validate(s)
}
