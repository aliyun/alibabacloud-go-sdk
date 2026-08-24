// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOneMetaSqlTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOneMetaSqlTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOneMetaSqlTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOneMetaSqlTemplateResponseBody) *DeleteOneMetaSqlTemplateResponse
	GetBody() *DeleteOneMetaSqlTemplateResponseBody
}

type DeleteOneMetaSqlTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOneMetaSqlTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOneMetaSqlTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOneMetaSqlTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteOneMetaSqlTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOneMetaSqlTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOneMetaSqlTemplateResponse) GetBody() *DeleteOneMetaSqlTemplateResponseBody {
	return s.Body
}

func (s *DeleteOneMetaSqlTemplateResponse) SetHeaders(v map[string]*string) *DeleteOneMetaSqlTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteOneMetaSqlTemplateResponse) SetStatusCode(v int32) *DeleteOneMetaSqlTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOneMetaSqlTemplateResponse) SetBody(v *DeleteOneMetaSqlTemplateResponseBody) *DeleteOneMetaSqlTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteOneMetaSqlTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
