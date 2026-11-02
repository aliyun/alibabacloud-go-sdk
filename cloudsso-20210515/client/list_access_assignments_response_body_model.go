// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessAssignmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessAssignments(v []*ListAccessAssignmentsResponseBodyAccessAssignments) *ListAccessAssignmentsResponseBody
	GetAccessAssignments() []*ListAccessAssignmentsResponseBodyAccessAssignments
	SetIsTruncated(v bool) *ListAccessAssignmentsResponseBody
	GetIsTruncated() *bool
	SetMaxResults(v int32) *ListAccessAssignmentsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAccessAssignmentsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAccessAssignmentsResponseBody
	GetRequestId() *string
	SetTotalCounts(v int32) *ListAccessAssignmentsResponseBody
	GetTotalCounts() *int32
}

type ListAccessAssignmentsResponseBody struct {
	// The list of access assignments.
	AccessAssignments []*ListAccessAssignmentsResponseBodyAccessAssignments `json:"AccessAssignments,omitempty" xml:"AccessAssignments,omitempty" type:"Repeated"`
	// Indicates whether the returned result is truncated. Valid values:
	//
	// - true: The result is truncated.
	//
	// - false: The result is not truncated.
	//
	// example:
	//
	// false
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The maximum number of entries returned on each page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page of query results.
	//
	// > This parameter is returned only when `IsTruncated` is set to `true`.
	//
	// example:
	//
	// K1c3o9K7pFxoTtxH1Nm7MMLb7zrDGvftYBQBPDVv7AD3a8yhRb3Mk8L9ivmN6bFSjfkZNTAg3h4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 66898413-EB80-556D-9429-06FE3548F672
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that meet the filter conditions specified in the request parameters.
	//
	// example:
	//
	// 1
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListAccessAssignmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccessAssignmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessAssignmentsResponseBody) GetAccessAssignments() []*ListAccessAssignmentsResponseBodyAccessAssignments {
	return s.AccessAssignments
}

func (s *ListAccessAssignmentsResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListAccessAssignmentsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAccessAssignmentsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAccessAssignmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccessAssignmentsResponseBody) GetTotalCounts() *int32 {
	return s.TotalCounts
}

func (s *ListAccessAssignmentsResponseBody) SetAccessAssignments(v []*ListAccessAssignmentsResponseBodyAccessAssignments) *ListAccessAssignmentsResponseBody {
	s.AccessAssignments = v
	return s
}

func (s *ListAccessAssignmentsResponseBody) SetIsTruncated(v bool) *ListAccessAssignmentsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListAccessAssignmentsResponseBody) SetMaxResults(v int32) *ListAccessAssignmentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAccessAssignmentsResponseBody) SetNextToken(v string) *ListAccessAssignmentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAccessAssignmentsResponseBody) SetRequestId(v string) *ListAccessAssignmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccessAssignmentsResponseBody) SetTotalCounts(v int32) *ListAccessAssignmentsResponseBody {
	s.TotalCounts = &v
	return s
}

func (s *ListAccessAssignmentsResponseBody) Validate() error {
	if s.AccessAssignments != nil {
		for _, item := range s.AccessAssignments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAccessAssignmentsResponseBodyAccessAssignments struct {
	// The ID of the access configuration.
	//
	// example:
	//
	// ac-00jhtfl8thteu6uj****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The name of the access configuration.
	//
	// example:
	//
	// ECS-Admin
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	// The time when the access assignment was successful. The time follows the RFC 3339 standard and is displayed in UTC in the YYYY-MM-DDTHH:mm:ssZ format.
	//
	// example:
	//
	// 2021-11-04T10:03:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the Alibaba Cloud account to which the authorization target belongs.
	//
	// example:
	//
	// 114240524784****
	OriginTargetId *string `json:"OriginTargetId,omitempty" xml:"OriginTargetId,omitempty"`
	// The ID of the CloudSSO identity.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The name of the CloudSSO identity.
	//
	// example:
	//
	// Alice
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the CloudSSO identity. Valid values:
	//
	// - User: user.
	//
	// - Group: group.
	//
	// example:
	//
	// User
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The ID of the task target.
	//
	// example:
	//
	// 114240524784****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the task target.
	//
	// example:
	//
	// dev-test
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The path ID of the task target in the resource directory.
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The path name of the task target in the resource directory.
	TargetPathName *string `json:"TargetPathName,omitempty" xml:"TargetPathName,omitempty"`
	// The type of the task target.
	//
	// Valid values:
	//
	// - RD-Account: The task target is an account in a resource directory.
	//
	// example:
	//
	// RD-Account
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListAccessAssignmentsResponseBodyAccessAssignments) String() string {
	return dara.Prettify(s)
}

func (s ListAccessAssignmentsResponseBodyAccessAssignments) GoString() string {
	return s.String()
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) GetAccessConfigurationName() *string {
	return s.AccessConfigurationName
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) GetOriginTargetId() *string {
	return s.OriginTargetId
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) GetTargetId() *string {
	return s.TargetId
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) GetTargetName() *string {
	return s.TargetName
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) GetTargetPath() *string {
	return s.TargetPath
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) GetTargetPathName() *string {
	return s.TargetPathName
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) GetTargetType() *string {
	return s.TargetType
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetAccessConfigurationId(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.AccessConfigurationId = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetAccessConfigurationName(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.AccessConfigurationName = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetCreateTime(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.CreateTime = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetOriginTargetId(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.OriginTargetId = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetPrincipalId(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.PrincipalId = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetPrincipalName(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.PrincipalName = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetPrincipalType(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.PrincipalType = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetTargetId(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.TargetId = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetTargetName(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.TargetName = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetTargetPath(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.TargetPath = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetTargetPathName(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.TargetPathName = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) SetTargetType(v string) *ListAccessAssignmentsResponseBodyAccessAssignments {
	s.TargetType = &v
	return s
}

func (s *ListAccessAssignmentsResponseBodyAccessAssignments) Validate() error {
	return dara.Validate(s)
}
