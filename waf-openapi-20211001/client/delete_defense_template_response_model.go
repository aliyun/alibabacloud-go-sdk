// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDefenseTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDefenseTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDefenseTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDefenseTemplateResponseBody) *DeleteDefenseTemplateResponse
	GetBody() *DeleteDefenseTemplateResponseBody
}

type DeleteDefenseTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDefenseTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDefenseTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDefenseTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteDefenseTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDefenseTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDefenseTemplateResponse) GetBody() *DeleteDefenseTemplateResponseBody {
	return s.Body
}

func (s *DeleteDefenseTemplateResponse) SetHeaders(v map[string]*string) *DeleteDefenseTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteDefenseTemplateResponse) SetStatusCode(v int32) *DeleteDefenseTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDefenseTemplateResponse) SetBody(v *DeleteDefenseTemplateResponseBody) *DeleteDefenseTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteDefenseTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
