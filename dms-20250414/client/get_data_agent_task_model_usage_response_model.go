// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAgentTaskModelUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataAgentTaskModelUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataAgentTaskModelUsageResponse
	GetStatusCode() *int32
	SetBody(v *GetDataAgentTaskModelUsageResponseBody) *GetDataAgentTaskModelUsageResponse
	GetBody() *GetDataAgentTaskModelUsageResponseBody
}

type GetDataAgentTaskModelUsageResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataAgentTaskModelUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataAgentTaskModelUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentTaskModelUsageResponse) GoString() string {
	return s.String()
}

func (s *GetDataAgentTaskModelUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataAgentTaskModelUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataAgentTaskModelUsageResponse) GetBody() *GetDataAgentTaskModelUsageResponseBody {
	return s.Body
}

func (s *GetDataAgentTaskModelUsageResponse) SetHeaders(v map[string]*string) *GetDataAgentTaskModelUsageResponse {
	s.Headers = v
	return s
}

func (s *GetDataAgentTaskModelUsageResponse) SetStatusCode(v int32) *GetDataAgentTaskModelUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataAgentTaskModelUsageResponse) SetBody(v *GetDataAgentTaskModelUsageResponseBody) *GetDataAgentTaskModelUsageResponse {
	s.Body = v
	return s
}

func (s *GetDataAgentTaskModelUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
