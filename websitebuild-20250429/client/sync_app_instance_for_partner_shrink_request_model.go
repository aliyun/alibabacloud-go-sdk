// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncAppInstanceForPartnerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceShrink(v string) *SyncAppInstanceForPartnerShrinkRequest
	GetAppInstanceShrink() *string
	SetEventType(v string) *SyncAppInstanceForPartnerShrinkRequest
	GetEventType() *string
	SetOperator(v string) *SyncAppInstanceForPartnerShrinkRequest
	GetOperator() *string
	SetSourceBizId(v string) *SyncAppInstanceForPartnerShrinkRequest
	GetSourceBizId() *string
	SetSourceType(v string) *SyncAppInstanceForPartnerShrinkRequest
	GetSourceType() *string
}

type SyncAppInstanceForPartnerShrinkRequest struct {
	// Application instance object data
	AppInstanceShrink *string `json:"AppInstance,omitempty" xml:"AppInstance,omitempty"`
	// Type of system event. CREATE, UPDATE, COMPLETE
	//
	// example:
	//
	// CREATE
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// Operator ID
	//
	// example:
	//
	// system
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// Source business ID.
	//
	// example:
	//
	// 31104757
	SourceBizId *string `json:"SourceBizId,omitempty" xml:"SourceBizId,omitempty"`
	// Source: MARKET_CLOUD_DREAM
	//
	// example:
	//
	// MARKET_CLOUD_DREAM
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s SyncAppInstanceForPartnerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncAppInstanceForPartnerShrinkRequest) GoString() string {
	return s.String()
}

func (s *SyncAppInstanceForPartnerShrinkRequest) GetAppInstanceShrink() *string {
	return s.AppInstanceShrink
}

func (s *SyncAppInstanceForPartnerShrinkRequest) GetEventType() *string {
	return s.EventType
}

func (s *SyncAppInstanceForPartnerShrinkRequest) GetOperator() *string {
	return s.Operator
}

func (s *SyncAppInstanceForPartnerShrinkRequest) GetSourceBizId() *string {
	return s.SourceBizId
}

func (s *SyncAppInstanceForPartnerShrinkRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *SyncAppInstanceForPartnerShrinkRequest) SetAppInstanceShrink(v string) *SyncAppInstanceForPartnerShrinkRequest {
	s.AppInstanceShrink = &v
	return s
}

func (s *SyncAppInstanceForPartnerShrinkRequest) SetEventType(v string) *SyncAppInstanceForPartnerShrinkRequest {
	s.EventType = &v
	return s
}

func (s *SyncAppInstanceForPartnerShrinkRequest) SetOperator(v string) *SyncAppInstanceForPartnerShrinkRequest {
	s.Operator = &v
	return s
}

func (s *SyncAppInstanceForPartnerShrinkRequest) SetSourceBizId(v string) *SyncAppInstanceForPartnerShrinkRequest {
	s.SourceBizId = &v
	return s
}

func (s *SyncAppInstanceForPartnerShrinkRequest) SetSourceType(v string) *SyncAppInstanceForPartnerShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *SyncAppInstanceForPartnerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
