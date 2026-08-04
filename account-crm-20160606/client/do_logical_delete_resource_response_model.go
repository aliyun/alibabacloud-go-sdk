// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDoLogicalDeleteResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DoLogicalDeleteResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DoLogicalDeleteResourceResponse
	GetStatusCode() *int32
	SetBody(v *DoLogicalDeleteResourceResponseBody) *DoLogicalDeleteResourceResponse
	GetBody() *DoLogicalDeleteResourceResponseBody
}

type DoLogicalDeleteResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DoLogicalDeleteResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DoLogicalDeleteResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DoLogicalDeleteResourceResponse) GoString() string {
	return s.String()
}

func (s *DoLogicalDeleteResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DoLogicalDeleteResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DoLogicalDeleteResourceResponse) GetBody() *DoLogicalDeleteResourceResponseBody {
	return s.Body
}

func (s *DoLogicalDeleteResourceResponse) SetHeaders(v map[string]*string) *DoLogicalDeleteResourceResponse {
	s.Headers = v
	return s
}

func (s *DoLogicalDeleteResourceResponse) SetStatusCode(v int32) *DoLogicalDeleteResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DoLogicalDeleteResourceResponse) SetBody(v *DoLogicalDeleteResourceResponseBody) *DoLogicalDeleteResourceResponse {
	s.Body = v
	return s
}

func (s *DoLogicalDeleteResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
