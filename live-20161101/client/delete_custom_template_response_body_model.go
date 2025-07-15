// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomTemplateResponseBody
	GetRequestId() *string
}

type DeleteCustomTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// BC1E78D3-FA8B-4457-DEE2-6093E1232254
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomTemplateResponseBody) SetRequestId(v string) *DeleteCustomTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
