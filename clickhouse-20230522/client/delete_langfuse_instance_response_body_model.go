// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLangfuseInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteLangfuseInstanceResponseBodyData) *DeleteLangfuseInstanceResponseBody
	GetData() *DeleteLangfuseInstanceResponseBodyData
	SetRequestId(v string) *DeleteLangfuseInstanceResponseBody
	GetRequestId() *string
}

type DeleteLangfuseInstanceResponseBody struct {
	// The response data.
	Data *DeleteLangfuseInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLangfuseInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLangfuseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLangfuseInstanceResponseBody) GetData() *DeleteLangfuseInstanceResponseBodyData {
	return s.Data
}

func (s *DeleteLangfuseInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLangfuseInstanceResponseBody) SetData(v *DeleteLangfuseInstanceResponseBodyData) *DeleteLangfuseInstanceResponseBody {
	s.Data = v
	return s
}

func (s *DeleteLangfuseInstanceResponseBody) SetRequestId(v string) *DeleteLangfuseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLangfuseInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteLangfuseInstanceResponseBodyData struct {
	// The Langfuse instance ID.
	//
	// example:
	//
	// lfs-****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DeleteLangfuseInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteLangfuseInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteLangfuseInstanceResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteLangfuseInstanceResponseBodyData) SetDBInstanceId(v string) *DeleteLangfuseInstanceResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteLangfuseInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
