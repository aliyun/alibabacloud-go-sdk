// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunInstanceInspectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RunInstanceInspectionRequest
	GetClientToken() *string
	SetItemOverrides(v []*RunInstanceInspectionRequestItemOverrides) *RunInstanceInspectionRequest
	GetItemOverrides() []*RunInstanceInspectionRequestItemOverrides
	SetRegionId(v string) *RunInstanceInspectionRequest
	GetRegionId() *string
}

type RunInstanceInspectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// inspection-request-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// [{"templateCode":"rmq.traffic-capacity","parameters":{"lookback":"6h"}}]
	ItemOverrides []*RunInstanceInspectionRequestItemOverrides `json:"itemOverrides,omitempty" xml:"itemOverrides,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s RunInstanceInspectionRequest) String() string {
	return dara.Prettify(s)
}

func (s RunInstanceInspectionRequest) GoString() string {
	return s.String()
}

func (s *RunInstanceInspectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RunInstanceInspectionRequest) GetItemOverrides() []*RunInstanceInspectionRequestItemOverrides {
	return s.ItemOverrides
}

func (s *RunInstanceInspectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunInstanceInspectionRequest) SetClientToken(v string) *RunInstanceInspectionRequest {
	s.ClientToken = &v
	return s
}

func (s *RunInstanceInspectionRequest) SetItemOverrides(v []*RunInstanceInspectionRequestItemOverrides) *RunInstanceInspectionRequest {
	s.ItemOverrides = v
	return s
}

func (s *RunInstanceInspectionRequest) SetRegionId(v string) *RunInstanceInspectionRequest {
	s.RegionId = &v
	return s
}

func (s *RunInstanceInspectionRequest) Validate() error {
	if s.ItemOverrides != nil {
		for _, item := range s.ItemOverrides {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunInstanceInspectionRequestItemOverrides struct {
	// This parameter is required.
	//
	// example:
	//
	// {"lookback":"30m"}
	Parameters interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rmq.consumer-health
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
}

func (s RunInstanceInspectionRequestItemOverrides) String() string {
	return dara.Prettify(s)
}

func (s RunInstanceInspectionRequestItemOverrides) GoString() string {
	return s.String()
}

func (s *RunInstanceInspectionRequestItemOverrides) GetParameters() interface{} {
	return s.Parameters
}

func (s *RunInstanceInspectionRequestItemOverrides) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *RunInstanceInspectionRequestItemOverrides) SetParameters(v interface{}) *RunInstanceInspectionRequestItemOverrides {
	s.Parameters = v
	return s
}

func (s *RunInstanceInspectionRequestItemOverrides) SetTemplateCode(v string) *RunInstanceInspectionRequestItemOverrides {
	s.TemplateCode = &v
	return s
}

func (s *RunInstanceInspectionRequestItemOverrides) Validate() error {
	return dara.Validate(s)
}
