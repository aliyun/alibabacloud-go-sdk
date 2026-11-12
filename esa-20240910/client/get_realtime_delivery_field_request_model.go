// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealtimeDeliveryFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *GetRealtimeDeliveryFieldRequest
	GetBusinessType() *string
	SetSiteId(v string) *GetRealtimeDeliveryFieldRequest
	GetSiteId() *string
}

type GetRealtimeDeliveryFieldRequest struct {
	// The business type. Valid values:
	//
	// - **dcdn_log_access_l1*	- (default): Access logs.
	//
	// - **dcdn_log_er**: Edge Routine logs.
	//
	// - **dcdn_log_waf**: security protection logs.
	//
	// - **dcdn_log_ipa**: Layer 4 acceleration logs.
	//
	// This parameter is required.
	//
	// example:
	//
	// dcdn_log_er
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The site ID.
	//
	// example:
	//
	// 1234567890123
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetRealtimeDeliveryFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeDeliveryFieldRequest) GoString() string {
	return s.String()
}

func (s *GetRealtimeDeliveryFieldRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GetRealtimeDeliveryFieldRequest) GetSiteId() *string {
	return s.SiteId
}

func (s *GetRealtimeDeliveryFieldRequest) SetBusinessType(v string) *GetRealtimeDeliveryFieldRequest {
	s.BusinessType = &v
	return s
}

func (s *GetRealtimeDeliveryFieldRequest) SetSiteId(v string) *GetRealtimeDeliveryFieldRequest {
	s.SiteId = &v
	return s
}

func (s *GetRealtimeDeliveryFieldRequest) Validate() error {
	return dara.Validate(s)
}
