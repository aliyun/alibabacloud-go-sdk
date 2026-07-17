// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListElasticPlansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListElasticPlansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListElasticPlansResponse
	GetStatusCode() *int32
	SetBody(v *ListElasticPlansResponseBody) *ListElasticPlansResponse
	GetBody() *ListElasticPlansResponseBody
}

type ListElasticPlansResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListElasticPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListElasticPlansResponse) String() string {
	return dara.Prettify(s)
}

func (s ListElasticPlansResponse) GoString() string {
	return s.String()
}

func (s *ListElasticPlansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListElasticPlansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListElasticPlansResponse) GetBody() *ListElasticPlansResponseBody {
	return s.Body
}

func (s *ListElasticPlansResponse) SetHeaders(v map[string]*string) *ListElasticPlansResponse {
	s.Headers = v
	return s
}

func (s *ListElasticPlansResponse) SetStatusCode(v int32) *ListElasticPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *ListElasticPlansResponse) SetBody(v *ListElasticPlansResponseBody) *ListElasticPlansResponse {
	s.Body = v
	return s
}

func (s *ListElasticPlansResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
