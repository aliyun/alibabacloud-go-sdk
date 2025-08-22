// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnSLSRealtimeLogDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *DescribeDcdnSLSRealtimeLogDeliveryRequest
	GetProjectName() *string
}

type DescribeDcdnSLSRealtimeLogDeliveryRequest struct {
	// The name of a real-time log delivery project.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DescribeDcdnSLSRealtimeLogDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSLSRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryRequest) SetProjectName(v string) *DescribeDcdnSLSRealtimeLogDeliveryRequest {
	s.ProjectName = &v
	return s
}

func (s *DescribeDcdnSLSRealtimeLogDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
