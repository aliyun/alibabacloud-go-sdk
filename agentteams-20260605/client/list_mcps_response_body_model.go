// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMcpsResponseBody
	GetCode() *string
	SetItems(v []*ListMcpsResponseBodyItems) *ListMcpsResponseBody
	GetItems() []*ListMcpsResponseBodyItems
	SetMaxResults(v int32) *ListMcpsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListMcpsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListMcpsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMcpsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMcpsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListMcpsResponseBody
	GetTotalCount() *int32
}

type ListMcpsResponseBody struct {
	Code       *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Items      []*ListMcpsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	MaxResults *int32                       `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message    *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken  *string                      `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMcpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMcpsResponseBody) GetItems() []*ListMcpsResponseBodyItems {
	return s.Items
}

func (s *ListMcpsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMcpsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMcpsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMcpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcpsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMcpsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMcpsResponseBody) SetCode(v string) *ListMcpsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMcpsResponseBody) SetItems(v []*ListMcpsResponseBodyItems) *ListMcpsResponseBody {
	s.Items = v
	return s
}

func (s *ListMcpsResponseBody) SetMaxResults(v int32) *ListMcpsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMcpsResponseBody) SetMessage(v string) *ListMcpsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMcpsResponseBody) SetNextToken(v string) *ListMcpsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMcpsResponseBody) SetRequestId(v string) *ListMcpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcpsResponseBody) SetSuccess(v bool) *ListMcpsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMcpsResponseBody) SetTotalCount(v int32) *ListMcpsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMcpsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMcpsResponseBodyItems struct {
	Addresses       []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	CreateType      *string   `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	DeployStatus    *string   `json:"DeployStatus,omitempty" xml:"DeployStatus,omitempty"`
	Description     *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Id              *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId      *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	McpServerConfig *string   `json:"McpServerConfig,omitempty" xml:"McpServerConfig,omitempty"`
	Name            *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Protocol        *string   `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Url             *string   `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListMcpsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListMcpsResponseBodyItems) GetAddresses() []*string {
	return s.Addresses
}

func (s *ListMcpsResponseBodyItems) GetCreateType() *string {
	return s.CreateType
}

func (s *ListMcpsResponseBodyItems) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *ListMcpsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListMcpsResponseBodyItems) GetId() *string {
	return s.Id
}

func (s *ListMcpsResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListMcpsResponseBodyItems) GetMcpServerConfig() *string {
	return s.McpServerConfig
}

func (s *ListMcpsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListMcpsResponseBodyItems) GetProtocol() *string {
	return s.Protocol
}

func (s *ListMcpsResponseBodyItems) GetUrl() *string {
	return s.Url
}

func (s *ListMcpsResponseBodyItems) SetAddresses(v []*string) *ListMcpsResponseBodyItems {
	s.Addresses = v
	return s
}

func (s *ListMcpsResponseBodyItems) SetCreateType(v string) *ListMcpsResponseBodyItems {
	s.CreateType = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetDeployStatus(v string) *ListMcpsResponseBodyItems {
	s.DeployStatus = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetDescription(v string) *ListMcpsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetId(v string) *ListMcpsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetInstanceId(v string) *ListMcpsResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetMcpServerConfig(v string) *ListMcpsResponseBodyItems {
	s.McpServerConfig = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetName(v string) *ListMcpsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetProtocol(v string) *ListMcpsResponseBodyItems {
	s.Protocol = &v
	return s
}

func (s *ListMcpsResponseBodyItems) SetUrl(v string) *ListMcpsResponseBodyItems {
	s.Url = &v
	return s
}

func (s *ListMcpsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
