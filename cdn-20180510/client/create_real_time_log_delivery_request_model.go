// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRealTimeLogDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *CreateRealTimeLogDeliveryRequest
	GetDomain() *string
	SetLogstore(v string) *CreateRealTimeLogDeliveryRequest
	GetLogstore() *string
	SetProject(v string) *CreateRealTimeLogDeliveryRequest
	GetProject() *string
	SetRegion(v string) *CreateRealTimeLogDeliveryRequest
	GetRegion() *string
}

type CreateRealTimeLogDeliveryRequest struct {
	// The accelerated domain name for which you want to configure real-time log delivery.
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
	// LogstoreName
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	// The name of the Log Service project that is used for real-time log delivery.
	//
	// This parameter is required.
	//
	// example:
	//
	// ProjectName
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The ID of the region where the Log Service project is deployed. For more information, see [Regions that support real-time log delivery](https://help.aliyun.com/document_detail/144883.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s CreateRealTimeLogDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRealTimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *CreateRealTimeLogDeliveryRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateRealTimeLogDeliveryRequest) GetLogstore() *string {
	return s.Logstore
}

func (s *CreateRealTimeLogDeliveryRequest) GetProject() *string {
	return s.Project
}

func (s *CreateRealTimeLogDeliveryRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateRealTimeLogDeliveryRequest) SetDomain(v string) *CreateRealTimeLogDeliveryRequest {
	s.Domain = &v
	return s
}

func (s *CreateRealTimeLogDeliveryRequest) SetLogstore(v string) *CreateRealTimeLogDeliveryRequest {
	s.Logstore = &v
	return s
}

func (s *CreateRealTimeLogDeliveryRequest) SetProject(v string) *CreateRealTimeLogDeliveryRequest {
	s.Project = &v
	return s
}

func (s *CreateRealTimeLogDeliveryRequest) SetRegion(v string) *CreateRealTimeLogDeliveryRequest {
	s.Region = &v
	return s
}

func (s *CreateRealTimeLogDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
