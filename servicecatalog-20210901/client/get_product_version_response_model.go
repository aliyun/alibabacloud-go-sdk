// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProductVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProductVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetProductVersionResponseBody) *GetProductVersionResponse
	GetBody() *GetProductVersionResponseBody
}

type GetProductVersionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProductVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProductVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProductVersionResponse) GoString() string {
	return s.String()
}

func (s *GetProductVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProductVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProductVersionResponse) GetBody() *GetProductVersionResponseBody {
	return s.Body
}

func (s *GetProductVersionResponse) SetHeaders(v map[string]*string) *GetProductVersionResponse {
	s.Headers = v
	return s
}

func (s *GetProductVersionResponse) SetStatusCode(v int32) *GetProductVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProductVersionResponse) SetBody(v *GetProductVersionResponseBody) *GetProductVersionResponse {
	s.Body = v
	return s
}

func (s *GetProductVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
