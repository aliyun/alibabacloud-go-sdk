// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivePackageChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLivePackageChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLivePackageChannelResponse
	GetStatusCode() *int32
	SetBody(v *CreateLivePackageChannelResponseBody) *CreateLivePackageChannelResponse
	GetBody() *CreateLivePackageChannelResponseBody
}

type CreateLivePackageChannelResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLivePackageChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLivePackageChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePackageChannelResponse) GoString() string {
	return s.String()
}

func (s *CreateLivePackageChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLivePackageChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLivePackageChannelResponse) GetBody() *CreateLivePackageChannelResponseBody {
	return s.Body
}

func (s *CreateLivePackageChannelResponse) SetHeaders(v map[string]*string) *CreateLivePackageChannelResponse {
	s.Headers = v
	return s
}

func (s *CreateLivePackageChannelResponse) SetStatusCode(v int32) *CreateLivePackageChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLivePackageChannelResponse) SetBody(v *CreateLivePackageChannelResponseBody) *CreateLivePackageChannelResponse {
	s.Body = v
	return s
}

func (s *CreateLivePackageChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
