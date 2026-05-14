// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkAgentStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkAgentStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkAgentStatusResponse
	GetStatusCode() *int32
	SetBody(v *ClinkAgentStatusResponseBody) *ClinkAgentStatusResponse
	GetBody() *ClinkAgentStatusResponseBody
}

type ClinkAgentStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkAgentStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkAgentStatusResponse) GoString() string {
	return s.String()
}

func (s *ClinkAgentStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkAgentStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkAgentStatusResponse) GetBody() *ClinkAgentStatusResponseBody {
	return s.Body
}

func (s *ClinkAgentStatusResponse) SetHeaders(v map[string]*string) *ClinkAgentStatusResponse {
	s.Headers = v
	return s
}

func (s *ClinkAgentStatusResponse) SetStatusCode(v int32) *ClinkAgentStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkAgentStatusResponse) SetBody(v *ClinkAgentStatusResponseBody) *ClinkAgentStatusResponse {
	s.Body = v
	return s
}

func (s *ClinkAgentStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
