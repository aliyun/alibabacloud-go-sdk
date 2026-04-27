// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudAgentUnlinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudAgentUnlinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudAgentUnlinkResponse
	GetStatusCode() *int32
	SetBody(v *CloudAgentUnlinkResponseBody) *CloudAgentUnlinkResponse
	GetBody() *CloudAgentUnlinkResponseBody
}

type CloudAgentUnlinkResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudAgentUnlinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudAgentUnlinkResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentUnlinkResponse) GoString() string {
	return s.String()
}

func (s *CloudAgentUnlinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudAgentUnlinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudAgentUnlinkResponse) GetBody() *CloudAgentUnlinkResponseBody {
	return s.Body
}

func (s *CloudAgentUnlinkResponse) SetHeaders(v map[string]*string) *CloudAgentUnlinkResponse {
	s.Headers = v
	return s
}

func (s *CloudAgentUnlinkResponse) SetStatusCode(v int32) *CloudAgentUnlinkResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudAgentUnlinkResponse) SetBody(v *CloudAgentUnlinkResponseBody) *CloudAgentUnlinkResponse {
	s.Body = v
	return s
}

func (s *CloudAgentUnlinkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
