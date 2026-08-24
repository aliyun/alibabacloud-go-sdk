// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOneMetaSqlTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOneMetaSqlTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOneMetaSqlTemplateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOneMetaSqlTemplateResponseBody) *UpdateOneMetaSqlTemplateResponse
	GetBody() *UpdateOneMetaSqlTemplateResponseBody
}

type UpdateOneMetaSqlTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOneMetaSqlTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOneMetaSqlTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOneMetaSqlTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateOneMetaSqlTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOneMetaSqlTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOneMetaSqlTemplateResponse) GetBody() *UpdateOneMetaSqlTemplateResponseBody {
	return s.Body
}

func (s *UpdateOneMetaSqlTemplateResponse) SetHeaders(v map[string]*string) *UpdateOneMetaSqlTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateOneMetaSqlTemplateResponse) SetStatusCode(v int32) *UpdateOneMetaSqlTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOneMetaSqlTemplateResponse) SetBody(v *UpdateOneMetaSqlTemplateResponseBody) *UpdateOneMetaSqlTemplateResponse {
	s.Body = v
	return s
}

func (s *UpdateOneMetaSqlTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
