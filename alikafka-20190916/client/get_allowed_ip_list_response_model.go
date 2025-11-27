// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllowedIpListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAllowedIpListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAllowedIpListResponse
	GetStatusCode() *int32
	SetBody(v *GetAllowedIpListResponseBody) *GetAllowedIpListResponse
	GetBody() *GetAllowedIpListResponseBody
}

type GetAllowedIpListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAllowedIpListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAllowedIpListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAllowedIpListResponse) GoString() string {
	return s.String()
}

func (s *GetAllowedIpListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAllowedIpListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAllowedIpListResponse) GetBody() *GetAllowedIpListResponseBody {
	return s.Body
}

func (s *GetAllowedIpListResponse) SetHeaders(v map[string]*string) *GetAllowedIpListResponse {
	s.Headers = v
	return s
}

func (s *GetAllowedIpListResponse) SetStatusCode(v int32) *GetAllowedIpListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAllowedIpListResponse) SetBody(v *GetAllowedIpListResponseBody) *GetAllowedIpListResponse {
	s.Body = v
	return s
}

func (s *GetAllowedIpListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
