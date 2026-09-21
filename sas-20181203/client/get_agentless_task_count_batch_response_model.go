// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentlessTaskCountBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentlessTaskCountBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentlessTaskCountBatchResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentlessTaskCountBatchResponseBody) *GetAgentlessTaskCountBatchResponse
	GetBody() *GetAgentlessTaskCountBatchResponseBody
}

type GetAgentlessTaskCountBatchResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentlessTaskCountBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentlessTaskCountBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentlessTaskCountBatchResponse) GoString() string {
	return s.String()
}

func (s *GetAgentlessTaskCountBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentlessTaskCountBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentlessTaskCountBatchResponse) GetBody() *GetAgentlessTaskCountBatchResponseBody {
	return s.Body
}

func (s *GetAgentlessTaskCountBatchResponse) SetHeaders(v map[string]*string) *GetAgentlessTaskCountBatchResponse {
	s.Headers = v
	return s
}

func (s *GetAgentlessTaskCountBatchResponse) SetStatusCode(v int32) *GetAgentlessTaskCountBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentlessTaskCountBatchResponse) SetBody(v *GetAgentlessTaskCountBatchResponseBody) *GetAgentlessTaskCountBatchResponse {
	s.Body = v
	return s
}

func (s *GetAgentlessTaskCountBatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
