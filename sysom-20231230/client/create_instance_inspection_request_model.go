// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceInspectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstance(v string) *CreateInstanceInspectionRequest
	GetInstance() *string
	SetItems(v []*string) *CreateInstanceInspectionRequest
	GetItems() []*string
	SetRegion(v string) *CreateInstanceInspectionRequest
	GetRegion() *string
	SetSource(v string) *CreateInstanceInspectionRequest
	GetSource() *string
}

type CreateInstanceInspectionRequest struct {
	// example:
	//
	// i-2ze5ru5rjurix7f71sxv
	Instance *string   `json:"instance,omitempty" xml:"instance,omitempty"`
	Items    []*string `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// cn-shenzhen
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// console
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s CreateInstanceInspectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionRequest) GetInstance() *string {
	return s.Instance
}

func (s *CreateInstanceInspectionRequest) GetItems() []*string {
	return s.Items
}

func (s *CreateInstanceInspectionRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateInstanceInspectionRequest) GetSource() *string {
	return s.Source
}

func (s *CreateInstanceInspectionRequest) SetInstance(v string) *CreateInstanceInspectionRequest {
	s.Instance = &v
	return s
}

func (s *CreateInstanceInspectionRequest) SetItems(v []*string) *CreateInstanceInspectionRequest {
	s.Items = v
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

func (s *CreateInstanceInspectionRequest) Validate() error {
	return dara.Validate(s)
}
