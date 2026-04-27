// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteAgentTelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudDeleteAgentTelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudDeleteAgentTelResponse
	GetStatusCode() *int32
	SetBody(v *CloudDeleteAgentTelResponseBody) *CloudDeleteAgentTelResponse
	GetBody() *CloudDeleteAgentTelResponseBody
}

type CloudDeleteAgentTelResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudDeleteAgentTelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudDeleteAgentTelResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteAgentTelResponse) GoString() string {
	return s.String()
}

func (s *CloudDeleteAgentTelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudDeleteAgentTelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudDeleteAgentTelResponse) GetBody() *CloudDeleteAgentTelResponseBody {
	return s.Body
}

func (s *CloudDeleteAgentTelResponse) SetHeaders(v map[string]*string) *CloudDeleteAgentTelResponse {
	s.Headers = v
	return s
}

func (s *CloudDeleteAgentTelResponse) SetStatusCode(v int32) *CloudDeleteAgentTelResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudDeleteAgentTelResponse) SetBody(v *CloudDeleteAgentTelResponseBody) *CloudDeleteAgentTelResponse {
	s.Body = v
	return s
}

func (s *CloudDeleteAgentTelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
