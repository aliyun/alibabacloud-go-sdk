// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodRealtimeLogDeliveryDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogstore(v string) *ListVodRealtimeLogDeliveryDomainsRequest
	GetLogstore() *string
	SetOwnerId(v int64) *ListVodRealtimeLogDeliveryDomainsRequest
	GetOwnerId() *int64
	SetProject(v string) *ListVodRealtimeLogDeliveryDomainsRequest
	GetProject() *string
	SetRegion(v string) *ListVodRealtimeLogDeliveryDomainsRequest
	GetRegion() *string
}

type ListVodRealtimeLogDeliveryDomainsRequest struct {
	// This parameter is required.
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// This parameter is required.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ListVodRealtimeLogDeliveryDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryDomainsRequest) GetLogstore() *string {
	return s.Logstore
}

func (s *ListVodRealtimeLogDeliveryDomainsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListVodRealtimeLogDeliveryDomainsRequest) GetProject() *string {
	return s.Project
}

func (s *ListVodRealtimeLogDeliveryDomainsRequest) GetRegion() *string {
	return s.Region
}

func (s *ListVodRealtimeLogDeliveryDomainsRequest) SetLogstore(v string) *ListVodRealtimeLogDeliveryDomainsRequest {
	s.Logstore = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryDomainsRequest) SetOwnerId(v int64) *ListVodRealtimeLogDeliveryDomainsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryDomainsRequest) SetProject(v string) *ListVodRealtimeLogDeliveryDomainsRequest {
	s.Project = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryDomainsRequest) SetRegion(v string) *ListVodRealtimeLogDeliveryDomainsRequest {
	s.Region = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryDomainsRequest) Validate() error {
	return dara.Validate(s)
}
