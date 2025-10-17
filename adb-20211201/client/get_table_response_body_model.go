// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetTableResponseBody
	GetCode() *int64
	SetMessage(v string) *GetTableResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTableResponseBody
	GetSuccess() *bool
	SetTable(v *TableModel) *GetTableResponseBody
	GetTable() *TableModel
}

type GetTableResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the query succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The information about the table.
	Table *TableModel `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s GetTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetTableResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTableResponseBody) GetTable() *TableModel {
	return s.Table
}

func (s *GetTableResponseBody) SetCode(v int64) *GetTableResponseBody {
	s.Code = &v
	return s
}

func (s *GetTableResponseBody) SetMessage(v string) *GetTableResponseBody {
	s.Message = &v
	return s
}

func (s *GetTableResponseBody) SetRequestId(v string) *GetTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableResponseBody) SetSuccess(v bool) *GetTableResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableResponseBody) SetTable(v *TableModel) *GetTableResponseBody {
	s.Table = v
	return s
}

func (s *GetTableResponseBody) Validate() error {
	if s.Table != nil {
		if err := s.Table.Validate(); err != nil {
			return err
		}
	}
	return nil
}
