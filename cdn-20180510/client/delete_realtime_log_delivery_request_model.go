// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRealtimeLogDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DeleteRealtimeLogDeliveryRequest
	GetDomain() *string
	SetLogstore(v string) *DeleteRealtimeLogDeliveryRequest
	GetLogstore() *string
	SetProject(v string) *DeleteRealtimeLogDeliveryRequest
	GetProject() *string
	SetRegion(v string) *DeleteRealtimeLogDeliveryRequest
	GetRegion() *string
}

type DeleteRealtimeLogDeliveryRequest struct {
	// The acceleration domain name for which you want to disable real-time log delivery. You can specify multiple domain names and separate them with commas (,).
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

func (s DeleteRealtimeLogDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *DeleteRealtimeLogDeliveryRequest) GetDomain() *string {
	return s.Domain
}

func (s *DeleteRealtimeLogDeliveryRequest) GetLogstore() *string {
	return s.Logstore
}

func (s *DeleteRealtimeLogDeliveryRequest) GetProject() *string {
	return s.Project
}

func (s *DeleteRealtimeLogDeliveryRequest) GetRegion() *string {
	return s.Region
}

func (s *DeleteRealtimeLogDeliveryRequest) SetDomain(v string) *DeleteRealtimeLogDeliveryRequest {
	s.Domain = &v
	return s
}

func (s *DeleteRealtimeLogDeliveryRequest) SetLogstore(v string) *DeleteRealtimeLogDeliveryRequest {
	s.Logstore = &v
	return s
}

func (s *DeleteRealtimeLogDeliveryRequest) SetProject(v string) *DeleteRealtimeLogDeliveryRequest {
	s.Project = &v
	return s
}

func (s *DeleteRealtimeLogDeliveryRequest) SetRegion(v string) *DeleteRealtimeLogDeliveryRequest {
	s.Region = &v
	return s
}

func (s *DeleteRealtimeLogDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
