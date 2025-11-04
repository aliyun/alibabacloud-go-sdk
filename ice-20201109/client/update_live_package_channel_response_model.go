// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePackageChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLivePackageChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLivePackageChannelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLivePackageChannelResponseBody) *UpdateLivePackageChannelResponse
	GetBody() *UpdateLivePackageChannelResponseBody
}

type UpdateLivePackageChannelResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLivePackageChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLivePackageChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePackageChannelResponse) GoString() string {
	return s.String()
}

func (s *UpdateLivePackageChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLivePackageChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLivePackageChannelResponse) GetBody() *UpdateLivePackageChannelResponseBody {
	return s.Body
}

func (s *UpdateLivePackageChannelResponse) SetHeaders(v map[string]*string) *UpdateLivePackageChannelResponse {
	s.Headers = v
	return s
}

func (s *UpdateLivePackageChannelResponse) SetStatusCode(v int32) *UpdateLivePackageChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLivePackageChannelResponse) SetBody(v *UpdateLivePackageChannelResponseBody) *UpdateLivePackageChannelResponse {
	s.Body = v
	return s
}

func (s *UpdateLivePackageChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
