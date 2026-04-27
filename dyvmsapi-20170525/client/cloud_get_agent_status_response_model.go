// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetAgentStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudGetAgentStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudGetAgentStatusResponse
	GetStatusCode() *int32
	SetBody(v *CloudGetAgentStatusResponseBody) *CloudGetAgentStatusResponse
	GetBody() *CloudGetAgentStatusResponseBody
}

type CloudGetAgentStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudGetAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudGetAgentStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAgentStatusResponse) GoString() string {
	return s.String()
}

func (s *CloudGetAgentStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudGetAgentStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudGetAgentStatusResponse) GetBody() *CloudGetAgentStatusResponseBody {
	return s.Body
}

func (s *CloudGetAgentStatusResponse) SetHeaders(v map[string]*string) *CloudGetAgentStatusResponse {
	s.Headers = v
	return s
}

func (s *CloudGetAgentStatusResponse) SetStatusCode(v int32) *CloudGetAgentStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudGetAgentStatusResponse) SetBody(v *CloudGetAgentStatusResponseBody) *CloudGetAgentStatusResponse {
	s.Body = v
	return s
}

func (s *CloudGetAgentStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
