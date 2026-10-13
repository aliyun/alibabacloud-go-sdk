// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductRefResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProductRefResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProductRefResponse
	GetStatusCode() *int32
	SetBody(v *GetProductRefResponseBody) *GetProductRefResponse
	GetBody() *GetProductRefResponseBody
}

type GetProductRefResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProductRefResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProductRefResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProductRefResponse) GoString() string {
	return s.String()
}

func (s *GetProductRefResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProductRefResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProductRefResponse) GetBody() *GetProductRefResponseBody {
	return s.Body
}

func (s *GetProductRefResponse) SetHeaders(v map[string]*string) *GetProductRefResponse {
	s.Headers = v
	return s
}

func (s *GetProductRefResponse) SetStatusCode(v int32) *GetProductRefResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProductRefResponse) SetBody(v *GetProductRefResponseBody) *GetProductRefResponse {
	s.Body = v
	return s
}

func (s *GetProductRefResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
