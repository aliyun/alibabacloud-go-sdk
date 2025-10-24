// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectIpWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProjectIpWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProjectIpWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProjectIpWhiteListResponseBody) *UpdateProjectIpWhiteListResponse
	GetBody() *UpdateProjectIpWhiteListResponseBody
}

type UpdateProjectIpWhiteListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProjectIpWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectIpWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectIpWhiteListResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectIpWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProjectIpWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProjectIpWhiteListResponse) GetBody() *UpdateProjectIpWhiteListResponseBody {
	return s.Body
}

func (s *UpdateProjectIpWhiteListResponse) SetHeaders(v map[string]*string) *UpdateProjectIpWhiteListResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectIpWhiteListResponse) SetStatusCode(v int32) *UpdateProjectIpWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectIpWhiteListResponse) SetBody(v *UpdateProjectIpWhiteListResponseBody) *UpdateProjectIpWhiteListResponse {
	s.Body = v
	return s
}

func (s *UpdateProjectIpWhiteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
