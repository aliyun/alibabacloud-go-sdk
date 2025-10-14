// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataIngestionTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataIngestionTemplateResponseBody
	GetRequestId() *string
}

type UpdateDataIngestionTemplateResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDataIngestionTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataIngestionTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataIngestionTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataIngestionTemplateResponseBody) SetRequestId(v string) *UpdateDataIngestionTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataIngestionTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
