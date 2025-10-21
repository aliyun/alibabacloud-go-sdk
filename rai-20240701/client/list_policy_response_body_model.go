// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPolicyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListPolicyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListPolicyResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListPolicyResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPolicyResponseBody
	GetPageSize() *int32
	SetPolicyInfoList(v []*ListPolicyResponseBodyPolicyInfoList) *ListPolicyResponseBody
	GetPolicyInfoList() []*ListPolicyResponseBodyPolicyInfoList
	SetRequestId(v string) *ListPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPolicyResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListPolicyResponseBody
	GetTotalCount() *int32
}

type ListPolicyResponseBody struct {
	// Status code, 00000 indicates success; other values indicate failure.
	//
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code description.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// If there is an error, returns the error message.
	//
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Page number, consistent with the PageNumber in the request.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, the maximum number of results returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// List of policy objects.
	PolicyInfoList []*ListPolicyResponseBodyPolicyInfoList `json:"PolicyInfoList,omitempty" xml:"PolicyInfoList,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Possible values are:
	//
	// - True: Call succeeded.
	//
	// - False: Call failed.
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
}

func (s ListPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPolicyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPolicyResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPolicyResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPolicyResponseBody) GetPolicyInfoList() []*ListPolicyResponseBodyPolicyInfoList {
	return s.PolicyInfoList
}

func (s *ListPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPolicyResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPolicyResponseBody) SetCode(v string) *ListPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *ListPolicyResponseBody) SetHttpStatusCode(v int32) *ListPolicyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPolicyResponseBody) SetMessage(v string) *ListPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *ListPolicyResponseBody) SetPageNumber(v int32) *ListPolicyResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPolicyResponseBody) SetPageSize(v int32) *ListPolicyResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPolicyResponseBody) SetPolicyInfoList(v []*ListPolicyResponseBodyPolicyInfoList) *ListPolicyResponseBody {
	s.PolicyInfoList = v
	return s
}

func (s *ListPolicyResponseBody) SetRequestId(v string) *ListPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPolicyResponseBody) SetSuccess(v bool) *ListPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *ListPolicyResponseBody) SetTotalCount(v int32) *ListPolicyResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPolicyResponseBody) Validate() error {
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

type ListPolicyResponseBodyPolicyInfoList struct {
	// Modification time.
	//
	// example:
	//
	// 1731204769000
	GmtModified     *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	IsSidecarPolicy *int32 `json:"IsSidecarPolicy,omitempty" xml:"IsSidecarPolicy,omitempty"`
	// Tag policy ID.
	//
	// example:
	//
	// 16
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Policy identifier.
	//
	// example:
	//
	// x1bc5xgs4uhx
	PolicyIdentifier *string `json:"PolicyIdentifier,omitempty" xml:"PolicyIdentifier,omitempty"`
	// Permission policy name.
	//
	// example:
	//
	// testPolicy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	SceneType  *int32  `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
}

func (s ListPolicyResponseBodyPolicyInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyResponseBodyPolicyInfoList) GoString() string {
	return s.String()
}

func (s *ListPolicyResponseBodyPolicyInfoList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListPolicyResponseBodyPolicyInfoList) GetIsSidecarPolicy() *int32 {
	return s.IsSidecarPolicy
}

func (s *ListPolicyResponseBodyPolicyInfoList) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *ListPolicyResponseBodyPolicyInfoList) GetPolicyIdentifier() *string {
	return s.PolicyIdentifier
}

func (s *ListPolicyResponseBodyPolicyInfoList) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListPolicyResponseBodyPolicyInfoList) GetSceneType() *int32 {
	return s.SceneType
}

func (s *ListPolicyResponseBodyPolicyInfoList) SetGmtModified(v int64) *ListPolicyResponseBodyPolicyInfoList {
	s.GmtModified = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyInfoList) SetIsSidecarPolicy(v int32) *ListPolicyResponseBodyPolicyInfoList {
	s.IsSidecarPolicy = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyInfoList) SetPolicyId(v int64) *ListPolicyResponseBodyPolicyInfoList {
	s.PolicyId = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyInfoList) SetPolicyIdentifier(v string) *ListPolicyResponseBodyPolicyInfoList {
	s.PolicyIdentifier = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyInfoList) SetPolicyName(v string) *ListPolicyResponseBodyPolicyInfoList {
	s.PolicyName = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyInfoList) SetSceneType(v int32) *ListPolicyResponseBodyPolicyInfoList {
	s.SceneType = &v
	return s
}

func (s *ListPolicyResponseBodyPolicyInfoList) Validate() error {
	return dara.Validate(s)
}
