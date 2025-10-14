// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataSourceTemplateResponseBody
	GetRequestId() *string
}

type UpdateDataSourceTemplateResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDataSourceTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataSourceTemplateResponseBody) SetRequestId(v string) *UpdateDataSourceTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataSourceTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
