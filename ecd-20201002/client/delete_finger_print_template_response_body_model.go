// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFingerPrintTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFingerPrintTemplateResponseBody
	GetRequestId() *string
}

type DeleteFingerPrintTemplateResponseBody struct {
	// example:
	//
	// 134BD0B2-B848-5743-9CE2-C1FD3D5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFingerPrintTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFingerPrintTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFingerPrintTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFingerPrintTemplateResponseBody) SetRequestId(v string) *DeleteFingerPrintTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFingerPrintTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
