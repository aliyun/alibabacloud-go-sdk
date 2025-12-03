// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationInstanceId(v string) *CreateSubscriptionRequest
	GetDestinationInstanceId() *string
	SetDestinationInstanceRegionId(v string) *CreateSubscriptionRequest
	GetDestinationInstanceRegionId() *string
	SetExtraContext(v string) *CreateSubscriptionRequest
	GetExtraContext() *string
	SetMapping(v string) *CreateSubscriptionRequest
	GetMapping() *string
	SetOwnerId(v int64) *CreateSubscriptionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateSubscriptionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSubscriptionRequest
	GetResourceOwnerId() *int64
	SetSlbServer(v string) *CreateSubscriptionRequest
	GetSlbServer() *string
	SetSourceInstanceId(v string) *CreateSubscriptionRequest
	GetSourceInstanceId() *string
	SetSourceInstanceRegionId(v string) *CreateSubscriptionRequest
	GetSourceInstanceRegionId() *string
	SetSubscriptionDescription(v string) *CreateSubscriptionRequest
	GetSubscriptionDescription() *string
	SetSubscriptionType(v string) *CreateSubscriptionRequest
	GetSubscriptionType() *string
	SetZoneId(v string) *CreateSubscriptionRequest
	GetZoneId() *string
}

type CreateSubscriptionRequest struct {
	// This parameter is required.
	DestinationInstanceId *string `json:"DestinationInstanceId,omitempty" xml:"DestinationInstanceId,omitempty"`
	// This parameter is required.
	DestinationInstanceRegionId *string `json:"DestinationInstanceRegionId,omitempty" xml:"DestinationInstanceRegionId,omitempty"`
	ExtraContext                *string `json:"ExtraContext,omitempty" xml:"ExtraContext,omitempty"`
	// This parameter is required.
	Mapping              *string `json:"Mapping,omitempty" xml:"Mapping,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SlbServer            *string `json:"SlbServer,omitempty" xml:"SlbServer,omitempty"`
	// This parameter is required.
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// This parameter is required.
	SourceInstanceRegionId  *string `json:"SourceInstanceRegionId,omitempty" xml:"SourceInstanceRegionId,omitempty"`
	SubscriptionDescription *string `json:"SubscriptionDescription,omitempty" xml:"SubscriptionDescription,omitempty"`
	// This parameter is required.
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	ZoneId           *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionRequest) GetDestinationInstanceId() *string {
	return s.DestinationInstanceId
}

func (s *CreateSubscriptionRequest) GetDestinationInstanceRegionId() *string {
	return s.DestinationInstanceRegionId
}

func (s *CreateSubscriptionRequest) GetExtraContext() *string {
	return s.ExtraContext
}

func (s *CreateSubscriptionRequest) GetMapping() *string {
	return s.Mapping
}

func (s *CreateSubscriptionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSubscriptionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSubscriptionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSubscriptionRequest) GetSlbServer() *string {
	return s.SlbServer
}

func (s *CreateSubscriptionRequest) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *CreateSubscriptionRequest) GetSourceInstanceRegionId() *string {
	return s.SourceInstanceRegionId
}

func (s *CreateSubscriptionRequest) GetSubscriptionDescription() *string {
	return s.SubscriptionDescription
}

func (s *CreateSubscriptionRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *CreateSubscriptionRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateSubscriptionRequest) SetDestinationInstanceId(v string) *CreateSubscriptionRequest {
	s.DestinationInstanceId = &v
	return s
}

func (s *CreateSubscriptionRequest) SetDestinationInstanceRegionId(v string) *CreateSubscriptionRequest {
	s.DestinationInstanceRegionId = &v
	return s
}

func (s *CreateSubscriptionRequest) SetExtraContext(v string) *CreateSubscriptionRequest {
	s.ExtraContext = &v
	return s
}

func (s *CreateSubscriptionRequest) SetMapping(v string) *CreateSubscriptionRequest {
	s.Mapping = &v
	return s
}

func (s *CreateSubscriptionRequest) SetOwnerId(v int64) *CreateSubscriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSubscriptionRequest) SetResourceOwnerAccount(v string) *CreateSubscriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSubscriptionRequest) SetResourceOwnerId(v int64) *CreateSubscriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSubscriptionRequest) SetSlbServer(v string) *CreateSubscriptionRequest {
	s.SlbServer = &v
	return s
}

func (s *CreateSubscriptionRequest) SetSourceInstanceId(v string) *CreateSubscriptionRequest {
	s.SourceInstanceId = &v
	return s
}

func (s *CreateSubscriptionRequest) SetSourceInstanceRegionId(v string) *CreateSubscriptionRequest {
	s.SourceInstanceRegionId = &v
	return s
}

func (s *CreateSubscriptionRequest) SetSubscriptionDescription(v string) *CreateSubscriptionRequest {
	s.SubscriptionDescription = &v
	return s
}

func (s *CreateSubscriptionRequest) SetSubscriptionType(v string) *CreateSubscriptionRequest {
	s.SubscriptionType = &v
	return s
}

func (s *CreateSubscriptionRequest) SetZoneId(v string) *CreateSubscriptionRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
