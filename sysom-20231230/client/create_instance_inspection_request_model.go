// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceInspectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *CreateInstanceInspectionRequest
	GetXDebugId() *string
	SetInstance(v string) *CreateInstanceInspectionRequest
	GetInstance() *string
	SetItems(v []*string) *CreateInstanceInspectionRequest
	GetItems() []*string
	SetMetricSource(v string) *CreateInstanceInspectionRequest
	GetMetricSource() *string
	SetRegion(v string) *CreateInstanceInspectionRequest
	GetRegion() *string
	SetSource(v string) *CreateInstanceInspectionRequest
	GetSource() *string
	SetXSysomInvokeSource(v string) *CreateInstanceInspectionRequest
	GetXSysomInvokeSource() *string
}

type CreateInstanceInspectionRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-2ze5ru5rjurix7f71sxv
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// The anomaly items.
	Items []*string `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The metric source.
	//
	// example:
	//
	// sysom
	MetricSource *string `json:"metricSource,omitempty" xml:"metricSource,omitempty"`
	// The region to which the instance belongs.
	//
	// example:
	//
	// cn-shenzhen
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The source.
	//
	// example:
	//
	// console
	Source             *string `json:"source,omitempty" xml:"source,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s CreateInstanceInspectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *CreateInstanceInspectionRequest) GetInstance() *string {
	return s.Instance
}

func (s *CreateInstanceInspectionRequest) GetItems() []*string {
	return s.Items
}

func (s *CreateInstanceInspectionRequest) GetMetricSource() *string {
	return s.MetricSource
}

func (s *CreateInstanceInspectionRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateInstanceInspectionRequest) GetSource() *string {
	return s.Source
}

func (s *CreateInstanceInspectionRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *CreateInstanceInspectionRequest) SetXDebugId(v string) *CreateInstanceInspectionRequest {
	s.XDebugId = &v
	return s
}

func (s *CreateInstanceInspectionRequest) SetInstance(v string) *CreateInstanceInspectionRequest {
	s.Instance = &v
	return s
}

func (s *CreateInstanceInspectionRequest) SetItems(v []*string) *CreateInstanceInspectionRequest {
	s.Items = v
	return s
}

func (s *CreateInstanceInspectionRequest) SetMetricSource(v string) *CreateInstanceInspectionRequest {
	s.MetricSource = &v
	return s
}

func (s *CreateInstanceInspectionRequest) SetRegion(v string) *CreateInstanceInspectionRequest {
	s.Region = &v
	return s
}

func (s *CreateInstanceInspectionRequest) SetSource(v string) *CreateInstanceInspectionRequest {
	s.Source = &v
	return s
}

func (s *CreateInstanceInspectionRequest) SetXSysomInvokeSource(v string) *CreateInstanceInspectionRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *CreateInstanceInspectionRequest) Validate() error {
	return dara.Validate(s)
}
