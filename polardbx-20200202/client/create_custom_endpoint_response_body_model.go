// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateCustomEndpointResponseBodyData) *CreateCustomEndpointResponseBody
	GetData() *CreateCustomEndpointResponseBodyData
	SetRequestId(v string) *CreateCustomEndpointResponseBody
	GetRequestId() *string
}

type CreateCustomEndpointResponseBody struct {
	Data *CreateCustomEndpointResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 2DFF784E-DC31-5BBC-9B25-9261CD9E0AA9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomEndpointResponseBody) GetData() *CreateCustomEndpointResponseBodyData {
	return s.Data
}

func (s *CreateCustomEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomEndpointResponseBody) SetData(v *CreateCustomEndpointResponseBodyData) *CreateCustomEndpointResponseBody {
	s.Data = v
	return s
}

func (s *CreateCustomEndpointResponseBody) SetRequestId(v string) *CreateCustomEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomEndpointResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCustomEndpointResponseBodyData struct {
	// example:
	//
	// pxe-c4gkgqg****7sgyg
	CustomEndpointId *string `json:"CustomEndpointId,omitempty" xml:"CustomEndpointId,omitempty"`
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
}

func (s CreateCustomEndpointResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomEndpointResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCustomEndpointResponseBodyData) GetCustomEndpointId() *string {
	return s.CustomEndpointId
}

func (s *CreateCustomEndpointResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateCustomEndpointResponseBodyData) SetCustomEndpointId(v string) *CreateCustomEndpointResponseBodyData {
	s.CustomEndpointId = &v
	return s
}

func (s *CreateCustomEndpointResponseBodyData) SetDBInstanceName(v string) *CreateCustomEndpointResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *CreateCustomEndpointResponseBodyData) Validate() error {
	return dara.Validate(s)
}
