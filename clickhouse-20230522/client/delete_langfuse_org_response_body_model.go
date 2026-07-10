// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLangfuseOrgResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteLangfuseOrgResponseBodyData) *DeleteLangfuseOrgResponseBody
	GetData() *DeleteLangfuseOrgResponseBodyData
	SetRequestId(v string) *DeleteLangfuseOrgResponseBody
	GetRequestId() *string
}

type DeleteLangfuseOrgResponseBody struct {
	// The response result.
	Data *DeleteLangfuseOrgResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLangfuseOrgResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLangfuseOrgResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLangfuseOrgResponseBody) GetData() *DeleteLangfuseOrgResponseBodyData {
	return s.Data
}

func (s *DeleteLangfuseOrgResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLangfuseOrgResponseBody) SetData(v *DeleteLangfuseOrgResponseBodyData) *DeleteLangfuseOrgResponseBody {
	s.Data = v
	return s
}

func (s *DeleteLangfuseOrgResponseBody) SetRequestId(v string) *DeleteLangfuseOrgResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLangfuseOrgResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteLangfuseOrgResponseBodyData struct {
	// The Langfuse instance ID.
	//
	// example:
	//
	// lfs-****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
}

func (s DeleteLangfuseOrgResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteLangfuseOrgResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteLangfuseOrgResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteLangfuseOrgResponseBodyData) SetDBInstanceId(v string) *DeleteLangfuseOrgResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteLangfuseOrgResponseBodyData) Validate() error {
	return dara.Validate(s)
}
