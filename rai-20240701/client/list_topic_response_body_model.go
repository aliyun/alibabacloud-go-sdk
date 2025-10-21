// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTopicResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListTopicResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListTopicResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListTopicResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTopicResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTopicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTopicResponseBody
	GetSuccess() *bool
	SetTopicInfoList(v []*ListTopicResponseBodyTopicInfoList) *ListTopicResponseBody
	GetTopicInfoList() []*ListTopicResponseBodyTopicInfoList
	SetTotalCount(v int32) *ListTopicResponseBody
	GetTotalCount() *int32
}

type ListTopicResponseBody struct {
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success       *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TopicInfoList []*ListTopicResponseBodyTopicInfoList `json:"TopicInfoList,omitempty" xml:"TopicInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTopicResponseBody) GoString() string {
	return s.String()
}

func (s *ListTopicResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTopicResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTopicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTopicResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTopicResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTopicResponseBody) GetTopicInfoList() []*ListTopicResponseBodyTopicInfoList {
	return s.TopicInfoList
}

func (s *ListTopicResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTopicResponseBody) SetCode(v string) *ListTopicResponseBody {
	s.Code = &v
	return s
}

func (s *ListTopicResponseBody) SetHttpStatusCode(v int32) *ListTopicResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTopicResponseBody) SetMessage(v string) *ListTopicResponseBody {
	s.Message = &v
	return s
}

func (s *ListTopicResponseBody) SetPageNumber(v int32) *ListTopicResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTopicResponseBody) SetPageSize(v int32) *ListTopicResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTopicResponseBody) SetRequestId(v string) *ListTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTopicResponseBody) SetSuccess(v bool) *ListTopicResponseBody {
	s.Success = &v
	return s
}

func (s *ListTopicResponseBody) SetTopicInfoList(v []*ListTopicResponseBodyTopicInfoList) *ListTopicResponseBody {
	s.TopicInfoList = v
	return s
}

func (s *ListTopicResponseBody) SetTotalCount(v int32) *ListTopicResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTopicResponseBody) Validate() error {
	if s.TopicInfoList != nil {
		for _, item := range s.TopicInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTopicResponseBodyTopicInfoList struct {
	// example:
	//
	// 1597738932000
	GmtModified     *int64                                              `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	PolicyInfoList  []*ListTopicResponseBodyTopicInfoListPolicyInfoList `json:"PolicyInfoList,omitempty" xml:"PolicyInfoList,omitempty" type:"Repeated"`
	TopicDefinition *string                                             `json:"TopicDefinition,omitempty" xml:"TopicDefinition,omitempty"`
	// example:
	//
	// 216
	TopicId   *int64  `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s ListTopicResponseBodyTopicInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListTopicResponseBodyTopicInfoList) GoString() string {
	return s.String()
}

func (s *ListTopicResponseBodyTopicInfoList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListTopicResponseBodyTopicInfoList) GetPolicyInfoList() []*ListTopicResponseBodyTopicInfoListPolicyInfoList {
	return s.PolicyInfoList
}

func (s *ListTopicResponseBodyTopicInfoList) GetTopicDefinition() *string {
	return s.TopicDefinition
}

func (s *ListTopicResponseBodyTopicInfoList) GetTopicId() *int64 {
	return s.TopicId
}

func (s *ListTopicResponseBodyTopicInfoList) GetTopicName() *string {
	return s.TopicName
}

func (s *ListTopicResponseBodyTopicInfoList) SetGmtModified(v int64) *ListTopicResponseBodyTopicInfoList {
	s.GmtModified = &v
	return s
}

func (s *ListTopicResponseBodyTopicInfoList) SetPolicyInfoList(v []*ListTopicResponseBodyTopicInfoListPolicyInfoList) *ListTopicResponseBodyTopicInfoList {
	s.PolicyInfoList = v
	return s
}

func (s *ListTopicResponseBodyTopicInfoList) SetTopicDefinition(v string) *ListTopicResponseBodyTopicInfoList {
	s.TopicDefinition = &v
	return s
}

func (s *ListTopicResponseBodyTopicInfoList) SetTopicId(v int64) *ListTopicResponseBodyTopicInfoList {
	s.TopicId = &v
	return s
}

func (s *ListTopicResponseBodyTopicInfoList) SetTopicName(v string) *ListTopicResponseBodyTopicInfoList {
	s.TopicName = &v
	return s
}

func (s *ListTopicResponseBodyTopicInfoList) Validate() error {
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

type ListTopicResponseBodyTopicInfoListPolicyInfoList struct {
	// example:
	//
	// 412
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// bqi1c3s99qx3
	PolicyIdentifier *string `json:"PolicyIdentifier,omitempty" xml:"PolicyIdentifier,omitempty"`
	PolicyName       *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s ListTopicResponseBodyTopicInfoListPolicyInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListTopicResponseBodyTopicInfoListPolicyInfoList) GoString() string {
	return s.String()
}

func (s *ListTopicResponseBodyTopicInfoListPolicyInfoList) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *ListTopicResponseBodyTopicInfoListPolicyInfoList) GetPolicyIdentifier() *string {
	return s.PolicyIdentifier
}

func (s *ListTopicResponseBodyTopicInfoListPolicyInfoList) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListTopicResponseBodyTopicInfoListPolicyInfoList) SetPolicyId(v int64) *ListTopicResponseBodyTopicInfoListPolicyInfoList {
	s.PolicyId = &v
	return s
}

func (s *ListTopicResponseBodyTopicInfoListPolicyInfoList) SetPolicyIdentifier(v string) *ListTopicResponseBodyTopicInfoListPolicyInfoList {
	s.PolicyIdentifier = &v
	return s
}

func (s *ListTopicResponseBodyTopicInfoListPolicyInfoList) SetPolicyName(v string) *ListTopicResponseBodyTopicInfoListPolicyInfoList {
	s.PolicyName = &v
	return s
}

func (s *ListTopicResponseBodyTopicInfoListPolicyInfoList) Validate() error {
	return dara.Validate(s)
}
