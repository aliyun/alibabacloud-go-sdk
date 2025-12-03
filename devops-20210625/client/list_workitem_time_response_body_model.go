// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkitemTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ListWorkitemTimeResponseBody
	GetCode() *int64
	SetErrorCode(v string) *ListWorkitemTimeResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListWorkitemTimeResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *ListWorkitemTimeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWorkitemTimeResponseBody
	GetSuccess() *bool
	SetWorkitemTime(v []*ListWorkitemTimeResponseBodyWorkitemTime) *ListWorkitemTimeResponseBody
	GetWorkitemTime() []*ListWorkitemTimeResponseBodyWorkitemTime
}

type ListWorkitemTimeResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success      *bool                                       `json:"success,omitempty" xml:"success,omitempty"`
	WorkitemTime []*ListWorkitemTimeResponseBodyWorkitemTime `json:"workitemTime,omitempty" xml:"workitemTime,omitempty" type:"Repeated"`
}

func (s ListWorkitemTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkitemTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkitemTimeResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ListWorkitemTimeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListWorkitemTimeResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListWorkitemTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkitemTimeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkitemTimeResponseBody) GetWorkitemTime() []*ListWorkitemTimeResponseBodyWorkitemTime {
	return s.WorkitemTime
}

func (s *ListWorkitemTimeResponseBody) SetCode(v int64) *ListWorkitemTimeResponseBody {
	s.Code = &v
	return s
}

func (s *ListWorkitemTimeResponseBody) SetErrorCode(v string) *ListWorkitemTimeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListWorkitemTimeResponseBody) SetErrorMsg(v string) *ListWorkitemTimeResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListWorkitemTimeResponseBody) SetRequestId(v string) *ListWorkitemTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkitemTimeResponseBody) SetSuccess(v bool) *ListWorkitemTimeResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkitemTimeResponseBody) SetWorkitemTime(v []*ListWorkitemTimeResponseBodyWorkitemTime) *ListWorkitemTimeResponseBody {
	s.WorkitemTime = v
	return s
}

func (s *ListWorkitemTimeResponseBody) Validate() error {
	if s.WorkitemTime != nil {
		for _, item := range s.WorkitemTime {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkitemTimeResponseBodyWorkitemTime struct {
	// example:
	//
	// 1
	ActualTime *float32 `json:"actualTime,omitempty" xml:"actualTime,omitempty"`
	// example:
	//
	// 开发代码等
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1653235200000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1653235200000
	GmtEnd *int64 `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
	// example:
	//
	// 1653235200000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 1653235200000
	GmtStart *int64 `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
	// example:
	//
	// a4ac3a81e90217db32b7......
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// 1967043931......
	RecordUser *string `json:"recordUser,omitempty" xml:"recordUser,omitempty"`
	// example:
	//
	// 研发
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 5daa9a15c7fd55523996......
	WorkitemIdentifier *string `json:"workitemIdentifier,omitempty" xml:"workitemIdentifier,omitempty"`
}

func (s ListWorkitemTimeResponseBodyWorkitemTime) String() string {
	return dara.Prettify(s)
}

func (s ListWorkitemTimeResponseBodyWorkitemTime) GoString() string {
	return s.String()
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) GetActualTime() *float32 {
	return s.ActualTime
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) GetDescription() *string {
	return s.Description
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) GetGmtEnd() *int64 {
	return s.GmtEnd
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) GetGmtStart() *int64 {
	return s.GmtStart
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) GetRecordUser() *string {
	return s.RecordUser
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) GetType() *string {
	return s.Type
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) GetWorkitemIdentifier() *string {
	return s.WorkitemIdentifier
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) SetActualTime(v float32) *ListWorkitemTimeResponseBodyWorkitemTime {
	s.ActualTime = &v
	return s
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) SetDescription(v string) *ListWorkitemTimeResponseBodyWorkitemTime {
	s.Description = &v
	return s
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) SetGmtCreate(v int64) *ListWorkitemTimeResponseBodyWorkitemTime {
	s.GmtCreate = &v
	return s
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) SetGmtEnd(v int64) *ListWorkitemTimeResponseBodyWorkitemTime {
	s.GmtEnd = &v
	return s
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) SetGmtModified(v int64) *ListWorkitemTimeResponseBodyWorkitemTime {
	s.GmtModified = &v
	return s
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) SetGmtStart(v int64) *ListWorkitemTimeResponseBodyWorkitemTime {
	s.GmtStart = &v
	return s
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) SetIdentifier(v string) *ListWorkitemTimeResponseBodyWorkitemTime {
	s.Identifier = &v
	return s
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) SetRecordUser(v string) *ListWorkitemTimeResponseBodyWorkitemTime {
	s.RecordUser = &v
	return s
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) SetType(v string) *ListWorkitemTimeResponseBodyWorkitemTime {
	s.Type = &v
	return s
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) SetWorkitemIdentifier(v string) *ListWorkitemTimeResponseBodyWorkitemTime {
	s.WorkitemIdentifier = &v
	return s
}

func (s *ListWorkitemTimeResponseBodyWorkitemTime) Validate() error {
	return dara.Validate(s)
}
