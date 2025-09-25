// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectFaceAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *DetectFaceAttributesRequest
	GetBizType() *string
	SetMaterialValue(v string) *DetectFaceAttributesRequest
	GetMaterialValue() *string
}

type DetectFaceAttributesRequest struct {
	// Identifier for the business scenario using real-person authentication services.
	//
	// example:
	//
	// RPBasicTest
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The photo to be detected, see the instructions for uploading image addresses for format description. A maximum of 5 faces can be detected in a single image.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://image-demo.img-cn-hangzhou.aliyuncs.com/example.jpg
	MaterialValue *string `json:"MaterialValue,omitempty" xml:"MaterialValue,omitempty"`
}

func (s DetectFaceAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectFaceAttributesRequest) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesRequest) GetBizType() *string {
	return s.BizType
}

func (s *DetectFaceAttributesRequest) GetMaterialValue() *string {
	return s.MaterialValue
}

func (s *DetectFaceAttributesRequest) SetBizType(v string) *DetectFaceAttributesRequest {
	s.BizType = &v
	return s
}

func (s *DetectFaceAttributesRequest) SetMaterialValue(v string) *DetectFaceAttributesRequest {
	s.MaterialValue = &v
	return s
}

func (s *DetectFaceAttributesRequest) Validate() error {
	return dara.Validate(s)
}
