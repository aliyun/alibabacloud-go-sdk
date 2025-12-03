// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkitemEstimateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ListWorkitemEstimateResponseBody
	GetCode() *int64
	SetErrorCode(v string) *ListWorkitemEstimateResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListWorkitemEstimateResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *ListWorkitemEstimateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWorkitemEstimateResponseBody
	GetSuccess() *bool
	SetWorkitemTimeEstimate(v []*ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) *ListWorkitemEstimateResponseBody
	GetWorkitemTimeEstimate() []*ListWorkitemEstimateResponseBodyWorkitemTimeEstimate
}

type ListWorkitemEstimateResponseBody struct {
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
	Success              *bool                                                   `json:"success,omitempty" xml:"success,omitempty"`
	WorkitemTimeEstimate []*ListWorkitemEstimateResponseBodyWorkitemTimeEstimate `json:"workitemTimeEstimate,omitempty" xml:"workitemTimeEstimate,omitempty" type:"Repeated"`
}

func (s ListWorkitemEstimateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkitemEstimateResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkitemEstimateResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ListWorkitemEstimateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListWorkitemEstimateResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListWorkitemEstimateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkitemEstimateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkitemEstimateResponseBody) GetWorkitemTimeEstimate() []*ListWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	return s.WorkitemTimeEstimate
}

func (s *ListWorkitemEstimateResponseBody) SetCode(v int64) *ListWorkitemEstimateResponseBody {
	s.Code = &v
	return s
}

func (s *ListWorkitemEstimateResponseBody) SetErrorCode(v string) *ListWorkitemEstimateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListWorkitemEstimateResponseBody) SetErrorMsg(v string) *ListWorkitemEstimateResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListWorkitemEstimateResponseBody) SetRequestId(v string) *ListWorkitemEstimateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkitemEstimateResponseBody) SetSuccess(v bool) *ListWorkitemEstimateResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkitemEstimateResponseBody) SetWorkitemTimeEstimate(v []*ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) *ListWorkitemEstimateResponseBody {
	s.WorkitemTimeEstimate = v
	return s
}

func (s *ListWorkitemEstimateResponseBody) Validate() error {
	if s.WorkitemTimeEstimate != nil {
		for _, item := range s.WorkitemTimeEstimate {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkitemEstimateResponseBodyWorkitemTimeEstimate struct {
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
	Identifier *string                                                         `json:"identifier,omitempty" xml:"identifier,omitempty"`
	RecordUser *ListWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser `json:"recordUser,omitempty" xml:"recordUser,omitempty" type:"Struct"`
	// example:
	//
	// 8
	SpentTime *float32 `json:"spentTime,omitempty" xml:"spentTime,omitempty"`
	// example:
	//
	// 研发
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 5daa9a15c7fd55523996......
	WorkitemIdentifier *string `json:"workitemIdentifier,omitempty" xml:"workitemIdentifier,omitempty"`
}

func (s ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) String() string {
	return dara.Prettify(s)
}

func (s ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) GoString() string {
	return s.String()
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) GetDescription() *string {
	return s.Description
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) GetGmtEnd() *int64 {
	return s.GmtEnd
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) GetGmtStart() *int64 {
	return s.GmtStart
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) GetRecordUser() *ListWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	return s.RecordUser
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) GetSpentTime() *float32 {
	return s.SpentTime
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) GetType() *string {
	return s.Type
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) GetWorkitemIdentifier() *string {
	return s.WorkitemIdentifier
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetDescription(v string) *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.Description = &v
	return s
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetGmtCreate(v int64) *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.GmtCreate = &v
	return s
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetGmtEnd(v int64) *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.GmtEnd = &v
	return s
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetGmtModified(v int64) *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.GmtModified = &v
	return s
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetGmtStart(v int64) *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.GmtStart = &v
	return s
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetIdentifier(v string) *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.Identifier = &v
	return s
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetRecordUser(v *ListWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.RecordUser = v
	return s
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetSpentTime(v float32) *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.SpentTime = &v
	return s
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetType(v string) *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.Type = &v
	return s
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetWorkitemIdentifier(v string) *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.WorkitemIdentifier = &v
	return s
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimate) Validate() error {
	if s.RecordUser != nil {
		if err := s.RecordUser.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser struct {
	// example:
	//
	// 132xxxx123
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) String() string {
	return dara.Prettify(s)
}

func (s ListWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GoString() string {
	return s.String()
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GetName() *string {
	return s.Name
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetIdentifier(v string) *ListWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.Identifier = &v
	return s
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetName(v string) *ListWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.Name = &v
	return s
}

func (s *ListWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) Validate() error {
	return dara.Validate(s)
}
