// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSpecLatestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentSpecLatestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentSpecLatestResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentSpecLatestResponseBody) *GetAgentSpecLatestResponse
	GetBody() *GetAgentSpecLatestResponseBody
}

type GetAgentSpecLatestResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentSpecLatestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentSpecLatestResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpecLatestResponse) GoString() string {
	return s.String()
}

func (s *GetAgentSpecLatestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentSpecLatestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentSpecLatestResponse) GetBody() *GetAgentSpecLatestResponseBody {
	return s.Body
}

func (s *GetAgentSpecLatestResponse) SetHeaders(v map[string]*string) *GetAgentSpecLatestResponse {
	s.Headers = v
	return s
}

func (s *GetAgentSpecLatestResponse) SetStatusCode(v int32) *GetAgentSpecLatestResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentSpecLatestResponse) SetBody(v *GetAgentSpecLatestResponseBody) *GetAgentSpecLatestResponse {
	s.Body = v
	return s
}

func (s *GetAgentSpecLatestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
