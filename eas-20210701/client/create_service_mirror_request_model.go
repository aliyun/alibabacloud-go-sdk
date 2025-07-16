// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceMirrorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRatio(v int32) *CreateServiceMirrorRequest
	GetRatio() *int32
	SetTarget(v []*string) *CreateServiceMirrorRequest
	GetTarget() []*string
}

type CreateServiceMirrorRequest struct {
	// The percentage of the traffic that is mirrored to the destination service. Valid values: 0 to 100.
	//
	// example:
	//
	// 30
	Ratio *int32 `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// The instances.
	Target []*string `json:"Target,omitempty" xml:"Target,omitempty" type:"Repeated"`
}

func (s CreateServiceMirrorRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceMirrorRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceMirrorRequest) GetRatio() *int32 {
	return s.Ratio
}

func (s *CreateServiceMirrorRequest) GetTarget() []*string {
	return s.Target
}

func (s *CreateServiceMirrorRequest) SetRatio(v int32) *CreateServiceMirrorRequest {
	s.Ratio = &v
	return s
}

func (s *CreateServiceMirrorRequest) SetTarget(v []*string) *CreateServiceMirrorRequest {
	s.Target = v
	return s
}

func (s *CreateServiceMirrorRequest) Validate() error {
	return dara.Validate(s)
}
