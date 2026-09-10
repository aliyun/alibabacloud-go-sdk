// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataCheckTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataCheckTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataCheckTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataCheckTemplateResponseBody) *DeleteDataCheckTemplateResponse
	GetBody() *DeleteDataCheckTemplateResponseBody
}

type DeleteDataCheckTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataCheckTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataCheckTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataCheckTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataCheckTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataCheckTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataCheckTemplateResponse) GetBody() *DeleteDataCheckTemplateResponseBody {
	return s.Body
}

func (s *DeleteDataCheckTemplateResponse) SetHeaders(v map[string]*string) *DeleteDataCheckTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataCheckTemplateResponse) SetStatusCode(v int32) *DeleteDataCheckTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataCheckTemplateResponse) SetBody(v *DeleteDataCheckTemplateResponseBody) *DeleteDataCheckTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteDataCheckTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
