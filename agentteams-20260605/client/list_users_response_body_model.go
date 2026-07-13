// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListUsersResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListUsersResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListUsersResponseBodyItems) *ListUsersResponseBody
	GetItems() []*ListUsersResponseBodyItems
	SetMaxResults(v int32) *ListUsersResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListUsersResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListUsersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListUsersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListUsersResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListUsersResponseBody
	GetTotalCount() *int64
}

type ListUsersResponseBody struct {
	Code           *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Items          []*ListUsersResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	MaxResults     *int32                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken      *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int64                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUsersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListUsersResponseBody) GetItems() []*ListUsersResponseBodyItems {
	return s.Items
}

func (s *ListUsersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUsersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUsersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListUsersResponseBody) SetCode(v string) *ListUsersResponseBody {
	s.Code = &v
	return s
}

func (s *ListUsersResponseBody) SetHttpStatusCode(v int32) *ListUsersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListUsersResponseBody) SetItems(v []*ListUsersResponseBodyItems) *ListUsersResponseBody {
	s.Items = v
	return s
}

func (s *ListUsersResponseBody) SetMaxResults(v int32) *ListUsersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUsersResponseBody) SetMessage(v string) *ListUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ListUsersResponseBody) SetNextToken(v string) *ListUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetSuccess(v bool) *ListUsersResponseBody {
	s.Success = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalCount(v int64) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBody) Validate() error {
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

type ListUsersResponseBodyItems struct {
	AuthMethod  *string `json:"AuthMethod,omitempty" xml:"AuthMethod,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListUsersResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyItems) GetAuthMethod() *string {
	return s.AuthMethod
}

func (s *ListUsersResponseBodyItems) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUsersResponseBodyItems) GetEmail() *string {
	return s.Email
}

func (s *ListUsersResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListUsersResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListUsersResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListUsersResponseBodyItems) SetAuthMethod(v string) *ListUsersResponseBodyItems {
	s.AuthMethod = &v
	return s
}

func (s *ListUsersResponseBodyItems) SetDisplayName(v string) *ListUsersResponseBodyItems {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyItems) SetEmail(v string) *ListUsersResponseBodyItems {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyItems) SetInstanceId(v string) *ListUsersResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *ListUsersResponseBodyItems) SetName(v string) *ListUsersResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListUsersResponseBodyItems) SetStatus(v string) *ListUsersResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListUsersResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
