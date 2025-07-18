// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ResumeInstanceResponseBody
	GetDBInstanceId() *string
	SetErrorMessage(v string) *ResumeInstanceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ResumeInstanceResponseBody
	GetRequestId() *string
	SetStatus(v bool) *ResumeInstanceResponseBody
	GetStatus() *bool
}

type ResumeInstanceResponseBody struct {
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
	// ******
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

func (s ResumeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeInstanceResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ResumeInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ResumeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeInstanceResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *ResumeInstanceResponseBody) SetDBInstanceId(v string) *ResumeInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetErrorMessage(v string) *ResumeInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetRequestId(v string) *ResumeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetStatus(v bool) *ResumeInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *ResumeInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
