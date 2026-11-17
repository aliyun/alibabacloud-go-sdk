// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebFilterTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWebFilterTemplateResponseBody
	GetRequestId() *string
}

type DeleteWebFilterTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWebFilterTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebFilterTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWebFilterTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWebFilterTemplateResponseBody) SetRequestId(v string) *DeleteWebFilterTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWebFilterTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
