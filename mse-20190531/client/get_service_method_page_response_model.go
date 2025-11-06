// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceMethodPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceMethodPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceMethodPageResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceMethodPageResponseBody) *GetServiceMethodPageResponse
	GetBody() *GetServiceMethodPageResponseBody
}

type GetServiceMethodPageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceMethodPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceMethodPageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceMethodPageResponse) GoString() string {
	return s.String()
}

func (s *GetServiceMethodPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceMethodPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceMethodPageResponse) GetBody() *GetServiceMethodPageResponseBody {
	return s.Body
}

func (s *GetServiceMethodPageResponse) SetHeaders(v map[string]*string) *GetServiceMethodPageResponse {
	s.Headers = v
	return s
}

func (s *GetServiceMethodPageResponse) SetStatusCode(v int32) *GetServiceMethodPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceMethodPageResponse) SetBody(v *GetServiceMethodPageResponseBody) *GetServiceMethodPageResponse {
	s.Body = v
	return s
}

func (s *GetServiceMethodPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
