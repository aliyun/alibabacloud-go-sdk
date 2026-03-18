// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSlsOpenStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSlsOpenStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSlsOpenStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetSlsOpenStatusResponseBody) *GetSlsOpenStatusResponse
	GetBody() *GetSlsOpenStatusResponseBody
}

type GetSlsOpenStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSlsOpenStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSlsOpenStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSlsOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *GetSlsOpenStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSlsOpenStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSlsOpenStatusResponse) GetBody() *GetSlsOpenStatusResponseBody {
	return s.Body
}

func (s *GetSlsOpenStatusResponse) SetHeaders(v map[string]*string) *GetSlsOpenStatusResponse {
	s.Headers = v
	return s
}

func (s *GetSlsOpenStatusResponse) SetStatusCode(v int32) *GetSlsOpenStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSlsOpenStatusResponse) SetBody(v *GetSlsOpenStatusResponseBody) *GetSlsOpenStatusResponse {
	s.Body = v
	return s
}

func (s *GetSlsOpenStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
