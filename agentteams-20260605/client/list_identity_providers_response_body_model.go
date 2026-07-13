// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityProvidersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListIdentityProvidersResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListIdentityProvidersResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListIdentityProvidersResponseBodyItems) *ListIdentityProvidersResponseBody
	GetItems() []*ListIdentityProvidersResponseBodyItems
	SetMaxResults(v int32) *ListIdentityProvidersResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListIdentityProvidersResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListIdentityProvidersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListIdentityProvidersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListIdentityProvidersResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListIdentityProvidersResponseBody
	GetTotalCount() *int64
}

type ListIdentityProvidersResponseBody struct {
	Code           *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Items          []*ListIdentityProvidersResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	MaxResults     *int32                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message        *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken      *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int64                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIdentityProvidersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ListIdentityProvidersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListIdentityProvidersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListIdentityProvidersResponseBody) GetItems() []*ListIdentityProvidersResponseBodyItems {
	return s.Items
}

func (s *ListIdentityProvidersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIdentityProvidersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListIdentityProvidersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIdentityProvidersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIdentityProvidersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListIdentityProvidersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListIdentityProvidersResponseBody) SetCode(v string) *ListIdentityProvidersResponseBody {
	s.Code = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetHttpStatusCode(v int32) *ListIdentityProvidersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetItems(v []*ListIdentityProvidersResponseBodyItems) *ListIdentityProvidersResponseBody {
	s.Items = v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetMaxResults(v int32) *ListIdentityProvidersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetMessage(v string) *ListIdentityProvidersResponseBody {
	s.Message = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetNextToken(v string) *ListIdentityProvidersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetRequestId(v string) *ListIdentityProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetSuccess(v bool) *ListIdentityProvidersResponseBody {
	s.Success = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) SetTotalCount(v int64) *ListIdentityProvidersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIdentityProvidersResponseBody) Validate() error {
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

type ListIdentityProvidersResponseBodyItems struct {
	IdentityProviderType *string `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LoginEnabled         *bool   `json:"LoginEnabled,omitempty" xml:"LoginEnabled,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SyncEnabled          *bool   `json:"SyncEnabled,omitempty" xml:"SyncEnabled,omitempty"`
}

func (s ListIdentityProvidersResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProvidersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListIdentityProvidersResponseBodyItems) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *ListIdentityProvidersResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListIdentityProvidersResponseBodyItems) GetLoginEnabled() *bool {
	return s.LoginEnabled
}

func (s *ListIdentityProvidersResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListIdentityProvidersResponseBodyItems) GetSyncEnabled() *bool {
	return s.SyncEnabled
}

func (s *ListIdentityProvidersResponseBodyItems) SetIdentityProviderType(v string) *ListIdentityProvidersResponseBodyItems {
	s.IdentityProviderType = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItems) SetInstanceId(v string) *ListIdentityProvidersResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItems) SetLoginEnabled(v bool) *ListIdentityProvidersResponseBodyItems {
	s.LoginEnabled = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItems) SetStatus(v string) *ListIdentityProvidersResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItems) SetSyncEnabled(v bool) *ListIdentityProvidersResponseBodyItems {
	s.SyncEnabled = &v
	return s
}

func (s *ListIdentityProvidersResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
