// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWatermarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWatermarkId(v string) *GetWatermarkRequest
	GetWatermarkId() *string
}

type GetWatermarkRequest struct {
	// The ID of the watermark template. You can specify only one watermark template ID. You can obtain the ID by using one of the following methods:
	//
	// 	- Obtain the watermark template ID from the response to the [AddWatermark](~~AddWatermark~~) operation that you call to create a watermark template.
	//
	// 	- Obtain the watermark template ID from the response to the [ListWatermark](~~ListWatermark~~) operation that you call to query all watermarks within your account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9bcc8bfadb843f*****09a2671d0df97
	WatermarkId *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s GetWatermarkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarkRequest) GoString() string {
	return s.String()
}

func (s *GetWatermarkRequest) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *GetWatermarkRequest) SetWatermarkId(v string) *GetWatermarkRequest {
	s.WatermarkId = &v
	return s
}

func (s *GetWatermarkRequest) Validate() error {
	return dara.Validate(s)
}
