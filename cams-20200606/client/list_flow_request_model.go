// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *ListFlowRequest
	GetCustSpaceId() *string
	SetFlowName(v string) *ListFlowRequest
	GetFlowName() *string
	SetOwnerId(v int64) *ListFlowRequest
	GetOwnerId() *int64
	SetPage(v *ListFlowRequestPage) *ListFlowRequest
	GetPage() *ListFlowRequestPage
	SetResourceOwnerAccount(v string) *ListFlowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListFlowRequest
	GetResourceOwnerId() *int64
}

type ListFlowRequest struct {
	// The space ID of the RAM user within the independent software vendor (ISV) account.
	//
	// example:
	//
	// 99948484
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The name of the Flow that you want to query. If FlowName is left empty, the information about all Flows is queried.
	//
	// example:
	//
	// flow_001
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The returned pages.
	Page                 *ListFlowRequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	ResourceOwnerAccount *string              `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64               `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFlowRequest) GoString() string {
	return s.String()
}

func (s *ListFlowRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListFlowRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *ListFlowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListFlowRequest) GetPage() *ListFlowRequestPage {
	return s.Page
}

func (s *ListFlowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListFlowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListFlowRequest) SetCustSpaceId(v string) *ListFlowRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListFlowRequest) SetFlowName(v string) *ListFlowRequest {
	s.FlowName = &v
	return s
}

func (s *ListFlowRequest) SetOwnerId(v int64) *ListFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *ListFlowRequest) SetPage(v *ListFlowRequestPage) *ListFlowRequest {
	s.Page = v
	return s
}

func (s *ListFlowRequest) SetResourceOwnerAccount(v string) *ListFlowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListFlowRequest) SetResourceOwnerId(v int64) *ListFlowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListFlowRequest) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFlowRequestPage struct {
	// The page number.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListFlowRequestPage) String() string {
	return dara.Prettify(s)
}

func (s ListFlowRequestPage) GoString() string {
	return s.String()
}

func (s *ListFlowRequestPage) GetIndex() *int32 {
	return s.Index
}

func (s *ListFlowRequestPage) GetSize() *int32 {
	return s.Size
}

func (s *ListFlowRequestPage) SetIndex(v int32) *ListFlowRequestPage {
	s.Index = &v
	return s
}

func (s *ListFlowRequestPage) SetSize(v int32) *ListFlowRequestPage {
	s.Size = &v
	return s
}

func (s *ListFlowRequestPage) Validate() error {
	return dara.Validate(s)
}
