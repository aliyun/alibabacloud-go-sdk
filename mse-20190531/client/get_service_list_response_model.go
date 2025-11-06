// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceListResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceListResponseBody) *GetServiceListResponse
	GetBody() *GetServiceListResponseBody
}

type GetServiceListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceListResponse) GoString() string {
	return s.String()
}

func (s *GetServiceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceListResponse) GetBody() *GetServiceListResponseBody {
	return s.Body
}

func (s *GetServiceListResponse) SetHeaders(v map[string]*string) *GetServiceListResponse {
	s.Headers = v
	return s
}

func (s *GetServiceListResponse) SetStatusCode(v int32) *GetServiceListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceListResponse) SetBody(v *GetServiceListResponseBody) *GetServiceListResponse {
	s.Body = v
	return s
}

func (s *GetServiceListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
