// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteApplicationTemplateResponseBody
	GetRequestId() *string
}

type DeleteApplicationTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApplicationTemplateResponseBody) SetRequestId(v string) *DeleteApplicationTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApplicationTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
