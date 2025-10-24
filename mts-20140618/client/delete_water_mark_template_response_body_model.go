// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWaterMarkTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWaterMarkTemplateResponseBody
	GetRequestId() *string
	SetWaterMarkTemplateId(v string) *DeleteWaterMarkTemplateResponseBody
	GetWaterMarkTemplateId() *string
}

type DeleteWaterMarkTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 030E2671-806A-52AF-A93C-DA8E308603A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the deleted watermark template.
	//
	// example:
	//
	// 3780bd69b2b74540bc7b1096f564****
	WaterMarkTemplateId *string `json:"WaterMarkTemplateId,omitempty" xml:"WaterMarkTemplateId,omitempty"`
}

func (s DeleteWaterMarkTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWaterMarkTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWaterMarkTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWaterMarkTemplateResponseBody) GetWaterMarkTemplateId() *string {
	return s.WaterMarkTemplateId
}

func (s *DeleteWaterMarkTemplateResponseBody) SetRequestId(v string) *DeleteWaterMarkTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWaterMarkTemplateResponseBody) SetWaterMarkTemplateId(v string) *DeleteWaterMarkTemplateResponseBody {
	s.WaterMarkTemplateId = &v
	return s
}

func (s *DeleteWaterMarkTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
