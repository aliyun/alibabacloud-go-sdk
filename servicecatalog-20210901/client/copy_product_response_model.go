// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyProductResponse
	GetStatusCode() *int32
	SetBody(v *CopyProductResponseBody) *CopyProductResponse
	GetBody() *CopyProductResponseBody
}

type CopyProductResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyProductResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyProductResponse) GoString() string {
	return s.String()
}

func (s *CopyProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyProductResponse) GetBody() *CopyProductResponseBody {
	return s.Body
}

func (s *CopyProductResponse) SetHeaders(v map[string]*string) *CopyProductResponse {
	s.Headers = v
	return s
}

func (s *CopyProductResponse) SetStatusCode(v int32) *CopyProductResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyProductResponse) SetBody(v *CopyProductResponseBody) *CopyProductResponse {
	s.Body = v
	return s
}

func (s *CopyProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
