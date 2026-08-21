// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWatermarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWatermarkId(v string) *DeleteWatermarkRequest
	GetWatermarkId() *string
}

type DeleteWatermarkRequest struct {
	// The ID of the watermark template to delete. Only a single watermark template ID can be specified. You can obtain the ID by using the following methods:
	//
	// - The ID is returned after you call the [AddWatermark](~~AddWatermark~~) operation to add a watermark template.
	//
	// - The ID is returned after you call the [ListWatermark](~~ListWatermark~~) operation to query the list of watermark templates.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9bcc8bfadb843f*****09a2671d0df97
	WatermarkId *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s DeleteWatermarkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWatermarkRequest) GoString() string {
	return s.String()
}

func (s *DeleteWatermarkRequest) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *DeleteWatermarkRequest) SetWatermarkId(v string) *DeleteWatermarkRequest {
	s.WatermarkId = &v
	return s
}

func (s *DeleteWatermarkRequest) Validate() error {
	return dara.Validate(s)
}
