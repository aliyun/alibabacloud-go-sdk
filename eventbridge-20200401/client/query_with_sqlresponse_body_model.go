// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWithSQLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryWithSQLResponseBody
	GetCode() *string
	SetData(v *QueryWithSQLResponseBodyData) *QueryWithSQLResponseBody
	GetData() *QueryWithSQLResponseBodyData
	SetMessage(v string) *QueryWithSQLResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryWithSQLResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryWithSQLResponseBody
	GetSuccess() *bool
}

type QueryWithSQLResponseBody struct {
	// The response code. Valid values:
	//
	// - Success: The request was successful.
	//
	// - Other values: An error occurred. For more information, see error codes.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The event trace information.
	Data *QueryWithSQLResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. A value of true indicates success. A value of false indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryWithSQLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryWithSQLResponseBody) GoString() string {
	return s.String()
}

func (s *QueryWithSQLResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryWithSQLResponseBody) GetData() *QueryWithSQLResponseBodyData {
	return s.Data
}

func (s *QueryWithSQLResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryWithSQLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryWithSQLResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryWithSQLResponseBody) SetCode(v string) *QueryWithSQLResponseBody {
	s.Code = &v
	return s
}

func (s *QueryWithSQLResponseBody) SetData(v *QueryWithSQLResponseBodyData) *QueryWithSQLResponseBody {
	s.Data = v
	return s
}

func (s *QueryWithSQLResponseBody) SetMessage(v string) *QueryWithSQLResponseBody {
	s.Message = &v
	return s
}

func (s *QueryWithSQLResponseBody) SetRequestId(v string) *QueryWithSQLResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryWithSQLResponseBody) SetSuccess(v bool) *QueryWithSQLResponseBody {
	s.Success = &v
	return s
}

func (s *QueryWithSQLResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryWithSQLResponseBodyData struct {
	// The query result rows.
	Rows []*Row `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryWithSQLResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryWithSQLResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryWithSQLResponseBodyData) GetRows() []*Row {
	return s.Rows
}

func (s *QueryWithSQLResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *QueryWithSQLResponseBodyData) SetRows(v []*Row) *QueryWithSQLResponseBodyData {
	s.Rows = v
	return s
}

func (s *QueryWithSQLResponseBodyData) SetTotal(v int32) *QueryWithSQLResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryWithSQLResponseBodyData) Validate() error {
	if s.Rows != nil {
		for _, item := range s.Rows {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
