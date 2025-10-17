// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartDBLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartDBLinkResponse
	GetStatusCode() *int32
	SetBody(v *RestartDBLinkResponseBody) *RestartDBLinkResponse
	GetBody() *RestartDBLinkResponseBody
}

type RestartDBLinkResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartDBLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartDBLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartDBLinkResponse) GoString() string {
	return s.String()
}

func (s *RestartDBLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartDBLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartDBLinkResponse) GetBody() *RestartDBLinkResponseBody {
	return s.Body
}

func (s *RestartDBLinkResponse) SetHeaders(v map[string]*string) *RestartDBLinkResponse {
	s.Headers = v
	return s
}

func (s *RestartDBLinkResponse) SetStatusCode(v int32) *RestartDBLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartDBLinkResponse) SetBody(v *RestartDBLinkResponseBody) *RestartDBLinkResponse {
	s.Body = v
	return s
}

func (s *RestartDBLinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
