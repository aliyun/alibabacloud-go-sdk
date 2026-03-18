// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedtxt2imgAddModelTrainJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v []*string) *Personalizedtxt2imgAddModelTrainJobRequest
	GetImageUrl() []*string
	SetName(v string) *Personalizedtxt2imgAddModelTrainJobRequest
	GetName() *string
	SetObjectType(v string) *Personalizedtxt2imgAddModelTrainJobRequest
	GetObjectType() *string
	SetTrainSteps(v int32) *Personalizedtxt2imgAddModelTrainJobRequest
	GetTrainSteps() *int32
}

type Personalizedtxt2imgAddModelTrainJobRequest struct {
	// This parameter is required.
	ImageUrl []*string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 熊猫图片生成
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dog
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// example:
	//
	// 800
	TrainSteps *int32 `json:"trainSteps,omitempty" xml:"trainSteps,omitempty"`
}

func (s Personalizedtxt2imgAddModelTrainJobRequest) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgAddModelTrainJobRequest) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddModelTrainJobRequest) GetImageUrl() []*string {
	return s.ImageUrl
}

func (s *Personalizedtxt2imgAddModelTrainJobRequest) GetName() *string {
	return s.Name
}

func (s *Personalizedtxt2imgAddModelTrainJobRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *Personalizedtxt2imgAddModelTrainJobRequest) GetTrainSteps() *int32 {
	return s.TrainSteps
}

func (s *Personalizedtxt2imgAddModelTrainJobRequest) SetImageUrl(v []*string) *Personalizedtxt2imgAddModelTrainJobRequest {
	s.ImageUrl = v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobRequest) SetName(v string) *Personalizedtxt2imgAddModelTrainJobRequest {
	s.Name = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobRequest) SetObjectType(v string) *Personalizedtxt2imgAddModelTrainJobRequest {
	s.ObjectType = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobRequest) SetTrainSteps(v int32) *Personalizedtxt2imgAddModelTrainJobRequest {
	s.TrainSteps = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobRequest) Validate() error {
	return dara.Validate(s)
}
