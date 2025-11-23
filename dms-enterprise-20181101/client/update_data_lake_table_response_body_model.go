// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataLakeTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateDataLakeTableResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateDataLakeTableResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateDataLakeTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataLakeTableResponseBody
	GetSuccess() *bool
	SetTable(v *DLTable) *UpdateDataLakeTableResponseBody
	GetTable() *DLTable
}

type UpdateDataLakeTableResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID. You can use the request ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// C5B8E84B-42B6-4374-AD5A-6264E1753325
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The information about the table.
	Table *DLTable `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s UpdateDataLakeTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataLakeTableResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataLakeTableResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateDataLakeTableResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateDataLakeTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataLakeTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataLakeTableResponseBody) GetTable() *DLTable {
	return s.Table
}

func (s *UpdateDataLakeTableResponseBody) SetErrorCode(v string) *UpdateDataLakeTableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDataLakeTableResponseBody) SetErrorMessage(v string) *UpdateDataLakeTableResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDataLakeTableResponseBody) SetRequestId(v string) *UpdateDataLakeTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataLakeTableResponseBody) SetSuccess(v bool) *UpdateDataLakeTableResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataLakeTableResponseBody) SetTable(v *DLTable) *UpdateDataLakeTableResponseBody {
	s.Table = v
	return s
}

func (s *UpdateDataLakeTableResponseBody) Validate() error {
	if s.Table != nil {
		if err := s.Table.Validate(); err != nil {
			return err
		}
	}
	return nil
}
