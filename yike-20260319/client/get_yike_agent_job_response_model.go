// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeAgentJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetYikeAgentJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetYikeAgentJobResponse
	GetStatusCode() *int32
	SetBody(v *GetYikeAgentJobResponseBody) *GetYikeAgentJobResponse
	GetBody() *GetYikeAgentJobResponseBody
}

type GetYikeAgentJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetYikeAgentJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetYikeAgentJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetYikeAgentJobResponse) GoString() string {
	return s.String()
}

func (s *GetYikeAgentJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetYikeAgentJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetYikeAgentJobResponse) GetBody() *GetYikeAgentJobResponseBody {
	return s.Body
}

func (s *GetYikeAgentJobResponse) SetHeaders(v map[string]*string) *GetYikeAgentJobResponse {
	s.Headers = v
	return s
}

func (s *GetYikeAgentJobResponse) SetStatusCode(v int32) *GetYikeAgentJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetYikeAgentJobResponse) SetBody(v *GetYikeAgentJobResponseBody) *GetYikeAgentJobResponse {
	s.Body = v
	return s
}

func (s *GetYikeAgentJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
