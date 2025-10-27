// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageVulWhitelistTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateImageVulWhitelistTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateImageVulWhitelistTargetResponse
	GetStatusCode() *int32
	SetBody(v *UpdateImageVulWhitelistTargetResponseBody) *UpdateImageVulWhitelistTargetResponse
	GetBody() *UpdateImageVulWhitelistTargetResponseBody
}

type UpdateImageVulWhitelistTargetResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateImageVulWhitelistTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateImageVulWhitelistTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageVulWhitelistTargetResponse) GoString() string {
	return s.String()
}

func (s *UpdateImageVulWhitelistTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateImageVulWhitelistTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateImageVulWhitelistTargetResponse) GetBody() *UpdateImageVulWhitelistTargetResponseBody {
	return s.Body
}

func (s *UpdateImageVulWhitelistTargetResponse) SetHeaders(v map[string]*string) *UpdateImageVulWhitelistTargetResponse {
	s.Headers = v
	return s
}

func (s *UpdateImageVulWhitelistTargetResponse) SetStatusCode(v int32) *UpdateImageVulWhitelistTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateImageVulWhitelistTargetResponse) SetBody(v *UpdateImageVulWhitelistTargetResponseBody) *UpdateImageVulWhitelistTargetResponse {
	s.Body = v
	return s
}

func (s *UpdateImageVulWhitelistTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
