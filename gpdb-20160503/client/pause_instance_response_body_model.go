// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *PauseInstanceResponseBody
	GetDBInstanceId() *string
	SetErrorMessage(v string) *PauseInstanceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *PauseInstanceResponseBody
	GetRequestId() *string
	SetStatus(v bool) *PauseInstanceResponseBody
	GetStatus() *bool
}

type PauseInstanceResponseBody struct {
	// The ID of the instance.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The error message returned.
	//
	// This parameter is returned only if **false*	- is returned for the **Status*	- parameter.
	//
	// example:
	//
	// *******
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 34b32a0a-08ef-4a87-b6be-cdd9********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **false**: The request failed.
	//
	// 	- **true**: The request was successful.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PauseInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PauseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *PauseInstanceResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *PauseInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *PauseInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PauseInstanceResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *PauseInstanceResponseBody) SetDBInstanceId(v string) *PauseInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *PauseInstanceResponseBody) SetErrorMessage(v string) *PauseInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *PauseInstanceResponseBody) SetRequestId(v string) *PauseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PauseInstanceResponseBody) SetStatus(v bool) *PauseInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *PauseInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
