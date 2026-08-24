// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOneMetaSqlTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOneMetaSqlTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOneMetaSqlTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateOneMetaSqlTemplateResponseBody) *CreateOneMetaSqlTemplateResponse
	GetBody() *CreateOneMetaSqlTemplateResponseBody
}

type CreateOneMetaSqlTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOneMetaSqlTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOneMetaSqlTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOneMetaSqlTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateOneMetaSqlTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOneMetaSqlTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOneMetaSqlTemplateResponse) GetBody() *CreateOneMetaSqlTemplateResponseBody {
	return s.Body
}

func (s *CreateOneMetaSqlTemplateResponse) SetHeaders(v map[string]*string) *CreateOneMetaSqlTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateOneMetaSqlTemplateResponse) SetStatusCode(v int32) *CreateOneMetaSqlTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOneMetaSqlTemplateResponse) SetBody(v *CreateOneMetaSqlTemplateResponseBody) *CreateOneMetaSqlTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateOneMetaSqlTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
