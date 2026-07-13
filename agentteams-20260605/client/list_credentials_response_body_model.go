// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCredentialsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListCredentialsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListCredentialsResponseBodyItems) *ListCredentialsResponseBody
	GetItems() []*ListCredentialsResponseBodyItems
	SetMaxResults(v int32) *ListCredentialsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListCredentialsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListCredentialsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListCredentialsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCredentialsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListCredentialsResponseBody
	GetTotalCount() *int64
}

type ListCredentialsResponseBody struct {
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Items          []*ListCredentialsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	MaxResults     *int32                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken      *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int64                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCredentialsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCredentialsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCredentialsResponseBody) GetItems() []*ListCredentialsResponseBodyItems {
	return s.Items
}

func (s *ListCredentialsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCredentialsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCredentialsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCredentialsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCredentialsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCredentialsResponseBody) SetCode(v string) *ListCredentialsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCredentialsResponseBody) SetHttpStatusCode(v int32) *ListCredentialsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCredentialsResponseBody) SetItems(v []*ListCredentialsResponseBodyItems) *ListCredentialsResponseBody {
	s.Items = v
	return s
}

func (s *ListCredentialsResponseBody) SetMaxResults(v int32) *ListCredentialsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCredentialsResponseBody) SetMessage(v string) *ListCredentialsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCredentialsResponseBody) SetNextToken(v string) *ListCredentialsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCredentialsResponseBody) SetRequestId(v string) *ListCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCredentialsResponseBody) SetSuccess(v bool) *ListCredentialsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCredentialsResponseBody) SetTotalCount(v int64) *ListCredentialsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCredentialsResponseBody) Validate() error {
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

type ListCredentialsResponseBodyItems struct {
	BoundWorkerCount *int32  `json:"BoundWorkerCount,omitempty" xml:"BoundWorkerCount,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	UpdateTime       *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListCredentialsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListCredentialsResponseBodyItems) GetBoundWorkerCount() *int32 {
	return s.BoundWorkerCount
}

func (s *ListCredentialsResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListCredentialsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListCredentialsResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCredentialsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListCredentialsResponseBodyItems) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListCredentialsResponseBodyItems) SetBoundWorkerCount(v int32) *ListCredentialsResponseBodyItems {
	s.BoundWorkerCount = &v
	return s
}

func (s *ListCredentialsResponseBodyItems) SetCreateTime(v string) *ListCredentialsResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *ListCredentialsResponseBodyItems) SetDescription(v string) *ListCredentialsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListCredentialsResponseBodyItems) SetInstanceId(v string) *ListCredentialsResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *ListCredentialsResponseBodyItems) SetName(v string) *ListCredentialsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListCredentialsResponseBodyItems) SetUpdateTime(v string) *ListCredentialsResponseBodyItems {
	s.UpdateTime = &v
	return s
}

func (s *ListCredentialsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
