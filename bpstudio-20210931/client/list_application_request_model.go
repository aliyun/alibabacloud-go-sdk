// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListApplicationRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListApplicationRequest
	GetMaxResults() *int32
	SetNextToken(v int32) *ListApplicationRequest
	GetNextToken() *int32
	SetOrderType(v int64) *ListApplicationRequest
	GetOrderType() *int64
	SetResourceGroupId(v string) *ListApplicationRequest
	GetResourceGroupId() *string
	SetResourceId(v string) *ListApplicationRequest
	GetResourceId() *string
	SetShowHide(v bool) *ListApplicationRequest
	GetShowHide() *bool
	SetStatus(v string) *ListApplicationRequest
	GetStatus() *string
	SetTemplateId(v string) *ListApplicationRequest
	GetTemplateId() *string
}

type ListApplicationRequest struct {
	// Keywords in the app name
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The pagination size of the resulting value cannot be less than the minimum value of 1 and cannot be greater than the maximum value of 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination page number of the resulting value cannot be less than the minimum value of 1 and cannot be greater than the maximum value of 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 1 update time,<br>2 creation time
	//
	// example:
	//
	// 1
	OrderType *int64 `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Resource Id
	//
	// example:
	//
	// vsw-xxxxxxxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ShowHide   *bool   `json:"ShowHide,omitempty" xml:"ShowHide,omitempty"`
	// The status of the applications to be returned.
	//
	// example:
	//
	// The following values are "success" and "release".
	//
	// If the input value is "success", the returned application list includes all applications in the Deployed_Success state of successful deployment.
	//
	// If the input value is release, the returned application list includes all applications in the release success (Destroyed_Success) and release failure (Destroyed_Failure) status.
	//
	// If this parameter is left blank, the returned app list includes apps in all states.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Template Id
	//
	// example:
	//
	// 0KSXXX6SJU03TXXX
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListApplicationRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationRequest) GetNextToken() *int32 {
	return s.NextToken
}

func (s *ListApplicationRequest) GetOrderType() *int64 {
	return s.OrderType
}

func (s *ListApplicationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListApplicationRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListApplicationRequest) GetShowHide() *bool {
	return s.ShowHide
}

func (s *ListApplicationRequest) GetStatus() *string {
	return s.Status
}

func (s *ListApplicationRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListApplicationRequest) SetKeyword(v string) *ListApplicationRequest {
	s.Keyword = &v
	return s
}

func (s *ListApplicationRequest) SetMaxResults(v int32) *ListApplicationRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationRequest) SetNextToken(v int32) *ListApplicationRequest {
	s.NextToken = &v
	return s
}

func (s *ListApplicationRequest) SetOrderType(v int64) *ListApplicationRequest {
	s.OrderType = &v
	return s
}

func (s *ListApplicationRequest) SetResourceGroupId(v string) *ListApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListApplicationRequest) SetResourceId(v string) *ListApplicationRequest {
	s.ResourceId = &v
	return s
}

func (s *ListApplicationRequest) SetShowHide(v bool) *ListApplicationRequest {
	s.ShowHide = &v
	return s
}

func (s *ListApplicationRequest) SetStatus(v string) *ListApplicationRequest {
	s.Status = &v
	return s
}

func (s *ListApplicationRequest) SetTemplateId(v string) *ListApplicationRequest {
	s.TemplateId = &v
	return s
}

func (s *ListApplicationRequest) Validate() error {
	return dara.Validate(s)
}
