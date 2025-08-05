// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelBackupJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelBackupJobResponseBody
	GetCode() *string
	SetMessage(v string) *CancelBackupJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelBackupJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelBackupJobResponseBody
	GetSuccess() *bool
}

type CancelBackupJobResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned. If the request is successful, a value of successful is returned. If the request fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// 	- true: The request is successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelBackupJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelBackupJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelBackupJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelBackupJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelBackupJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelBackupJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelBackupJobResponseBody) SetCode(v string) *CancelBackupJobResponseBody {
	s.Code = &v
	return s
}

func (s *CancelBackupJobResponseBody) SetMessage(v string) *CancelBackupJobResponseBody {
	s.Message = &v
	return s
}

func (s *CancelBackupJobResponseBody) SetRequestId(v string) *CancelBackupJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelBackupJobResponseBody) SetSuccess(v bool) *CancelBackupJobResponseBody {
	s.Success = &v
	return s
}

func (s *CancelBackupJobResponseBody) Validate() error {
	return dara.Validate(s)
}
