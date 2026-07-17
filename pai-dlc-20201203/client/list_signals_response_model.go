// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSignalsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSignalsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSignalsResponse
	GetStatusCode() *int32
	SetBody(v *ListSignalsResponseBody) *ListSignalsResponse
	GetBody() *ListSignalsResponseBody
}

type ListSignalsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSignalsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSignalsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSignalsResponse) GoString() string {
	return s.String()
}

func (s *ListSignalsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSignalsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSignalsResponse) GetBody() *ListSignalsResponseBody {
	return s.Body
}

func (s *ListSignalsResponse) SetHeaders(v map[string]*string) *ListSignalsResponse {
	s.Headers = v
	return s
}

func (s *ListSignalsResponse) SetStatusCode(v int32) *ListSignalsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSignalsResponse) SetBody(v *ListSignalsResponseBody) *ListSignalsResponse {
	s.Body = v
	return s
}

func (s *ListSignalsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
