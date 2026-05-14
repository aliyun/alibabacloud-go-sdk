// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListCdrIbAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkListCdrIbAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkListCdrIbAgentResponse
	GetStatusCode() *int32
	SetBody(v *ClinkListCdrIbAgentResponseBody) *ClinkListCdrIbAgentResponse
	GetBody() *ClinkListCdrIbAgentResponseBody
}

type ClinkListCdrIbAgentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkListCdrIbAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkListCdrIbAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrIbAgentResponse) GoString() string {
	return s.String()
}

func (s *ClinkListCdrIbAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkListCdrIbAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkListCdrIbAgentResponse) GetBody() *ClinkListCdrIbAgentResponseBody {
	return s.Body
}

func (s *ClinkListCdrIbAgentResponse) SetHeaders(v map[string]*string) *ClinkListCdrIbAgentResponse {
	s.Headers = v
	return s
}

func (s *ClinkListCdrIbAgentResponse) SetStatusCode(v int32) *ClinkListCdrIbAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkListCdrIbAgentResponse) SetBody(v *ClinkListCdrIbAgentResponseBody) *ClinkListCdrIbAgentResponse {
	s.Body = v
	return s
}

func (s *ClinkListCdrIbAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
