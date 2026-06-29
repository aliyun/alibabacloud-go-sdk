// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentStoragesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentStorages(v []*ListAgentStoragesResponseBodyAgentStorages) *ListAgentStoragesResponseBody
	GetAgentStorages() []*ListAgentStoragesResponseBodyAgentStorages
	SetNextToken(v string) *ListAgentStoragesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAgentStoragesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListAgentStoragesResponseBody
	GetTotalCount() *int64
}

type ListAgentStoragesResponseBody struct {
	// The list of agent storage information.
	AgentStorages []*ListAgentStoragesResponseBodyAgentStorages `json:"AgentStorages,omitempty" xml:"AgentStorages,omitempty" type:"Repeated"`
	// The token used to retrieve the next page of results when the total number of tag resources exceeds the value of MaxResults. This parameter has a value only when not all tag resources are returned.
	//
	// example:
	//
	// CAESCG15aC1xxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 39871ED2-62C0-578F-A32E-B88072D5582F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of agent storages returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAgentStoragesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentStoragesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentStoragesResponseBody) GetAgentStorages() []*ListAgentStoragesResponseBodyAgentStorages {
	return s.AgentStorages
}

func (s *ListAgentStoragesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentStoragesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentStoragesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAgentStoragesResponseBody) SetAgentStorages(v []*ListAgentStoragesResponseBodyAgentStorages) *ListAgentStoragesResponseBody {
	s.AgentStorages = v
	return s
}

func (s *ListAgentStoragesResponseBody) SetNextToken(v string) *ListAgentStoragesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAgentStoragesResponseBody) SetRequestId(v string) *ListAgentStoragesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentStoragesResponseBody) SetTotalCount(v int64) *ListAgentStoragesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAgentStoragesResponseBody) Validate() error {
	if s.AgentStorages != nil {
		for _, item := range s.AgentStorages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgentStoragesResponseBodyAgentStorages struct {
	// agent storage description
	//
	// example:
	//
	// description for agent storage
	AgentStorageDescription *string `json:"AgentStorageDescription,omitempty" xml:"AgentStorageDescription,omitempty"`
	// The agent storage name, which is a unique key.
	//
	// example:
	//
	// agent-test
	AgentStorageName *string `json:"AgentStorageName,omitempty" xml:"AgentStorageName,omitempty"`
	// The specifications of the agent storage.
	//
	// example:
	//
	// SSD
	AgentStorageSpecification *string `json:"AgentStorageSpecification,omitempty" xml:"AgentStorageSpecification,omitempty"`
	// The status of the agent storage. Valid values:
	//
	// - normal: Normal.
	//
	// - forbidden: Disabled.
	//
	// - deleting: Being released.
	//
	// example:
	//
	// normal
	AgentStorageStatus *string `json:"AgentStorageStatus,omitempty" xml:"AgentStorageStatus,omitempty"`
	// The alias of the agent storage.
	//
	// example:
	//
	// btnots
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The creation time of the agent storage.
	//
	// example:
	//
	// 2025-04-16T06:02:59Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The region ID of the agent storage.
	//
	// example:
	//
	// cn-shanghai-cloudspe
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxh4em5jnbcd
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 13542356466
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListAgentStoragesResponseBodyAgentStorages) String() string {
	return dara.Prettify(s)
}

func (s ListAgentStoragesResponseBodyAgentStorages) GoString() string {
	return s.String()
}

func (s *ListAgentStoragesResponseBodyAgentStorages) GetAgentStorageDescription() *string {
	return s.AgentStorageDescription
}

func (s *ListAgentStoragesResponseBodyAgentStorages) GetAgentStorageName() *string {
	return s.AgentStorageName
}

func (s *ListAgentStoragesResponseBodyAgentStorages) GetAgentStorageSpecification() *string {
	return s.AgentStorageSpecification
}

func (s *ListAgentStoragesResponseBodyAgentStorages) GetAgentStorageStatus() *string {
	return s.AgentStorageStatus
}

func (s *ListAgentStoragesResponseBodyAgentStorages) GetAliasName() *string {
	return s.AliasName
}

func (s *ListAgentStoragesResponseBodyAgentStorages) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAgentStoragesResponseBodyAgentStorages) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAgentStoragesResponseBodyAgentStorages) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListAgentStoragesResponseBodyAgentStorages) GetUserId() *string {
	return s.UserId
}

func (s *ListAgentStoragesResponseBodyAgentStorages) SetAgentStorageDescription(v string) *ListAgentStoragesResponseBodyAgentStorages {
	s.AgentStorageDescription = &v
	return s
}

func (s *ListAgentStoragesResponseBodyAgentStorages) SetAgentStorageName(v string) *ListAgentStoragesResponseBodyAgentStorages {
	s.AgentStorageName = &v
	return s
}

func (s *ListAgentStoragesResponseBodyAgentStorages) SetAgentStorageSpecification(v string) *ListAgentStoragesResponseBodyAgentStorages {
	s.AgentStorageSpecification = &v
	return s
}

func (s *ListAgentStoragesResponseBodyAgentStorages) SetAgentStorageStatus(v string) *ListAgentStoragesResponseBodyAgentStorages {
	s.AgentStorageStatus = &v
	return s
}

func (s *ListAgentStoragesResponseBodyAgentStorages) SetAliasName(v string) *ListAgentStoragesResponseBodyAgentStorages {
	s.AliasName = &v
	return s
}

func (s *ListAgentStoragesResponseBodyAgentStorages) SetCreateTime(v string) *ListAgentStoragesResponseBodyAgentStorages {
	s.CreateTime = &v
	return s
}

func (s *ListAgentStoragesResponseBodyAgentStorages) SetRegionId(v string) *ListAgentStoragesResponseBodyAgentStorages {
	s.RegionId = &v
	return s
}

func (s *ListAgentStoragesResponseBodyAgentStorages) SetResourceGroupId(v string) *ListAgentStoragesResponseBodyAgentStorages {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAgentStoragesResponseBodyAgentStorages) SetUserId(v string) *ListAgentStoragesResponseBodyAgentStorages {
	s.UserId = &v
	return s
}

func (s *ListAgentStoragesResponseBodyAgentStorages) Validate() error {
	return dara.Validate(s)
}
