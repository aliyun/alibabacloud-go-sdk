// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalGuardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MultiModalGuardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MultiModalGuardResponse
	GetStatusCode() *int32
	SetBody(v *MultiModalGuardResponseBody) *MultiModalGuardResponse
	GetBody() *MultiModalGuardResponseBody
}

type MultiModalGuardResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MultiModalGuardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MultiModalGuardResponse) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardResponse) GoString() string {
	return s.String()
}

func (s *MultiModalGuardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MultiModalGuardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MultiModalGuardResponse) GetBody() *MultiModalGuardResponseBody {
	return s.Body
}

func (s *MultiModalGuardResponse) SetHeaders(v map[string]*string) *MultiModalGuardResponse {
	s.Headers = v
	return s
}

func (s *MultiModalGuardResponse) SetStatusCode(v int32) *MultiModalGuardResponse {
	s.StatusCode = &v
	return s
}

func (s *MultiModalGuardResponse) SetBody(v *MultiModalGuardResponseBody) *MultiModalGuardResponse {
	s.Body = v
	return s
}

func (s *MultiModalGuardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
