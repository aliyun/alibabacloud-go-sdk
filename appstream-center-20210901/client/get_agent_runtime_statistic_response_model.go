// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentRuntimeStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentRuntimeStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentRuntimeStatisticResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentRuntimeStatisticResponseBody) *GetAgentRuntimeStatisticResponse
	GetBody() *GetAgentRuntimeStatisticResponseBody
}

type GetAgentRuntimeStatisticResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentRuntimeStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentRuntimeStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentRuntimeStatisticResponse) GoString() string {
	return s.String()
}

func (s *GetAgentRuntimeStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentRuntimeStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentRuntimeStatisticResponse) GetBody() *GetAgentRuntimeStatisticResponseBody {
	return s.Body
}

func (s *GetAgentRuntimeStatisticResponse) SetHeaders(v map[string]*string) *GetAgentRuntimeStatisticResponse {
	s.Headers = v
	return s
}

func (s *GetAgentRuntimeStatisticResponse) SetStatusCode(v int32) *GetAgentRuntimeStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentRuntimeStatisticResponse) SetBody(v *GetAgentRuntimeStatisticResponseBody) *GetAgentRuntimeStatisticResponse {
	s.Body = v
	return s
}

func (s *GetAgentRuntimeStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
