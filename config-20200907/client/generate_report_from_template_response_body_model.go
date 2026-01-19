// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateReportFromTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GenerateReportFromTemplateResponseBody
	GetRequestId() *string
}

type GenerateReportFromTemplateResponseBody struct {
	// example:
	//
	// 6CE4ABA1-9A57-41A9-8EA9-E8B17D4671CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateReportFromTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateReportFromTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateReportFromTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateReportFromTemplateResponseBody) SetRequestId(v string) *GenerateReportFromTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateReportFromTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
