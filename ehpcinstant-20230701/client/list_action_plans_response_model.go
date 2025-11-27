// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActionPlansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListActionPlansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListActionPlansResponse
	GetStatusCode() *int32
	SetBody(v *ListActionPlansResponseBody) *ListActionPlansResponse
	GetBody() *ListActionPlansResponseBody
}

type ListActionPlansResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListActionPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListActionPlansResponse) String() string {
	return dara.Prettify(s)
}

func (s ListActionPlansResponse) GoString() string {
	return s.String()
}

func (s *ListActionPlansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListActionPlansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListActionPlansResponse) GetBody() *ListActionPlansResponseBody {
	return s.Body
}

func (s *ListActionPlansResponse) SetHeaders(v map[string]*string) *ListActionPlansResponse {
	s.Headers = v
	return s
}

func (s *ListActionPlansResponse) SetStatusCode(v int32) *ListActionPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *ListActionPlansResponse) SetBody(v *ListActionPlansResponseBody) *ListActionPlansResponse {
	s.Body = v
	return s
}

func (s *ListActionPlansResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
