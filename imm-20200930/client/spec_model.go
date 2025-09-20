// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSpec interface {
	dara.Model
	String() string
	GoString() string
	SetBackbone(v *CustomParams) *Spec
	GetBackbone() *CustomParams
	SetClassNum(v int64) *Spec
	GetClassNum() *int64
	SetHead(v *CustomParams) *Spec
	GetHead() *CustomParams
	SetInputChannel(v int64) *Spec
	GetInputChannel() *int64
	SetLoss(v *CustomParams) *Spec
	GetLoss() *CustomParams
	SetName(v string) *Spec
	GetName() *string
	SetNeck(v *CustomParams) *Spec
	GetNeck() *CustomParams
	SetNumLandmarks(v int64) *Spec
	GetNumLandmarks() *int64
	SetPretrainedPath(v string) *Spec
	GetPretrainedPath() *string
}

type Spec struct {
	Backbone *CustomParams `json:"Backbone,omitempty" xml:"Backbone,omitempty"`
	// example:
	//
	// 10
	ClassNum *int64        `json:"ClassNum,omitempty" xml:"ClassNum,omitempty"`
	Head     *CustomParams `json:"Head,omitempty" xml:"Head,omitempty"`
	// example:
	//
	// 3
	InputChannel *int64        `json:"InputChannel,omitempty" xml:"InputChannel,omitempty"`
	Loss         *CustomParams `json:"Loss,omitempty" xml:"Loss,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ClsResNet
	Name *string       `json:"Name,omitempty" xml:"Name,omitempty"`
	Neck *CustomParams `json:"Neck,omitempty" xml:"Neck,omitempty"`
	// example:
	//
	// 5
	NumLandmarks *int64 `json:"NumLandmarks,omitempty" xml:"NumLandmarks,omitempty"`
	// example:
	//
	// oss://bucket/abc/xxx.json
	PretrainedPath *string `json:"PretrainedPath,omitempty" xml:"PretrainedPath,omitempty"`
}

func (s Spec) String() string {
	return dara.Prettify(s)
}

func (s Spec) GoString() string {
	return s.String()
}

func (s *Spec) GetBackbone() *CustomParams {
	return s.Backbone
}

func (s *Spec) GetClassNum() *int64 {
	return s.ClassNum
}

func (s *Spec) GetHead() *CustomParams {
	return s.Head
}

func (s *Spec) GetInputChannel() *int64 {
	return s.InputChannel
}

func (s *Spec) GetLoss() *CustomParams {
	return s.Loss
}

func (s *Spec) GetName() *string {
	return s.Name
}

func (s *Spec) GetNeck() *CustomParams {
	return s.Neck
}

func (s *Spec) GetNumLandmarks() *int64 {
	return s.NumLandmarks
}

func (s *Spec) GetPretrainedPath() *string {
	return s.PretrainedPath
}

func (s *Spec) SetBackbone(v *CustomParams) *Spec {
	s.Backbone = v
	return s
}

func (s *Spec) SetClassNum(v int64) *Spec {
	s.ClassNum = &v
	return s
}

func (s *Spec) SetHead(v *CustomParams) *Spec {
	s.Head = v
	return s
}

func (s *Spec) SetInputChannel(v int64) *Spec {
	s.InputChannel = &v
	return s
}

func (s *Spec) SetLoss(v *CustomParams) *Spec {
	s.Loss = v
	return s
}

func (s *Spec) SetName(v string) *Spec {
	s.Name = &v
	return s
}

func (s *Spec) SetNeck(v *CustomParams) *Spec {
	s.Neck = v
	return s
}

func (s *Spec) SetNumLandmarks(v int64) *Spec {
	s.NumLandmarks = &v
	return s
}

func (s *Spec) SetPretrainedPath(v string) *Spec {
	s.PretrainedPath = &v
	return s
}

func (s *Spec) Validate() error {
	return dara.Validate(s)
}
