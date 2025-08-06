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
	// The name of the watermark template to which you want to change.
	//
	// 	- Only letters and digits are supported.
	//
	// 	- The name cannot exceed 128 bytes.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration information of the watermark such as the display position and special effects. The value must be a JSON string. The configuration parameters for image and text watermarks are different. For more information about the parameter structure, see [WatermarkConfig](~~98618#section-h01-44s-2lr~~).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Width":"55","Height":"55","Dx":"9","Dy":"9","ReferPos":"BottonLeft","Type":"Image"}
	WatermarkConfig *string `json:"WatermarkConfig,omitempty" xml:"WatermarkConfig,omitempty"`
	// The ID of the watermark template. You can specify only one watermark template ID. You can obtain the ID by using one of the following methods:
	//
	// 	- Obtain the watermark template ID from the response to the [AddWatermark](~~AddWatermark~~) operation that you call to create a watermark template.
	//
	// 	- Obtain the watermark template ID from the response to the [ListWatermark](~~ListWatermark~~) operation that you call to query all watermark templates within your account.
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
