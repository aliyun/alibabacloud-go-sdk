// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateDBResponseBodyData) *CreateDBResponseBody
	GetData() *CreateDBResponseBodyData
	SetRequestId(v string) *CreateDBResponseBody
	GetRequestId() *string
}

type CreateDBResponseBody struct {
	// The data returned.
	Data *CreateDBResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 94F92113-FF63-5E57-8401-6FE123AD11DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBResponseBody) GetData() *CreateDBResponseBodyData {
	return s.Data
}

func (s *CreateDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBResponseBody) SetData(v *CreateDBResponseBodyData) *CreateDBResponseBody {
	s.Data = v
	return s
}

func (s *CreateDBResponseBody) SetRequestId(v string) *CreateDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDBResponseBodyData struct {
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

func (s CreateDBResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDBResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBResponseBodyData) GetDBName() *string {
	return s.DBName
}

func (s *CreateDBResponseBodyData) SetDBInstanceId(v string) *CreateDBResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBResponseBodyData) SetDBName(v string) *CreateDBResponseBodyData {
	s.DBName = &v
	return s
}

func (s *CreateDBResponseBodyData) Validate() error {
	return dara.Validate(s)
}
