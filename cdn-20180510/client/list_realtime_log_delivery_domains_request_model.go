// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRealtimeLogDeliveryDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogstore(v string) *ListRealtimeLogDeliveryDomainsRequest
	GetLogstore() *string
	SetProject(v string) *ListRealtimeLogDeliveryDomainsRequest
	GetProject() *string
	SetRegion(v string) *ListRealtimeLogDeliveryDomainsRequest
	GetRegion() *string
}

type ListRealtimeLogDeliveryDomainsRequest struct {
	// The name of the Logstore that collects log data from Alibaba Cloud CDN in real time. You can specify multiple Logstore names and separate them with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// LogstoreName
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	// The name of the Log Service project that is used for real-time log delivery. You can specify multiple project names and separate them with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ProjectName
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The ID of the region where the Log Service project is deployed. You can specify multiple region IDs and separate them with commas (,).
	//
	// For more information about regions, see [Regions that support real-time log delivery](https://help.aliyun.com/document_detail/144883.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ch-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ListRealtimeLogDeliveryDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeLogDeliveryDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryDomainsRequest) GetLogstore() *string {
	return s.Logstore
}

func (s *ListRealtimeLogDeliveryDomainsRequest) GetProject() *string {
	return s.Project
}

func (s *ListRealtimeLogDeliveryDomainsRequest) GetRegion() *string {
	return s.Region
}

func (s *ListRealtimeLogDeliveryDomainsRequest) SetLogstore(v string) *ListRealtimeLogDeliveryDomainsRequest {
	s.Logstore = &v
	return s
}

func (s *ListRealtimeLogDeliveryDomainsRequest) SetProject(v string) *ListRealtimeLogDeliveryDomainsRequest {
	s.Project = &v
	return s
}

func (s *ListRealtimeLogDeliveryDomainsRequest) SetRegion(v string) *ListRealtimeLogDeliveryDomainsRequest {
	s.Region = &v
	return s
}

func (s *ListRealtimeLogDeliveryDomainsRequest) Validate() error {
	return dara.Validate(s)
}
