// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContactTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContactTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContactTemplateResponseBody) *DeleteContactTemplateResponse
	GetBody() *DeleteContactTemplateResponseBody
}

type DeleteContactTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContactTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContactTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteContactTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContactTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContactTemplateResponse) GetBody() *DeleteContactTemplateResponseBody {
	return s.Body
}

func (s *DeleteContactTemplateResponse) SetHeaders(v map[string]*string) *DeleteContactTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteContactTemplateResponse) SetStatusCode(v int32) *DeleteContactTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContactTemplateResponse) SetBody(v *DeleteContactTemplateResponseBody) *DeleteContactTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteContactTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
