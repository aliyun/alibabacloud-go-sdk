// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterVpnAttachmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *ListTransitRouterVpnAttachmentsRequest
	GetCenId() *string
	SetMaxResults(v int32) *ListTransitRouterVpnAttachmentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterVpnAttachmentsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTransitRouterVpnAttachmentsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTransitRouterVpnAttachmentsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListTransitRouterVpnAttachmentsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListTransitRouterVpnAttachmentsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTransitRouterVpnAttachmentsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*ListTransitRouterVpnAttachmentsRequestTag) *ListTransitRouterVpnAttachmentsRequest
	GetTag() []*ListTransitRouterVpnAttachmentsRequestTag
	SetTransitRouterAttachmentId(v string) *ListTransitRouterVpnAttachmentsRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterId(v string) *ListTransitRouterVpnAttachmentsRequest
	GetTransitRouterId() *string
}

type ListTransitRouterVpnAttachmentsRequest struct {
	CenId                     *string                                      `json:"CenId,omitempty" xml:"CenId,omitempty"`
	MaxResults                *int32                                       `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                 *string                                      `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount              *string                                      `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                   *int64                                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                  *string                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount      *string                                      `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64                                       `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tag                       []*ListTransitRouterVpnAttachmentsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	TransitRouterAttachmentId *string                                      `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	TransitRouterId           *string                                      `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListTransitRouterVpnAttachmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVpnAttachmentsRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetCenId() *string {
	return s.CenId
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetTag() []*ListTransitRouterVpnAttachmentsRequestTag {
	return s.Tag
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListTransitRouterVpnAttachmentsRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetCenId(v string) *ListTransitRouterVpnAttachmentsRequest {
	s.CenId = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetMaxResults(v int32) *ListTransitRouterVpnAttachmentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetNextToken(v string) *ListTransitRouterVpnAttachmentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetOwnerAccount(v string) *ListTransitRouterVpnAttachmentsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetOwnerId(v int64) *ListTransitRouterVpnAttachmentsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetRegionId(v string) *ListTransitRouterVpnAttachmentsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetResourceOwnerAccount(v string) *ListTransitRouterVpnAttachmentsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetResourceOwnerId(v int64) *ListTransitRouterVpnAttachmentsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetTag(v []*ListTransitRouterVpnAttachmentsRequestTag) *ListTransitRouterVpnAttachmentsRequest {
	s.Tag = v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetTransitRouterAttachmentId(v string) *ListTransitRouterVpnAttachmentsRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) SetTransitRouterId(v string) *ListTransitRouterVpnAttachmentsRequest {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTransitRouterVpnAttachmentsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTransitRouterVpnAttachmentsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterVpnAttachmentsRequestTag) GoString() string {
	return s.String()
}

func (s *ListTransitRouterVpnAttachmentsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListTransitRouterVpnAttachmentsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListTransitRouterVpnAttachmentsRequestTag) SetKey(v string) *ListTransitRouterVpnAttachmentsRequestTag {
	s.Key = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequestTag) SetValue(v string) *ListTransitRouterVpnAttachmentsRequestTag {
	s.Value = &v
	return s
}

func (s *ListTransitRouterVpnAttachmentsRequestTag) Validate() error {
	return dara.Validate(s)
}
