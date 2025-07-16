// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRealtimeLogDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *ModifyRealtimeLogDeliveryRequest
	GetDomain() *string
	SetLogstore(v string) *ModifyRealtimeLogDeliveryRequest
	GetLogstore() *string
	SetProject(v string) *ModifyRealtimeLogDeliveryRequest
	GetProject() *string
	SetRegion(v string) *ModifyRealtimeLogDeliveryRequest
	GetRegion() *string
}

type ModifyRealtimeLogDeliveryRequest struct {
	// The accelerated domain name for which you want to modify the configurations of real-time log delivery. Only one domain name is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The name of the Logstore where log entries are stored.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestLog
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	// The name of the Log Service project that is used for real-time log delivery.
	//
	// This parameter is required.
	//
	// example:
	//
	// testProject
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The ID of the region where the Log Service project is deployed. For more information, see [Regions that support real-time log delivery](https://help.aliyun.com/document_detail/144883.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ch-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ModifyRealtimeLogDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *ModifyRealtimeLogDeliveryRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyRealtimeLogDeliveryRequest) GetLogstore() *string {
	return s.Logstore
}

func (s *ModifyRealtimeLogDeliveryRequest) GetProject() *string {
	return s.Project
}

func (s *ModifyRealtimeLogDeliveryRequest) GetRegion() *string {
	return s.Region
}

func (s *ModifyRealtimeLogDeliveryRequest) SetDomain(v string) *ModifyRealtimeLogDeliveryRequest {
	s.Domain = &v
	return s
}

func (s *ModifyRealtimeLogDeliveryRequest) SetLogstore(v string) *ModifyRealtimeLogDeliveryRequest {
	s.Logstore = &v
	return s
}

func (s *ModifyRealtimeLogDeliveryRequest) SetProject(v string) *ModifyRealtimeLogDeliveryRequest {
	s.Project = &v
	return s
}

func (s *ModifyRealtimeLogDeliveryRequest) SetRegion(v string) *ModifyRealtimeLogDeliveryRequest {
	s.Region = &v
	return s
}

func (s *ModifyRealtimeLogDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
