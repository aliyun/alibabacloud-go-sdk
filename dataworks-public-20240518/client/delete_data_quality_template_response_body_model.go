// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataQualityTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataQualityTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataQualityTemplateResponseBody
	GetSuccess() *bool
}

type DeleteDataQualityTemplateResponseBody struct {
	// example:
	//
	// 0bc14115****159376359
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataQualityTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataQualityTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataQualityTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataQualityTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataQualityTemplateResponseBody) SetRequestId(v string) *DeleteDataQualityTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataQualityTemplateResponseBody) SetSuccess(v bool) *DeleteDataQualityTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataQualityTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
