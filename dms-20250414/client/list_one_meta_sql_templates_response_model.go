// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOneMetaSqlTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOneMetaSqlTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOneMetaSqlTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListOneMetaSqlTemplatesResponseBody) *ListOneMetaSqlTemplatesResponse
	GetBody() *ListOneMetaSqlTemplatesResponseBody
}

type ListOneMetaSqlTemplatesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOneMetaSqlTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOneMetaSqlTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOneMetaSqlTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListOneMetaSqlTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOneMetaSqlTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOneMetaSqlTemplatesResponse) GetBody() *ListOneMetaSqlTemplatesResponseBody {
	return s.Body
}

func (s *ListOneMetaSqlTemplatesResponse) SetHeaders(v map[string]*string) *ListOneMetaSqlTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListOneMetaSqlTemplatesResponse) SetStatusCode(v int32) *ListOneMetaSqlTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOneMetaSqlTemplatesResponse) SetBody(v *ListOneMetaSqlTemplatesResponseBody) *ListOneMetaSqlTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListOneMetaSqlTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
