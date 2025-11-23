// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCorrectPreCheckDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListDataCorrectPreCheckDBResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataCorrectPreCheckDBResponseBody
	GetErrorMessage() *string
	SetPreCheckDBList(v []*ListDataCorrectPreCheckDBResponseBodyPreCheckDBList) *ListDataCorrectPreCheckDBResponseBody
	GetPreCheckDBList() []*ListDataCorrectPreCheckDBResponseBodyPreCheckDBList
	SetRequestId(v string) *ListDataCorrectPreCheckDBResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataCorrectPreCheckDBResponseBody
	GetSuccess() *bool
}

type ListDataCorrectPreCheckDBResponseBody struct {
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
	// The information about the databases that are involved in the precheck.
	PreCheckDBList []*ListDataCorrectPreCheckDBResponseBodyPreCheckDBList `json:"PreCheckDBList,omitempty" xml:"PreCheckDBList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 95A972AF-FAED-4768-9360-7C0DF5D594D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataCorrectPreCheckDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataCorrectPreCheckDBResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataCorrectPreCheckDBResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataCorrectPreCheckDBResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataCorrectPreCheckDBResponseBody) GetPreCheckDBList() []*ListDataCorrectPreCheckDBResponseBodyPreCheckDBList {
	return s.PreCheckDBList
}

func (s *ListDataCorrectPreCheckDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataCorrectPreCheckDBResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataCorrectPreCheckDBResponseBody) SetErrorCode(v string) *ListDataCorrectPreCheckDBResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataCorrectPreCheckDBResponseBody) SetErrorMessage(v string) *ListDataCorrectPreCheckDBResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataCorrectPreCheckDBResponseBody) SetPreCheckDBList(v []*ListDataCorrectPreCheckDBResponseBodyPreCheckDBList) *ListDataCorrectPreCheckDBResponseBody {
	s.PreCheckDBList = v
	return s
}

func (s *ListDataCorrectPreCheckDBResponseBody) SetRequestId(v string) *ListDataCorrectPreCheckDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataCorrectPreCheckDBResponseBody) SetSuccess(v bool) *ListDataCorrectPreCheckDBResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataCorrectPreCheckDBResponseBody) Validate() error {
	if s.PreCheckDBList != nil {
		for _, item := range s.PreCheckDBList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataCorrectPreCheckDBResponseBodyPreCheckDBList struct {
	// The ID of the database.
	//
	// example:
	//
	// 43***
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// test@localhost:3306
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	// The number of SQL statements.
	//
	// example:
	//
	// 1
	SqlNum *int64 `json:"SqlNum,omitempty" xml:"SqlNum,omitempty"`
}

func (s ListDataCorrectPreCheckDBResponseBodyPreCheckDBList) String() string {
	return dara.Prettify(s)
}

func (s ListDataCorrectPreCheckDBResponseBodyPreCheckDBList) GoString() string {
	return s.String()
}

func (s *ListDataCorrectPreCheckDBResponseBodyPreCheckDBList) GetDbId() *int64 {
	return s.DbId
}

func (s *ListDataCorrectPreCheckDBResponseBodyPreCheckDBList) GetSearchName() *string {
	return s.SearchName
}

func (s *ListDataCorrectPreCheckDBResponseBodyPreCheckDBList) GetSqlNum() *int64 {
	return s.SqlNum
}

func (s *ListDataCorrectPreCheckDBResponseBodyPreCheckDBList) SetDbId(v int64) *ListDataCorrectPreCheckDBResponseBodyPreCheckDBList {
	s.DbId = &v
	return s
}

func (s *ListDataCorrectPreCheckDBResponseBodyPreCheckDBList) SetSearchName(v string) *ListDataCorrectPreCheckDBResponseBodyPreCheckDBList {
	s.SearchName = &v
	return s
}

func (s *ListDataCorrectPreCheckDBResponseBodyPreCheckDBList) SetSqlNum(v int64) *ListDataCorrectPreCheckDBResponseBodyPreCheckDBList {
	s.SqlNum = &v
	return s
}

func (s *ListDataCorrectPreCheckDBResponseBodyPreCheckDBList) Validate() error {
	return dara.Validate(s)
}
