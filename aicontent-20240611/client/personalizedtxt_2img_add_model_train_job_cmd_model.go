// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedtxt2imgAddModelTrainJobCmd interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v []*string) *Personalizedtxt2imgAddModelTrainJobCmd
	GetImageUrl() []*string
	SetName(v string) *Personalizedtxt2imgAddModelTrainJobCmd
	GetName() *string
	SetObjectType(v string) *Personalizedtxt2imgAddModelTrainJobCmd
	GetObjectType() *string
}

type Personalizedtxt2imgAddModelTrainJobCmd struct {
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
}

func (s Personalizedtxt2imgAddModelTrainJobCmd) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgAddModelTrainJobCmd) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgAddModelTrainJobCmd) GetImageUrl() []*string {
	return s.ImageUrl
}

func (s *Personalizedtxt2imgAddModelTrainJobCmd) GetName() *string {
	return s.Name
}

func (s *Personalizedtxt2imgAddModelTrainJobCmd) GetObjectType() *string {
	return s.ObjectType
}

func (s *Personalizedtxt2imgAddModelTrainJobCmd) SetImageUrl(v []*string) *Personalizedtxt2imgAddModelTrainJobCmd {
	s.ImageUrl = v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobCmd) SetName(v string) *Personalizedtxt2imgAddModelTrainJobCmd {
	s.Name = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobCmd) SetObjectType(v string) *Personalizedtxt2imgAddModelTrainJobCmd {
	s.ObjectType = &v
	return s
}

func (s *Personalizedtxt2imgAddModelTrainJobCmd) Validate() error {
	return dara.Validate(s)
}
