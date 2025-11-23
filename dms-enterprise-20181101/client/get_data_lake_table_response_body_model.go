// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakeTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetDataLakeTableResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataLakeTableResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDataLakeTableResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetDataLakeTableResponseBody
	GetSuccess() *string
	SetTable(v *DLTable) *GetDataLakeTableResponseBody
	GetTable() *DLTable
}

type GetDataLakeTableResponseBody struct {
	// The error code that is returned if the request failed.
	//
	// example:
	//
	// 400
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// Unknown server error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID. You can use the request ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 4E1D2B4D-3E53-4ABC-999D-1D2520B3471A
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
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The information of the table.
	Table *DLTable `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s GetDataLakeTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakeTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataLakeTableResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataLakeTableResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataLakeTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataLakeTableResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetDataLakeTableResponseBody) GetTable() *DLTable {
	return s.Table
}

func (s *GetDataLakeTableResponseBody) SetErrorCode(v string) *GetDataLakeTableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataLakeTableResponseBody) SetErrorMessage(v string) *GetDataLakeTableResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataLakeTableResponseBody) SetRequestId(v string) *GetDataLakeTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataLakeTableResponseBody) SetSuccess(v string) *GetDataLakeTableResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataLakeTableResponseBody) SetTable(v *DLTable) *GetDataLakeTableResponseBody {
	s.Table = v
	return s
}

func (s *GetDataLakeTableResponseBody) Validate() error {
	if s.Table != nil {
		if err := s.Table.Validate(); err != nil {
			return err
		}
	}
	return nil
}
