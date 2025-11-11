// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuditContentErrorTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAuditContentErrorTypesResponseBody
	GetCode() *string
	SetData(v []*ListAuditContentErrorTypesResponseBodyData) *ListAuditContentErrorTypesResponseBody
	GetData() []*ListAuditContentErrorTypesResponseBodyData
	SetHttpStatusCode(v int32) *ListAuditContentErrorTypesResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListAuditContentErrorTypesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListAuditContentErrorTypesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListAuditContentErrorTypesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAuditContentErrorTypesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAuditContentErrorTypesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListAuditContentErrorTypesResponseBody
	GetTotalCount() *int32
}

type ListAuditContentErrorTypesResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListAuditContentErrorTypesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAuditContentErrorTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuditContentErrorTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuditContentErrorTypesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAuditContentErrorTypesResponseBody) GetData() []*ListAuditContentErrorTypesResponseBodyData {
	return s.Data
}

func (s *ListAuditContentErrorTypesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAuditContentErrorTypesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAuditContentErrorTypesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAuditContentErrorTypesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAuditContentErrorTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuditContentErrorTypesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAuditContentErrorTypesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAuditContentErrorTypesResponseBody) SetCode(v string) *ListAuditContentErrorTypesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAuditContentErrorTypesResponseBody) SetData(v []*ListAuditContentErrorTypesResponseBodyData) *ListAuditContentErrorTypesResponseBody {
	s.Data = v
	return s
}

func (s *ListAuditContentErrorTypesResponseBody) SetHttpStatusCode(v int32) *ListAuditContentErrorTypesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAuditContentErrorTypesResponseBody) SetMaxResults(v int32) *ListAuditContentErrorTypesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAuditContentErrorTypesResponseBody) SetMessage(v string) *ListAuditContentErrorTypesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAuditContentErrorTypesResponseBody) SetNextToken(v string) *ListAuditContentErrorTypesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAuditContentErrorTypesResponseBody) SetRequestId(v string) *ListAuditContentErrorTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuditContentErrorTypesResponseBody) SetSuccess(v bool) *ListAuditContentErrorTypesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAuditContentErrorTypesResponseBody) SetTotalCount(v int32) *ListAuditContentErrorTypesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAuditContentErrorTypesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAuditContentErrorTypesResponseBodyData struct {
	// example:
	//
	// ContentAccuracy
	MajorClassCode *string                                                 `json:"MajorClassCode,omitempty" xml:"MajorClassCode,omitempty"`
	MajorClassName *string                                                 `json:"MajorClassName,omitempty" xml:"MajorClassName,omitempty"`
	SubClasses     []*ListAuditContentErrorTypesResponseBodyDataSubClasses `json:"SubClasses,omitempty" xml:"SubClasses,omitempty" type:"Repeated"`
}

func (s ListAuditContentErrorTypesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAuditContentErrorTypesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAuditContentErrorTypesResponseBodyData) GetMajorClassCode() *string {
	return s.MajorClassCode
}

func (s *ListAuditContentErrorTypesResponseBodyData) GetMajorClassName() *string {
	return s.MajorClassName
}

func (s *ListAuditContentErrorTypesResponseBodyData) GetSubClasses() []*ListAuditContentErrorTypesResponseBodyDataSubClasses {
	return s.SubClasses
}

func (s *ListAuditContentErrorTypesResponseBodyData) SetMajorClassCode(v string) *ListAuditContentErrorTypesResponseBodyData {
	s.MajorClassCode = &v
	return s
}

func (s *ListAuditContentErrorTypesResponseBodyData) SetMajorClassName(v string) *ListAuditContentErrorTypesResponseBodyData {
	s.MajorClassName = &v
	return s
}

func (s *ListAuditContentErrorTypesResponseBodyData) SetSubClasses(v []*ListAuditContentErrorTypesResponseBodyDataSubClasses) *ListAuditContentErrorTypesResponseBodyData {
	s.SubClasses = v
	return s
}

func (s *ListAuditContentErrorTypesResponseBodyData) Validate() error {
	if s.SubClasses != nil {
		for _, item := range s.SubClasses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAuditContentErrorTypesResponseBodyDataSubClasses struct {
	// example:
	//
	// PunctuationError
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
}

func (s ListAuditContentErrorTypesResponseBodyDataSubClasses) String() string {
	return dara.Prettify(s)
}

func (s ListAuditContentErrorTypesResponseBodyDataSubClasses) GoString() string {
	return s.String()
}

func (s *ListAuditContentErrorTypesResponseBodyDataSubClasses) GetClassCode() *string {
	return s.ClassCode
}

func (s *ListAuditContentErrorTypesResponseBodyDataSubClasses) GetClassName() *string {
	return s.ClassName
}

func (s *ListAuditContentErrorTypesResponseBodyDataSubClasses) SetClassCode(v string) *ListAuditContentErrorTypesResponseBodyDataSubClasses {
	s.ClassCode = &v
	return s
}

func (s *ListAuditContentErrorTypesResponseBodyDataSubClasses) SetClassName(v string) *ListAuditContentErrorTypesResponseBodyDataSubClasses {
	s.ClassName = &v
	return s
}

func (s *ListAuditContentErrorTypesResponseBodyDataSubClasses) Validate() error {
	return dara.Validate(s)
}
