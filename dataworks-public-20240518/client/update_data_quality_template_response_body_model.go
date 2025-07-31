// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataQualityTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataQualityTemplateResponseBody
	GetSuccess() *bool
}

type UpdateDataQualityTemplateResponseBody struct {
	// example:
	//
	// 0bc14115***159376359
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataQualityTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataQualityTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataQualityTemplateResponseBody) SetRequestId(v string) *UpdateDataQualityTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataQualityTemplateResponseBody) SetSuccess(v bool) *UpdateDataQualityTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataQualityTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
