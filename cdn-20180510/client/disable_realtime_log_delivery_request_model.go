// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableRealtimeLogDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DisableRealtimeLogDeliveryRequest
	GetDomain() *string
	SetLogstore(v string) *DisableRealtimeLogDeliveryRequest
	GetLogstore() *string
	SetProject(v string) *DisableRealtimeLogDeliveryRequest
	GetProject() *string
	SetRegion(v string) *DisableRealtimeLogDeliveryRequest
	GetRegion() *string
}

type DisableRealtimeLogDeliveryRequest struct {
	// The accelerated domain name for which you want to disable real-time log delivery. You can specify multiple domain names and separate them with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The name of the Logstore where log entries are stored.
	//
	// example:
	//
	// LogstoreName
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	// The name of the Log Service project that is used for real-time log delivery.
	//
	// example:
	//
	// ProjectName
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The ID of the region where the Log Service project is deployed.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DisableRealtimeLogDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *DisableRealtimeLogDeliveryRequest) GetDomain() *string {
	return s.Domain
}

func (s *DisableRealtimeLogDeliveryRequest) GetLogstore() *string {
	return s.Logstore
}

func (s *DisableRealtimeLogDeliveryRequest) GetProject() *string {
	return s.Project
}

func (s *DisableRealtimeLogDeliveryRequest) GetRegion() *string {
	return s.Region
}

func (s *DisableRealtimeLogDeliveryRequest) SetDomain(v string) *DisableRealtimeLogDeliveryRequest {
	s.Domain = &v
	return s
}

func (s *DisableRealtimeLogDeliveryRequest) SetLogstore(v string) *DisableRealtimeLogDeliveryRequest {
	s.Logstore = &v
	return s
}

func (s *DisableRealtimeLogDeliveryRequest) SetProject(v string) *DisableRealtimeLogDeliveryRequest {
	s.Project = &v
	return s
}

func (s *DisableRealtimeLogDeliveryRequest) SetRegion(v string) *DisableRealtimeLogDeliveryRequest {
	s.Region = &v
	return s
}

func (s *DisableRealtimeLogDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
