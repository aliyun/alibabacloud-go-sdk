// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSeoStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppSeoStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppSeoStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetAppSeoStatusResponseBody) *GetAppSeoStatusResponse
	GetBody() *GetAppSeoStatusResponseBody
}

type GetAppSeoStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppSeoStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppSeoStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppSeoStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAppSeoStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppSeoStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppSeoStatusResponse) GetBody() *GetAppSeoStatusResponseBody {
	return s.Body
}

func (s *GetAppSeoStatusResponse) SetHeaders(v map[string]*string) *GetAppSeoStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAppSeoStatusResponse) SetStatusCode(v int32) *GetAppSeoStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppSeoStatusResponse) SetBody(v *GetAppSeoStatusResponseBody) *GetAppSeoStatusResponse {
	s.Body = v
	return s
}

func (s *GetAppSeoStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
