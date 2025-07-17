// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDataCorrectSQLJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *RestartDataCorrectSQLJobResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RestartDataCorrectSQLJobResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *RestartDataCorrectSQLJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RestartDataCorrectSQLJobResponseBody
	GetSuccess() *bool
}

type RestartDataCorrectSQLJobResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 809B52F2-CD2B-53DA-88C8-F7042787E673
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RestartDataCorrectSQLJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartDataCorrectSQLJobResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDataCorrectSQLJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RestartDataCorrectSQLJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RestartDataCorrectSQLJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartDataCorrectSQLJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RestartDataCorrectSQLJobResponseBody) SetErrorCode(v string) *RestartDataCorrectSQLJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RestartDataCorrectSQLJobResponseBody) SetErrorMessage(v string) *RestartDataCorrectSQLJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RestartDataCorrectSQLJobResponseBody) SetRequestId(v string) *RestartDataCorrectSQLJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartDataCorrectSQLJobResponseBody) SetSuccess(v bool) *RestartDataCorrectSQLJobResponseBody {
	s.Success = &v
	return s
}

func (s *RestartDataCorrectSQLJobResponseBody) Validate() error {
	return dara.Validate(s)
}
