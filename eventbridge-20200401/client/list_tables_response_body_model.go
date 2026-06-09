// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTablesResponseBody
	GetCode() *string
	SetData(v *ListTablesResponseBodyData) *ListTablesResponseBody
	GetData() *ListTablesResponseBodyData
	SetMessage(v string) *ListTablesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTablesResponseBody
	GetSuccess() *bool
}

type ListTablesResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"Tables":[{"Name":"my_table","Comment":"测试事件表"}],"NextToken":"10","Total":1}
	Data *ListTablesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTablesResponseBody) GetData() *ListTablesResponseBodyData {
	return s.Data
}

func (s *ListTablesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTablesResponseBody) SetCode(v string) *ListTablesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTablesResponseBody) SetData(v *ListTablesResponseBodyData) *ListTablesResponseBody {
	s.Data = v
	return s
}

func (s *ListTablesResponseBody) SetMessage(v string) *ListTablesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTablesResponseBody) SetRequestId(v string) *ListTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTablesResponseBody) SetSuccess(v bool) *ListTablesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTablesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTablesResponseBodyData struct {
	// example:
	//
	// 10
	NextToken *string  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Tables    []*Table `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListTablesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTablesResponseBodyData) GetTables() []*Table {
	return s.Tables
}

func (s *ListTablesResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListTablesResponseBodyData) SetNextToken(v string) *ListTablesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListTablesResponseBodyData) SetTables(v []*Table) *ListTablesResponseBodyData {
	s.Tables = v
	return s
}

func (s *ListTablesResponseBodyData) SetTotal(v int32) *ListTablesResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListTablesResponseBodyData) Validate() error {
	if s.Tables != nil {
		for _, item := range s.Tables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
