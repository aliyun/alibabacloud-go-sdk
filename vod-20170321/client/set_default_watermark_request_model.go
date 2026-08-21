// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultWatermarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWatermarkId(v string) *SetDefaultWatermarkRequest
	GetWatermarkId() *string
}

type SetDefaultWatermarkRequest struct {
	// The ID of the watermark template to set as the default. Only a single watermark template ID is supported. You can obtain the ID by using one of the following methods:
	//
	// - Call the [AddWatermark](~~AddWatermark~~) operation to add a watermark template. The ID is returned in the response.
	//
	// - Call the [ListWatermark](~~ListWatermark~~) operation to query the list of watermark templates. The ID is returned in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9bcc8bfadb843f*****09a2671d0df97
	WatermarkId *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s SetDefaultWatermarkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultWatermarkRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultWatermarkRequest) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *SetDefaultWatermarkRequest) SetWatermarkId(v string) *SetDefaultWatermarkRequest {
	s.WatermarkId = &v
	return s
}

func (s *SetDefaultWatermarkRequest) Validate() error {
	return dara.Validate(s)
}
