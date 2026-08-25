// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceStatusResponseBody) *GetServiceStatusResponse
	GetBody() *GetServiceStatusResponseBody
}

type GetServiceStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceStatusResponse) GetBody() *GetServiceStatusResponseBody {
	return s.Body
}

func (s *GetServiceStatusResponse) SetHeaders(v map[string]*string) *GetServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetServiceStatusResponse) SetStatusCode(v int32) *GetServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceStatusResponse) SetBody(v *GetServiceStatusResponseBody) *GetServiceStatusResponse {
	s.Body = v
	return s
}

func (s *GetServiceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
