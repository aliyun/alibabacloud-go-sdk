// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelProvidersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListModelProvidersResponseBody
	GetCode() *string
	SetItems(v []*ListModelProvidersResponseBodyItems) *ListModelProvidersResponseBody
	GetItems() []*ListModelProvidersResponseBodyItems
	SetMaxResults(v int32) *ListModelProvidersResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListModelProvidersResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListModelProvidersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListModelProvidersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListModelProvidersResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListModelProvidersResponseBody
	GetTotalCount() *int32
}

type ListModelProvidersResponseBody struct {
	Code       *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Items      []*ListModelProvidersResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	MaxResults *int32                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message    *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken  *string                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListModelProvidersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelProvidersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListModelProvidersResponseBody) GetItems() []*ListModelProvidersResponseBodyItems {
	return s.Items
}

func (s *ListModelProvidersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListModelProvidersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListModelProvidersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListModelProvidersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelProvidersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListModelProvidersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListModelProvidersResponseBody) SetCode(v string) *ListModelProvidersResponseBody {
	s.Code = &v
	return s
}

func (s *ListModelProvidersResponseBody) SetItems(v []*ListModelProvidersResponseBodyItems) *ListModelProvidersResponseBody {
	s.Items = v
	return s
}

func (s *ListModelProvidersResponseBody) SetMaxResults(v int32) *ListModelProvidersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListModelProvidersResponseBody) SetMessage(v string) *ListModelProvidersResponseBody {
	s.Message = &v
	return s
}

func (s *ListModelProvidersResponseBody) SetNextToken(v string) *ListModelProvidersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListModelProvidersResponseBody) SetRequestId(v string) *ListModelProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelProvidersResponseBody) SetSuccess(v bool) *ListModelProvidersResponseBody {
	s.Success = &v
	return s
}

func (s *ListModelProvidersResponseBody) SetTotalCount(v int32) *ListModelProvidersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelProvidersResponseBody) Validate() error {
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

type ListModelProvidersResponseBodyItems struct {
	Address      *string   `json:"Address,omitempty" xml:"Address,omitempty"`
	ApiKeys      []*string `json:"ApiKeys,omitempty" xml:"ApiKeys,omitempty" type:"Repeated"`
	CreateTime   *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeployStatus *string   `json:"DeployStatus,omitempty" xml:"DeployStatus,omitempty"`
	Description  *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Id           *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name         *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Protocols    []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	Provider     *string   `json:"Provider,omitempty" xml:"Provider,omitempty"`
}

func (s ListModelProvidersResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListModelProvidersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListModelProvidersResponseBodyItems) GetAddress() *string {
	return s.Address
}

func (s *ListModelProvidersResponseBodyItems) GetApiKeys() []*string {
	return s.ApiKeys
}

func (s *ListModelProvidersResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListModelProvidersResponseBodyItems) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *ListModelProvidersResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListModelProvidersResponseBodyItems) GetId() *string {
	return s.Id
}

func (s *ListModelProvidersResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListModelProvidersResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListModelProvidersResponseBodyItems) GetProtocols() []*string {
	return s.Protocols
}

func (s *ListModelProvidersResponseBodyItems) GetProvider() *string {
	return s.Provider
}

func (s *ListModelProvidersResponseBodyItems) SetAddress(v string) *ListModelProvidersResponseBodyItems {
	s.Address = &v
	return s
}

func (s *ListModelProvidersResponseBodyItems) SetApiKeys(v []*string) *ListModelProvidersResponseBodyItems {
	s.ApiKeys = v
	return s
}

func (s *ListModelProvidersResponseBodyItems) SetCreateTime(v string) *ListModelProvidersResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *ListModelProvidersResponseBodyItems) SetDeployStatus(v string) *ListModelProvidersResponseBodyItems {
	s.DeployStatus = &v
	return s
}

func (s *ListModelProvidersResponseBodyItems) SetDescription(v string) *ListModelProvidersResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListModelProvidersResponseBodyItems) SetId(v string) *ListModelProvidersResponseBodyItems {
	s.Id = &v
	return s
}

func (s *ListModelProvidersResponseBodyItems) SetInstanceId(v string) *ListModelProvidersResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *ListModelProvidersResponseBodyItems) SetName(v string) *ListModelProvidersResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListModelProvidersResponseBodyItems) SetProtocols(v []*string) *ListModelProvidersResponseBodyItems {
	s.Protocols = v
	return s
}

func (s *ListModelProvidersResponseBodyItems) SetProvider(v string) *ListModelProvidersResponseBodyItems {
	s.Provider = &v
	return s
}

func (s *ListModelProvidersResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
