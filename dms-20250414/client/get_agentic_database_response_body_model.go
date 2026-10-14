// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgenticDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AgenticDatabase) *GetAgenticDatabaseResponseBody
	GetData() *AgenticDatabase
	SetErrorCode(v string) *GetAgenticDatabaseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetAgenticDatabaseResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetAgenticDatabaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgenticDatabaseResponseBody
	GetSuccess() *bool
}

type GetAgenticDatabaseResponseBody struct {
	// The database details.
	Data *AgenticDatabase `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned when an exception occurs.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgenticDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgenticDatabaseResponseBody) GetData() *AgenticDatabase {
	return s.Data
}

func (s *GetAgenticDatabaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAgenticDatabaseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAgenticDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgenticDatabaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgenticDatabaseResponseBody) SetData(v *AgenticDatabase) *GetAgenticDatabaseResponseBody {
	s.Data = v
	return s
}

func (s *GetAgenticDatabaseResponseBody) SetErrorCode(v string) *GetAgenticDatabaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAgenticDatabaseResponseBody) SetErrorMessage(v string) *GetAgenticDatabaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAgenticDatabaseResponseBody) SetRequestId(v string) *GetAgenticDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgenticDatabaseResponseBody) SetSuccess(v bool) *GetAgenticDatabaseResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgenticDatabaseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
