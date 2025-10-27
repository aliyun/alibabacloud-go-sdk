// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceIpWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceIpWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceIpWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceIpWhiteListResponseBody) *GetInstanceIpWhiteListResponse
	GetBody() *GetInstanceIpWhiteListResponseBody
}

type GetInstanceIpWhiteListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceIpWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceIpWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceIpWhiteListResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceIpWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceIpWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceIpWhiteListResponse) GetBody() *GetInstanceIpWhiteListResponseBody {
	return s.Body
}

func (s *GetInstanceIpWhiteListResponse) SetHeaders(v map[string]*string) *GetInstanceIpWhiteListResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceIpWhiteListResponse) SetStatusCode(v int32) *GetInstanceIpWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceIpWhiteListResponse) SetBody(v *GetInstanceIpWhiteListResponseBody) *GetInstanceIpWhiteListResponse {
	s.Body = v
	return s
}

func (s *GetInstanceIpWhiteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
