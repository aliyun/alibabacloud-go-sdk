// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedtxt2imgQueryImageAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncodeFormat(v string) *Personalizedtxt2imgQueryImageAssetRequest
	GetEncodeFormat() *string
	SetImageId(v string) *Personalizedtxt2imgQueryImageAssetRequest
	GetImageId() *string
	SetModelId(v string) *Personalizedtxt2imgQueryImageAssetRequest
	GetModelId() *string
	SetPromptId(v string) *Personalizedtxt2imgQueryImageAssetRequest
	GetPromptId() *string
}

type Personalizedtxt2imgQueryImageAssetRequest struct {
	// example:
	//
	// base64
	EncodeFormat *string `json:"encodeFormat,omitempty" xml:"encodeFormat,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0000.png
	ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// girl-xxxx-xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxx
	PromptId *string `json:"promptId,omitempty" xml:"promptId,omitempty"`
}

func (s Personalizedtxt2imgQueryImageAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgQueryImageAssetRequest) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryImageAssetRequest) GetEncodeFormat() *string {
	return s.EncodeFormat
}

func (s *Personalizedtxt2imgQueryImageAssetRequest) GetImageId() *string {
	return s.ImageId
}

func (s *Personalizedtxt2imgQueryImageAssetRequest) GetModelId() *string {
	return s.ModelId
}

func (s *Personalizedtxt2imgQueryImageAssetRequest) GetPromptId() *string {
	return s.PromptId
}

func (s *Personalizedtxt2imgQueryImageAssetRequest) SetEncodeFormat(v string) *Personalizedtxt2imgQueryImageAssetRequest {
	s.EncodeFormat = &v
	return s
}

func (s *Personalizedtxt2imgQueryImageAssetRequest) SetImageId(v string) *Personalizedtxt2imgQueryImageAssetRequest {
	s.ImageId = &v
	return s
}

func (s *Personalizedtxt2imgQueryImageAssetRequest) SetModelId(v string) *Personalizedtxt2imgQueryImageAssetRequest {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgQueryImageAssetRequest) SetPromptId(v string) *Personalizedtxt2imgQueryImageAssetRequest {
	s.PromptId = &v
	return s
}

func (s *Personalizedtxt2imgQueryImageAssetRequest) Validate() error {
	return dara.Validate(s)
}
