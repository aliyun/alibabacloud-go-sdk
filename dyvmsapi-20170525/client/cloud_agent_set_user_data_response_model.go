// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudAgentSetUserDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudAgentSetUserDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudAgentSetUserDataResponse
	GetStatusCode() *int32
	SetBody(v *CloudAgentSetUserDataResponseBody) *CloudAgentSetUserDataResponse
	GetBody() *CloudAgentSetUserDataResponseBody
}

type CloudAgentSetUserDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudAgentSetUserDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudAgentSetUserDataResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentSetUserDataResponse) GoString() string {
	return s.String()
}

func (s *CloudAgentSetUserDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudAgentSetUserDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudAgentSetUserDataResponse) GetBody() *CloudAgentSetUserDataResponseBody {
	return s.Body
}

func (s *CloudAgentSetUserDataResponse) SetHeaders(v map[string]*string) *CloudAgentSetUserDataResponse {
	s.Headers = v
	return s
}

func (s *CloudAgentSetUserDataResponse) SetStatusCode(v int32) *CloudAgentSetUserDataResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudAgentSetUserDataResponse) SetBody(v *CloudAgentSetUserDataResponseBody) *CloudAgentSetUserDataResponse {
	s.Body = v
	return s
}

func (s *CloudAgentSetUserDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
