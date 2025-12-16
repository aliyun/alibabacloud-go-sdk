// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionCurrentVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *GetFunctionCurrentVersionRequest
	GetCategory() *string
	SetDomain(v string) *GetFunctionCurrentVersionRequest
	GetDomain() *string
	SetFunctionType(v string) *GetFunctionCurrentVersionRequest
	GetFunctionType() *string
	SetModelType(v string) *GetFunctionCurrentVersionRequest
	GetModelType() *string
}

type GetFunctionCurrentVersionRequest struct {
	// The category. By default, this parameter is left empty.
	//
	// example:
	//
	// general
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The industry. By default, this parameter is left empty, which indicates General-purpose Edition.
	//
	// example:
	//
	// ecommerce
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The type of the feature. Valid values:
	//
	// 	- PAAS. This is the default value.
	//
	// 	- SAAS.
	//
	// example:
	//
	// PAAS
	FunctionType *string `json:"functionType,omitempty" xml:"functionType,omitempty"`
	// The type of the model. The following features correspond to different model types:
	//
	// 	- CTR model: tf_checkpoint
	//
	// 	- Popularity model: pop
	//
	// 	- Category model: offline_inference
	//
	// 	- Hotword model: offline_inference
	//
	// 	- Shading model: offline_inference
	//
	// 	- Drop-down suggestion model: offline_inference
	//
	// 	- Word segmentation model: text
	//
	// 	- Word weight model: tf_checkpoint
	//
	// This parameter is required.
	//
	// example:
	//
	// tf_checkpoint
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
}

func (s GetFunctionCurrentVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionCurrentVersionRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionRequest) GetCategory() *string {
	return s.Category
}

func (s *GetFunctionCurrentVersionRequest) GetDomain() *string {
	return s.Domain
}

func (s *GetFunctionCurrentVersionRequest) GetFunctionType() *string {
	return s.FunctionType
}

func (s *GetFunctionCurrentVersionRequest) GetModelType() *string {
	return s.ModelType
}

func (s *GetFunctionCurrentVersionRequest) SetCategory(v string) *GetFunctionCurrentVersionRequest {
	s.Category = &v
	return s
}

func (s *GetFunctionCurrentVersionRequest) SetDomain(v string) *GetFunctionCurrentVersionRequest {
	s.Domain = &v
	return s
}

func (s *GetFunctionCurrentVersionRequest) SetFunctionType(v string) *GetFunctionCurrentVersionRequest {
	s.FunctionType = &v
	return s
}

func (s *GetFunctionCurrentVersionRequest) SetModelType(v string) *GetFunctionCurrentVersionRequest {
	s.ModelType = &v
	return s
}

func (s *GetFunctionCurrentVersionRequest) Validate() error {
	return dara.Validate(s)
}
