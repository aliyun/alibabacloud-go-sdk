// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTemplateResponseBody
	GetRequestId() *string
}

type DeleteTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2075899A-585D-4A41-A9B2-28DA8534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTemplateResponseBody) SetRequestId(v string) *DeleteTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
