// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ListTablesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListTablesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTablesResponseBody
	GetRequestId() *string
	SetStatus(v string) *ListTablesResponseBody
	GetStatus() *string
	SetTables(v *ListTablesResponseBodyTables) *ListTablesResponseBody
	GetTables() *ListTablesResponseBodyTables
}

type ListTablesResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the operation. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The queried tables.
	Tables *ListTablesResponseBodyTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Struct"`
}

func (s ListTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTablesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTablesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListTablesResponseBody) GetTables() *ListTablesResponseBodyTables {
	return s.Tables
}

func (s *ListTablesResponseBody) SetMessage(v string) *ListTablesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTablesResponseBody) SetNextToken(v string) *ListTablesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTablesResponseBody) SetRequestId(v string) *ListTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTablesResponseBody) SetStatus(v string) *ListTablesResponseBody {
	s.Status = &v
	return s
}

func (s *ListTablesResponseBody) SetTables(v *ListTablesResponseBodyTables) *ListTablesResponseBody {
	s.Tables = v
	return s
}

func (s *ListTablesResponseBody) Validate() error {
	if s.Tables != nil {
		if err := s.Tables.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTablesResponseBodyTables struct {
	Tables []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s ListTablesResponseBodyTables) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBodyTables) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyTables) GetTables() []*string {
	return s.Tables
}

func (s *ListTablesResponseBodyTables) SetTables(v []*string) *ListTablesResponseBodyTables {
	s.Tables = v
	return s
}

func (s *ListTablesResponseBodyTables) Validate() error {
	return dara.Validate(s)
}
