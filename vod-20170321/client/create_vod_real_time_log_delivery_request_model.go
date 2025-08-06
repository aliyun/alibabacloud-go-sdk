// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVodRealTimeLogDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *CreateVodRealTimeLogDeliveryRequest
	GetDomainName() *string
	SetLogstore(v string) *CreateVodRealTimeLogDeliveryRequest
	GetLogstore() *string
	SetOwnerId(v int64) *CreateVodRealTimeLogDeliveryRequest
	GetOwnerId() *int64
	SetProject(v string) *CreateVodRealTimeLogDeliveryRequest
	GetProject() *string
	SetRegion(v string) *CreateVodRealTimeLogDeliveryRequest
	GetRegion() *string
}

type CreateVodRealTimeLogDeliveryRequest struct {
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

func (s CreateVodRealTimeLogDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVodRealTimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *CreateVodRealTimeLogDeliveryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateVodRealTimeLogDeliveryRequest) GetLogstore() *string {
	return s.Logstore
}

func (s *CreateVodRealTimeLogDeliveryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateVodRealTimeLogDeliveryRequest) GetProject() *string {
	return s.Project
}

func (s *CreateVodRealTimeLogDeliveryRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateVodRealTimeLogDeliveryRequest) SetDomainName(v string) *CreateVodRealTimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

func (s *CreateVodRealTimeLogDeliveryRequest) SetLogstore(v string) *CreateVodRealTimeLogDeliveryRequest {
	s.Logstore = &v
	return s
}

func (s *CreateVodRealTimeLogDeliveryRequest) SetOwnerId(v int64) *CreateVodRealTimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVodRealTimeLogDeliveryRequest) SetProject(v string) *CreateVodRealTimeLogDeliveryRequest {
	s.Project = &v
	return s
}

func (s *CreateVodRealTimeLogDeliveryRequest) SetRegion(v string) *CreateVodRealTimeLogDeliveryRequest {
	s.Region = &v
	return s
}

func (s *CreateVodRealTimeLogDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
