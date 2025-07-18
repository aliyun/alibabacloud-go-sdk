// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnloadSampleDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UnloadSampleDataResponseBody
	GetDBInstanceId() *string
	SetErrorMessage(v string) *UnloadSampleDataResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UnloadSampleDataResponseBody
	GetRequestId() *string
	SetStatus(v bool) *UnloadSampleDataResponseBody
	GetStatus() *bool
}

type UnloadSampleDataResponseBody struct {
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
	// ********
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4E42ABC3-4DBD-5343-9BCA-66B7D091311F_6914
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

func (s UnloadSampleDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnloadSampleDataResponseBody) GoString() string {
	return s.String()
}

func (s *UnloadSampleDataResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UnloadSampleDataResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UnloadSampleDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnloadSampleDataResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *UnloadSampleDataResponseBody) SetDBInstanceId(v string) *UnloadSampleDataResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *UnloadSampleDataResponseBody) SetErrorMessage(v string) *UnloadSampleDataResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UnloadSampleDataResponseBody) SetRequestId(v string) *UnloadSampleDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnloadSampleDataResponseBody) SetStatus(v bool) *UnloadSampleDataResponseBody {
	s.Status = &v
	return s
}

func (s *UnloadSampleDataResponseBody) Validate() error {
	return dara.Validate(s)
}
