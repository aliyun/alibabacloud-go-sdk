// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSampleDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateSampleDataResponseBody
	GetDBInstanceId() *string
	SetErrorMessage(v string) *CreateSampleDataResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateSampleDataResponseBody
	GetRequestId() *string
	SetStatus(v bool) *CreateSampleDataResponseBody
	GetStatus() *bool
}

type CreateSampleDataResponseBody struct {
	// The ID of the instance.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The error message returned if an error occurs. This message does not affect the execution of the operation.
	//
	// example:
	//
	// *********
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 15A6881F-AAAB-5E4D-9B3F-6DCC1BDF4F2E_99
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution state of the operation. Valid values:
	//
	// 	- **false**: The operation fails.
	//
	// 	- **true**: The operation is successful.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateSampleDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSampleDataResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSampleDataResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateSampleDataResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateSampleDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSampleDataResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *CreateSampleDataResponseBody) SetDBInstanceId(v string) *CreateSampleDataResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateSampleDataResponseBody) SetErrorMessage(v string) *CreateSampleDataResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateSampleDataResponseBody) SetRequestId(v string) *CreateSampleDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSampleDataResponseBody) SetStatus(v bool) *CreateSampleDataResponseBody {
	s.Status = &v
	return s
}

func (s *CreateSampleDataResponseBody) Validate() error {
	return dara.Validate(s)
}
