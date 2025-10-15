// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReIndexResponse
	GetStatusCode() *int32
	SetBody(v *ReIndexResponseBody) *ReIndexResponse
	GetBody() *ReIndexResponseBody
}

type ReIndexResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s ReIndexResponse) GoString() string {
	return s.String()
}

func (s *ReIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReIndexResponse) GetBody() *ReIndexResponseBody {
	return s.Body
}

func (s *ReIndexResponse) SetHeaders(v map[string]*string) *ReIndexResponse {
	s.Headers = v
	return s
}

func (s *ReIndexResponse) SetStatusCode(v int32) *ReIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *ReIndexResponse) SetBody(v *ReIndexResponseBody) *ReIndexResponse {
	s.Body = v
	return s
}

func (s *ReIndexResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
