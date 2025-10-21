// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWordGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListWordGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListWordGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListWordGroupResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListWordGroupResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWordGroupResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListWordGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWordGroupResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListWordGroupResponseBody
	GetTotalCount() *int32
	SetWordGroupInfoList(v []*ListWordGroupResponseBodyWordGroupInfoList) *ListWordGroupResponseBody
	GetWordGroupInfoList() []*ListWordGroupResponseBodyWordGroupInfoList
}

type ListWordGroupResponseBody struct {
	// Status code, 00000 indicates success; other values indicate failure.
	//
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// If an error occurs, returns the error message.
	//
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, the maximum number of results returned per page.
	//
	// Maximum limit: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// BE2558EC-A9EA-5276-ADB5-107B09CF3D11
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful: true means the call was successful; false means the call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total count.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// List of keyword group objects.
	WordGroupInfoList []*ListWordGroupResponseBodyWordGroupInfoList `json:"WordGroupInfoList,omitempty" xml:"WordGroupInfoList,omitempty" type:"Repeated"`
}

func (s ListWordGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWordGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListWordGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListWordGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListWordGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWordGroupResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWordGroupResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWordGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWordGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWordGroupResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListWordGroupResponseBody) GetWordGroupInfoList() []*ListWordGroupResponseBodyWordGroupInfoList {
	return s.WordGroupInfoList
}

func (s *ListWordGroupResponseBody) SetCode(v string) *ListWordGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListWordGroupResponseBody) SetHttpStatusCode(v int32) *ListWordGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListWordGroupResponseBody) SetMessage(v string) *ListWordGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListWordGroupResponseBody) SetPageNumber(v int32) *ListWordGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListWordGroupResponseBody) SetPageSize(v int32) *ListWordGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListWordGroupResponseBody) SetRequestId(v string) *ListWordGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWordGroupResponseBody) SetSuccess(v bool) *ListWordGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ListWordGroupResponseBody) SetTotalCount(v int32) *ListWordGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWordGroupResponseBody) SetWordGroupInfoList(v []*ListWordGroupResponseBodyWordGroupInfoList) *ListWordGroupResponseBody {
	s.WordGroupInfoList = v
	return s
}

func (s *ListWordGroupResponseBody) Validate() error {
	if s.WordGroupInfoList != nil {
		for _, item := range s.WordGroupInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWordGroupResponseBodyWordGroupInfoList struct {
	// Policy modification time.
	//
	// example:
	//
	// 1673578650000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Keyword group ID.
	//
	// example:
	//
	// 16
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Keyword group name.
	//
	// example:
	//
	// testGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// List of associated policy objects.
	PolicyInfoList []*ListWordGroupResponseBodyWordGroupInfoListPolicyInfoList `json:"PolicyInfoList,omitempty" xml:"PolicyInfoList,omitempty" type:"Repeated"`
}

func (s ListWordGroupResponseBodyWordGroupInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListWordGroupResponseBodyWordGroupInfoList) GoString() string {
	return s.String()
}

func (s *ListWordGroupResponseBodyWordGroupInfoList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListWordGroupResponseBodyWordGroupInfoList) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListWordGroupResponseBodyWordGroupInfoList) GetGroupName() *string {
	return s.GroupName
}

func (s *ListWordGroupResponseBodyWordGroupInfoList) GetPolicyInfoList() []*ListWordGroupResponseBodyWordGroupInfoListPolicyInfoList {
	return s.PolicyInfoList
}

func (s *ListWordGroupResponseBodyWordGroupInfoList) SetGmtModified(v int64) *ListWordGroupResponseBodyWordGroupInfoList {
	s.GmtModified = &v
	return s
}

func (s *ListWordGroupResponseBodyWordGroupInfoList) SetGroupId(v int64) *ListWordGroupResponseBodyWordGroupInfoList {
	s.GroupId = &v
	return s
}

func (s *ListWordGroupResponseBodyWordGroupInfoList) SetGroupName(v string) *ListWordGroupResponseBodyWordGroupInfoList {
	s.GroupName = &v
	return s
}

func (s *ListWordGroupResponseBodyWordGroupInfoList) SetPolicyInfoList(v []*ListWordGroupResponseBodyWordGroupInfoListPolicyInfoList) *ListWordGroupResponseBodyWordGroupInfoList {
	s.PolicyInfoList = v
	return s
}

func (s *ListWordGroupResponseBodyWordGroupInfoList) Validate() error {
	if s.PolicyInfoList != nil {
		for _, item := range s.PolicyInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWordGroupResponseBodyWordGroupInfoListPolicyInfoList struct {
	// Detection policy ID.
	//
	// example:
	//
	// 16
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Policy identifier.
	//
	// example:
	//
	// mai934jhuekf
	PolicyIdentifier *string `json:"PolicyIdentifier,omitempty" xml:"PolicyIdentifier,omitempty"`
	// Detection policy name.
	//
	// example:
	//
	// testPoliy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s ListWordGroupResponseBodyWordGroupInfoListPolicyInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListWordGroupResponseBodyWordGroupInfoListPolicyInfoList) GoString() string {
	return s.String()
}

func (s *ListWordGroupResponseBodyWordGroupInfoListPolicyInfoList) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *ListWordGroupResponseBodyWordGroupInfoListPolicyInfoList) GetPolicyIdentifier() *string {
	return s.PolicyIdentifier
}

func (s *ListWordGroupResponseBodyWordGroupInfoListPolicyInfoList) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListWordGroupResponseBodyWordGroupInfoListPolicyInfoList) SetPolicyId(v int64) *ListWordGroupResponseBodyWordGroupInfoListPolicyInfoList {
	s.PolicyId = &v
	return s
}

func (s *ListWordGroupResponseBodyWordGroupInfoListPolicyInfoList) SetPolicyIdentifier(v string) *ListWordGroupResponseBodyWordGroupInfoListPolicyInfoList {
	s.PolicyIdentifier = &v
	return s
}

func (s *ListWordGroupResponseBodyWordGroupInfoListPolicyInfoList) SetPolicyName(v string) *ListWordGroupResponseBodyWordGroupInfoListPolicyInfoList {
	s.PolicyName = &v
	return s
}

func (s *ListWordGroupResponseBodyWordGroupInfoListPolicyInfoList) Validate() error {
	return dara.Validate(s)
}
