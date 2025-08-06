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
	// The ID of watermark template that you want to delete. You can specify only one watermark template ID. You can obtain the ID by using one of the following methods:
	//
	// 	- Obtain the watermark template ID from the response to the [AddWatermark](~~AddWatermark~~) operation that you call to create a watermark template.
	//
	// 	- Obtain the watermark template ID from the response to the [ListWatermark](~~ListWatermark~~) operation that you call to query all watermark templates within your account.
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
