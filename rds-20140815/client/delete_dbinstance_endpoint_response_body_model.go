// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteDBInstanceEndpointResponseBodyData) *DeleteDBInstanceEndpointResponseBody
	GetData() *DeleteDBInstanceEndpointResponseBodyData
	SetRequestId(v string) *DeleteDBInstanceEndpointResponseBody
	GetRequestId() *string
}

type DeleteDBInstanceEndpointResponseBody struct {
	// The data returned.
	Data *DeleteDBInstanceEndpointResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 224DB9F7-3100-4899-AB9C-C938BCCB43E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBInstanceEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceEndpointResponseBody) GetData() *DeleteDBInstanceEndpointResponseBodyData {
	return s.Data
}

func (s *DeleteDBInstanceEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBInstanceEndpointResponseBody) SetData(v *DeleteDBInstanceEndpointResponseBodyData) *DeleteDBInstanceEndpointResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDBInstanceEndpointResponseBody) SetRequestId(v string) *DeleteDBInstanceEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBInstanceEndpointResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteDBInstanceEndpointResponseBodyData struct {
	// The endpoint ID of the instance.
	//
	// example:
	//
	// ep-****
	DBInstanceEndpointId *string `json:"DBInstanceEndpointId,omitempty" xml:"DBInstanceEndpointId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// rm-****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
}

func (s DeleteDBInstanceEndpointResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceEndpointResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceEndpointResponseBodyData) GetDBInstanceEndpointId() *string {
	return s.DBInstanceEndpointId
}

func (s *DeleteDBInstanceEndpointResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeleteDBInstanceEndpointResponseBodyData) SetDBInstanceEndpointId(v string) *DeleteDBInstanceEndpointResponseBodyData {
	s.DBInstanceEndpointId = &v
	return s
}

func (s *DeleteDBInstanceEndpointResponseBodyData) SetDBInstanceName(v string) *DeleteDBInstanceEndpointResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteDBInstanceEndpointResponseBodyData) Validate() error {
	return dara.Validate(s)
}
