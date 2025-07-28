// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDefenseTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDefenseTemplateResponseBody
	GetRequestId() *string
}

type DeleteDefenseTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E2DE6F11-6FED-5909-95F2-2520B58C5BC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDefenseTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDefenseTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDefenseTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDefenseTemplateResponseBody) SetRequestId(v string) *DeleteDefenseTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDefenseTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
