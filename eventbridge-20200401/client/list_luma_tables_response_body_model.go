// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLumaTablesResponseBody
	GetCode() *string
	SetData(v *ListLumaTablesResponseBodyData) *ListLumaTablesResponseBody
	GetData() *ListLumaTablesResponseBodyData
	SetMessage(v string) *ListLumaTablesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListLumaTablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLumaTablesResponseBody
	GetSuccess() *bool
}

type ListLumaTablesResponseBody struct {
	// The response code of the operation. A value of Success indicates success. An error code is returned if the call fails.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of event tables bound to the Agent. All results are returned at once without pagination.
	Data *ListLumaTablesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned by the operation. The value is Operation success if the call succeeds, or a specific error description if the call fails.
	//
	// example:
	//
	// Operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, used for troubleshooting and ticket feedback.
	//
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. A value of true indicates success.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLumaTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLumaTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLumaTablesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLumaTablesResponseBody) GetData() *ListLumaTablesResponseBodyData {
	return s.Data
}

func (s *ListLumaTablesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLumaTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLumaTablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLumaTablesResponseBody) SetCode(v string) *ListLumaTablesResponseBody {
	s.Code = &v
	return s
}

func (s *ListLumaTablesResponseBody) SetData(v *ListLumaTablesResponseBodyData) *ListLumaTablesResponseBody {
	s.Data = v
	return s
}

func (s *ListLumaTablesResponseBody) SetMessage(v string) *ListLumaTablesResponseBody {
	s.Message = &v
	return s
}

func (s *ListLumaTablesResponseBody) SetRequestId(v string) *ListLumaTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLumaTablesResponseBody) SetSuccess(v bool) *ListLumaTablesResponseBody {
	s.Success = &v
	return s
}

func (s *ListLumaTablesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLumaTablesResponseBodyData struct {
	// The list of event tables bound to the Agent.
	//
	// example:
	//
	// [{"Name":"my_table","Namespace":"my_namespace"}]
	Tables []*LumaTable `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s ListLumaTablesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLumaTablesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLumaTablesResponseBodyData) GetTables() []*LumaTable {
	return s.Tables
}

func (s *ListLumaTablesResponseBodyData) SetTables(v []*LumaTable) *ListLumaTablesResponseBodyData {
	s.Tables = v
	return s
}

func (s *ListLumaTablesResponseBodyData) Validate() error {
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
