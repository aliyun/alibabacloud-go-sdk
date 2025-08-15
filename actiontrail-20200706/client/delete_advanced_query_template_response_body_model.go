// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAdvancedQueryTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAdvancedQueryTemplateResponseBody
	GetRequestId() *string
}

type DeleteAdvancedQueryTemplateResponseBody struct {
	// example:
	//
	// 95F2CD1D-9BD3-564A-A74A-743FFC5E46E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAdvancedQueryTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAdvancedQueryTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAdvancedQueryTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAdvancedQueryTemplateResponseBody) SetRequestId(v string) *DeleteAdvancedQueryTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAdvancedQueryTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
