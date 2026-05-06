// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelProviderTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteModelProviderTemplateResponseBody
	GetRequestId() *string
}

type DeleteModelProviderTemplateResponseBody struct {
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteModelProviderTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelProviderTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelProviderTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteModelProviderTemplateResponseBody) SetRequestId(v string) *DeleteModelProviderTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteModelProviderTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
