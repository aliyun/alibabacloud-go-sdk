// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceMirrorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRatio(v int32) *UpdateServiceMirrorRequest
	GetRatio() *int32
	SetTarget(v []*string) *UpdateServiceMirrorRequest
	GetTarget() []*string
}

type UpdateServiceMirrorRequest struct {
	// The percentage of traffic that you want to mirror. Valid values: 0 to 100.
	//
	// example:
	//
	// 30
	Ratio *int32 `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// The service instances.
	Target []*string `json:"Target,omitempty" xml:"Target,omitempty" type:"Repeated"`
}

func (s UpdateServiceMirrorRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceMirrorRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceMirrorRequest) GetRatio() *int32 {
	return s.Ratio
}

func (s *UpdateServiceMirrorRequest) GetTarget() []*string {
	return s.Target
}

func (s *UpdateServiceMirrorRequest) SetRatio(v int32) *UpdateServiceMirrorRequest {
	s.Ratio = &v
	return s
}

func (s *UpdateServiceMirrorRequest) SetTarget(v []*string) *UpdateServiceMirrorRequest {
	s.Target = v
	return s
}

func (s *UpdateServiceMirrorRequest) Validate() error {
	return dara.Validate(s)
}
