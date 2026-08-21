// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWatermarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UpdateWatermarkRequest
	GetName() *string
	SetWatermarkConfig(v string) *UpdateWatermarkRequest
	GetWatermarkConfig() *string
	SetWatermarkId(v string) *UpdateWatermarkRequest
	GetWatermarkId() *string
}

type UpdateWatermarkRequest struct {
	// The new name of the watermark template.
	//
	// - Only Chinese characters, letters, and digits are supported.
	//
	// - The name can be up to 128 bytes in length.
	//
	// - UTF-8 encoding.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration information of the image and text watermark (JSON character string), including the watermark display position and watermark effect. The configuration parameters for image watermarks and text watermarks are different. For details about the parameter structure, see [WatermarkConfig](~~98618#section-h01-44s-2lr~~).
	//
	// >Modifying across templatetypes is not supported. You can invoke the [GetWatermark](~~GetWatermark~~) operation to query the type of the watermark template before modifying the configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Width":"55","Height":"55","Dx":"9","Dy":"9","ReferPos":"BottomLeft"}
	WatermarkConfig *string `json:"WatermarkConfig,omitempty" xml:"WatermarkConfig,omitempty"`
	// The ID of the image and text watermark template to modify. Only a single watermark template ID is supported. You can obtain the ID by using one of the following methods:
	//
	// - The ID is returned after you call the [AddWatermark](~~AddWatermark~~) operation to add an image and text watermark template.
	//
	// - The ID is returned after you call the [ListWatermark](~~ListWatermark~~) operation to query the list of image and text watermark templates.
	//
	// This parameter is required.
	//
	// example:
	//
	// af2afe4761992c*****bd947dae97337
	WatermarkId *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s UpdateWatermarkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWatermarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWatermarkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateWatermarkRequest) GetWatermarkConfig() *string {
	return s.WatermarkConfig
}

func (s *UpdateWatermarkRequest) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *UpdateWatermarkRequest) SetName(v string) *UpdateWatermarkRequest {
	s.Name = &v
	return s
}

func (s *UpdateWatermarkRequest) SetWatermarkConfig(v string) *UpdateWatermarkRequest {
	s.WatermarkConfig = &v
	return s
}

func (s *UpdateWatermarkRequest) SetWatermarkId(v string) *UpdateWatermarkRequest {
	s.WatermarkId = &v
	return s
}

func (s *UpdateWatermarkRequest) Validate() error {
	return dara.Validate(s)
}
