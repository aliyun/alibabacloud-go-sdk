// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseDataCorrectSQLJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *PauseDataCorrectSQLJobResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *PauseDataCorrectSQLJobResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *PauseDataCorrectSQLJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PauseDataCorrectSQLJobResponseBody
	GetSuccess() *bool
}

type PauseDataCorrectSQLJobResponseBody struct {
	// The error code that is returned.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 335C5BC8-490C-56EF-BDDE-94A4B53FEB72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PauseDataCorrectSQLJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PauseDataCorrectSQLJobResponseBody) GoString() string {
	return s.String()
}

func (s *PauseDataCorrectSQLJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *PauseDataCorrectSQLJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *PauseDataCorrectSQLJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PauseDataCorrectSQLJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PauseDataCorrectSQLJobResponseBody) SetErrorCode(v string) *PauseDataCorrectSQLJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PauseDataCorrectSQLJobResponseBody) SetErrorMessage(v string) *PauseDataCorrectSQLJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *PauseDataCorrectSQLJobResponseBody) SetRequestId(v string) *PauseDataCorrectSQLJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *PauseDataCorrectSQLJobResponseBody) SetSuccess(v bool) *PauseDataCorrectSQLJobResponseBody {
	s.Success = &v
	return s
}

func (s *PauseDataCorrectSQLJobResponseBody) Validate() error {
	return dara.Validate(s)
}
