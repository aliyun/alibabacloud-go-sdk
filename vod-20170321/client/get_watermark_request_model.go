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
	// The ID of the watermark template to query. Only a single watermark template ID is supported. You can obtain the ID by using one of the following methods:
	//
	// - Call the [AddWatermark](~~AddWatermark~~) operation to add a watermark template. The ID is returned.
	//
	// - Call the [ListWatermark](~~ListWatermark~~) operation to query the list of watermark templates. The ID is returned.
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
