// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyDBInstanceConfigResponseBodyData) *ModifyDBInstanceConfigResponseBody
	GetData() *ModifyDBInstanceConfigResponseBodyData
	SetRequestId(v string) *ModifyDBInstanceConfigResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceConfigResponseBody struct {
	Data *ModifyDBInstanceConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 05321590-BB65-4720-8C***********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConfigResponseBody) GetData() *ModifyDBInstanceConfigResponseBodyData {
	return s.Data
}

func (s *ModifyDBInstanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceConfigResponseBody) SetData(v *ModifyDBInstanceConfigResponseBodyData) *ModifyDBInstanceConfigResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDBInstanceConfigResponseBody) SetRequestId(v string) *ModifyDBInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyDBInstanceConfigResponseBodyData struct {
	// example:
	//
	// cc-uf6lkzf*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s ModifyDBInstanceConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConfigResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceConfigResponseBodyData) SetDBInstanceId(v string) *ModifyDBInstanceConfigResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
