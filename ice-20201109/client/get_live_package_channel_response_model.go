// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLivePackageChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLivePackageChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLivePackageChannelResponse
	GetStatusCode() *int32
	SetBody(v *GetLivePackageChannelResponseBody) *GetLivePackageChannelResponse
	GetBody() *GetLivePackageChannelResponseBody
}

type GetLivePackageChannelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLivePackageChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLivePackageChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLivePackageChannelResponse) GoString() string {
	return s.String()
}

func (s *GetLivePackageChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLivePackageChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLivePackageChannelResponse) GetBody() *GetLivePackageChannelResponseBody {
	return s.Body
}

func (s *GetLivePackageChannelResponse) SetHeaders(v map[string]*string) *GetLivePackageChannelResponse {
	s.Headers = v
	return s
}

func (s *GetLivePackageChannelResponse) SetStatusCode(v int32) *GetLivePackageChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLivePackageChannelResponse) SetBody(v *GetLivePackageChannelResponseBody) *GetLivePackageChannelResponse {
	s.Body = v
	return s
}

func (s *GetLivePackageChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
