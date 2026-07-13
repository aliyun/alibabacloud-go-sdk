// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListModelsResponseBody
	GetCode() *string
	SetItems(v []*ListModelsResponseBodyItems) *ListModelsResponseBody
	GetItems() []*ListModelsResponseBodyItems
	SetMaxResults(v int32) *ListModelsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListModelsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListModelsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListModelsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListModelsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListModelsResponseBody
	GetTotalCount() *int32
}

type ListModelsResponseBody struct {
	Code       *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Items      []*ListModelsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	MaxResults *int32                         `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message    *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken  *string                        `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int32                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListModelsResponseBody) GetItems() []*ListModelsResponseBodyItems {
	return s.Items
}

func (s *ListModelsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListModelsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListModelsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListModelsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListModelsResponseBody) SetCode(v string) *ListModelsResponseBody {
	s.Code = &v
	return s
}

func (s *ListModelsResponseBody) SetItems(v []*ListModelsResponseBodyItems) *ListModelsResponseBody {
	s.Items = v
	return s
}

func (s *ListModelsResponseBody) SetMaxResults(v int32) *ListModelsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListModelsResponseBody) SetMessage(v string) *ListModelsResponseBody {
	s.Message = &v
	return s
}

func (s *ListModelsResponseBody) SetNextToken(v string) *ListModelsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListModelsResponseBody) SetRequestId(v string) *ListModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelsResponseBody) SetSuccess(v bool) *ListModelsResponseBody {
	s.Success = &v
	return s
}

func (s *ListModelsResponseBody) SetTotalCount(v int32) *ListModelsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelsResponseBody) Validate() error {
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

type ListModelsResponseBodyItems struct {
	Description  *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Id           *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name         *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Protocols    []*string `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	Provider     *string   `json:"Provider,omitempty" xml:"Provider,omitempty"`
	ProviderId   *string   `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	ProviderName *string   `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	WorkerNum    *int64    `json:"WorkerNum,omitempty" xml:"WorkerNum,omitempty"`
}

func (s ListModelsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListModelsResponseBodyItems) GetId() *string {
	return s.Id
}

func (s *ListModelsResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListModelsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListModelsResponseBodyItems) GetProtocols() []*string {
	return s.Protocols
}

func (s *ListModelsResponseBodyItems) GetProvider() *string {
	return s.Provider
}

func (s *ListModelsResponseBodyItems) GetProviderId() *string {
	return s.ProviderId
}

func (s *ListModelsResponseBodyItems) GetProviderName() *string {
	return s.ProviderName
}

func (s *ListModelsResponseBodyItems) GetWorkerNum() *int64 {
	return s.WorkerNum
}

func (s *ListModelsResponseBodyItems) SetDescription(v string) *ListModelsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListModelsResponseBodyItems) SetId(v string) *ListModelsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *ListModelsResponseBodyItems) SetInstanceId(v string) *ListModelsResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *ListModelsResponseBodyItems) SetName(v string) *ListModelsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListModelsResponseBodyItems) SetProtocols(v []*string) *ListModelsResponseBodyItems {
	s.Protocols = v
	return s
}

func (s *ListModelsResponseBodyItems) SetProvider(v string) *ListModelsResponseBodyItems {
	s.Provider = &v
	return s
}

func (s *ListModelsResponseBodyItems) SetProviderId(v string) *ListModelsResponseBodyItems {
	s.ProviderId = &v
	return s
}

func (s *ListModelsResponseBodyItems) SetProviderName(v string) *ListModelsResponseBodyItems {
	s.ProviderName = &v
	return s
}

func (s *ListModelsResponseBodyItems) SetWorkerNum(v int64) *ListModelsResponseBodyItems {
	s.WorkerNum = &v
	return s
}

func (s *ListModelsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
