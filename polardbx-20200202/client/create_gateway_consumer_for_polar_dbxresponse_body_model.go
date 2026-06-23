// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayConsumerForPolarDBXResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateGatewayConsumerForPolarDBXResponseBody
	GetDBInstanceName() *string
	SetRequestId(v string) *CreateGatewayConsumerForPolarDBXResponseBody
	GetRequestId() *string
	SetTaskId(v int32) *CreateGatewayConsumerForPolarDBXResponseBody
	GetTaskId() *int32
}

type CreateGatewayConsumerForPolarDBXResponseBody struct {
	// The instance name.
	//
	// example:
	//
	// pxc-**************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A501A191-BD70-5E50-98A9-C2A486A82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The backend task ID.
	//
	// example:
	//
	// 422922413
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateGatewayConsumerForPolarDBXResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayConsumerForPolarDBXResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewayConsumerForPolarDBXResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateGatewayConsumerForPolarDBXResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGatewayConsumerForPolarDBXResponseBody) GetTaskId() *int32 {
	return s.TaskId
}

func (s *CreateGatewayConsumerForPolarDBXResponseBody) SetDBInstanceName(v string) *CreateGatewayConsumerForPolarDBXResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *CreateGatewayConsumerForPolarDBXResponseBody) SetRequestId(v string) *CreateGatewayConsumerForPolarDBXResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewayConsumerForPolarDBXResponseBody) SetTaskId(v int32) *CreateGatewayConsumerForPolarDBXResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateGatewayConsumerForPolarDBXResponseBody) Validate() error {
	return dara.Validate(s)
}
