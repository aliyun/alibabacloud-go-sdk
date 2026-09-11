// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateResourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateResourceResponseBody) *CreateResourceResponse
	GetBody() *CreateResourceResponseBody
}

type CreateResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateResourceResponse) GetBody() *CreateResourceResponseBody {
	return s.Body
}

func (s *CreateResourceResponse) SetHeaders(v map[string]*string) *CreateResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceResponse) SetStatusCode(v int32) *CreateResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceResponse) SetBody(v *CreateResourceResponseBody) *CreateResourceResponse {
	s.Body = v
	return s
}

func (s *CreateResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
