// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodRealtimeLogDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DeleteVodRealtimeLogDeliveryRequest
	GetDomainName() *string
	SetLogstore(v string) *DeleteVodRealtimeLogDeliveryRequest
	GetLogstore() *string
	SetOwnerId(v int64) *DeleteVodRealtimeLogDeliveryRequest
	GetOwnerId() *int64
	SetProject(v string) *DeleteVodRealtimeLogDeliveryRequest
	GetProject() *string
	SetRegion(v string) *DeleteVodRealtimeLogDeliveryRequest
	GetRegion() *string
}

type DeleteVodRealtimeLogDeliveryRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// This parameter is required.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DeleteVodRealtimeLogDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *DeleteVodRealtimeLogDeliveryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteVodRealtimeLogDeliveryRequest) GetLogstore() *string {
	return s.Logstore
}

func (s *DeleteVodRealtimeLogDeliveryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVodRealtimeLogDeliveryRequest) GetProject() *string {
	return s.Project
}

func (s *DeleteVodRealtimeLogDeliveryRequest) GetRegion() *string {
	return s.Region
}

func (s *DeleteVodRealtimeLogDeliveryRequest) SetDomainName(v string) *DeleteVodRealtimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteVodRealtimeLogDeliveryRequest) SetLogstore(v string) *DeleteVodRealtimeLogDeliveryRequest {
	s.Logstore = &v
	return s
}

func (s *DeleteVodRealtimeLogDeliveryRequest) SetOwnerId(v int64) *DeleteVodRealtimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVodRealtimeLogDeliveryRequest) SetProject(v string) *DeleteVodRealtimeLogDeliveryRequest {
	s.Project = &v
	return s
}

func (s *DeleteVodRealtimeLogDeliveryRequest) SetRegion(v string) *DeleteVodRealtimeLogDeliveryRequest {
	s.Region = &v
	return s
}

func (s *DeleteVodRealtimeLogDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
