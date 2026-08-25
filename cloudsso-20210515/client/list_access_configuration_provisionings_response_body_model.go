// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessConfigurationProvisioningsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessConfigurationProvisionings(v []*ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) *ListAccessConfigurationProvisioningsResponseBody
	GetAccessConfigurationProvisionings() []*ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings
	SetIsTruncated(v bool) *ListAccessConfigurationProvisioningsResponseBody
	GetIsTruncated() *bool
	SetMaxResults(v int32) *ListAccessConfigurationProvisioningsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAccessConfigurationProvisioningsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAccessConfigurationProvisioningsResponseBody
	GetRequestId() *string
	SetTotalCounts(v int32) *ListAccessConfigurationProvisioningsResponseBody
	GetTotalCounts() *int32
}

type ListAccessConfigurationProvisioningsResponseBody struct {
	// The accounts for which the access configuration is provisioned.
	AccessConfigurationProvisionings []*ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings `json:"AccessConfigurationProvisionings,omitempty" xml:"AccessConfigurationProvisionings,omitempty" type:"Repeated"`
	// Indicates whether the queried entries are truncated. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// > This parameter is returned only when the value of `IsTruncated` is `true`.
	//
	// example:
	//
	// K1c3o9K7pFxoTtxH1Nm7MMLb7zrDGvftYBQBPDVv7AD3a8yhRb3Mk8L9ivmN6bFSjfkZNTAg3h4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6BA1BDF1-D845-5D2C-B742-74BE2970E4C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListAccessConfigurationProvisioningsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccessConfigurationProvisioningsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessConfigurationProvisioningsResponseBody) GetAccessConfigurationProvisionings() []*ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	return s.AccessConfigurationProvisionings
}

func (s *ListAccessConfigurationProvisioningsResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListAccessConfigurationProvisioningsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAccessConfigurationProvisioningsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAccessConfigurationProvisioningsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccessConfigurationProvisioningsResponseBody) GetTotalCounts() *int32 {
	return s.TotalCounts
}

func (s *ListAccessConfigurationProvisioningsResponseBody) SetAccessConfigurationProvisionings(v []*ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) *ListAccessConfigurationProvisioningsResponseBody {
	s.AccessConfigurationProvisionings = v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBody) SetIsTruncated(v bool) *ListAccessConfigurationProvisioningsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBody) SetMaxResults(v int32) *ListAccessConfigurationProvisioningsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBody) SetNextToken(v string) *ListAccessConfigurationProvisioningsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBody) SetRequestId(v string) *ListAccessConfigurationProvisioningsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBody) SetTotalCounts(v int32) *ListAccessConfigurationProvisioningsResponseBody {
	s.TotalCounts = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBody) Validate() error {
	if s.AccessConfigurationProvisionings != nil {
		for _, item := range s.AccessConfigurationProvisionings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings struct {
	// The ID of the access configuration.
	//
	// example:
	//
	// ac-00ccule7tadaijxc****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The name of the access configuration.
	//
	// example:
	//
	// VPC-Admin
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	// The first time when the access configuration was provisioned.
	//
	// example:
	//
	// 2021-07-26T08:54:14Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 114240524784****
	OriginTargetId *string `json:"OriginTargetId,omitempty" xml:"OriginTargetId,omitempty"`
	// The name of the custom policy that is created for an account in the resource directory.
	RAMPolicyNames []*string `json:"RAMPolicyNames,omitempty" xml:"RAMPolicyNames,omitempty" type:"Repeated"`
	// The name of the RAM role that is created for an account in the resource directory.
	//
	// example:
	//
	// AliyunReservedSSO-VPC-Admin
	RAMRoleName *string `json:"RAMRoleName,omitempty" xml:"RAMRoleName,omitempty"`
	// The name of the Security Assertion Markup Language (SAML) identity provider (IdP) that is created within an account in the resource directory.
	//
	// example:
	//
	// AliyunReservedSSO-d-00fc2p61****
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty"`
	// The status of the access configuration. Valid values:
	//
	// - Provisioned: The access configuration is provisioned.
	//
	// - ReprovisionRequired: The access configuration needs to be re-provisioned.
	//
	// - DeprovisionFailed: The access configuration failed to be provisioned.
	//
	// example:
	//
	// Provisioned
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task object.
	//
	// If the value of TargetType is `RD-Account`, the value of this parameter is the UID of an account in the resource directory.
	//
	// example:
	//
	// 101522521960****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the task object.
	//
	// example:
	//
	// SharedServices_5009****
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The path ID of the task object in the resource directory.
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The path name of the task object in the resource directory.
	TargetPathName *string `json:"TargetPathName,omitempty" xml:"TargetPathName,omitempty"`
	// The type of the task object.
	//
	// Set the value to RD-Account, which specifies the accounts in the resource directory.
	//
	// example:
	//
	// RD-Account
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The last time when the access configuration was provisioned.
	//
	// example:
	//
	// 2021-07-26T08:54:18Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) String() string {
	return dara.Prettify(s)
}

func (s ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) GoString() string {
	return s.String()
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) GetAccessConfigurationName() *string {
	return s.AccessConfigurationName
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) GetOriginTargetId() *string {
	return s.OriginTargetId
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) GetRAMPolicyNames() []*string {
	return s.RAMPolicyNames
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) GetRAMRoleName() *string {
	return s.RAMRoleName
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) GetSAMLProviderName() *string {
	return s.SAMLProviderName
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) GetStatus() *string {
	return s.Status
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) GetTargetId() *string {
	return s.TargetId
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) GetTargetName() *string {
	return s.TargetName
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) GetTargetPath() *string {
	return s.TargetPath
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) GetTargetPathName() *string {
	return s.TargetPathName
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) GetTargetType() *string {
	return s.TargetType
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetAccessConfigurationId(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.AccessConfigurationId = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetAccessConfigurationName(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.AccessConfigurationName = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetCreateTime(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.CreateTime = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetOriginTargetId(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.OriginTargetId = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetRAMPolicyNames(v []*string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.RAMPolicyNames = v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetRAMRoleName(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.RAMRoleName = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetSAMLProviderName(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.SAMLProviderName = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetStatus(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.Status = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetTargetId(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.TargetId = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetTargetName(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.TargetName = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetTargetPath(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.TargetPath = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetTargetPathName(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.TargetPathName = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetTargetType(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.TargetType = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) SetUpdateTime(v string) *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings {
	s.UpdateTime = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsResponseBodyAccessConfigurationProvisionings) Validate() error {
	return dara.Validate(s)
}
