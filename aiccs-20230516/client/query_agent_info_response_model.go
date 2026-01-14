// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAgentInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAgentInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAgentInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryAgentInfoResponseBody) *QueryAgentInfoResponse
	GetBody() *QueryAgentInfoResponseBody
}

type QueryAgentInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAgentInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAgentInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAgentInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryAgentInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAgentInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAgentInfoResponse) GetBody() *QueryAgentInfoResponseBody {
	return s.Body
}

func (s *QueryAgentInfoResponse) SetHeaders(v map[string]*string) *QueryAgentInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryAgentInfoResponse) SetStatusCode(v int32) *QueryAgentInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAgentInfoResponse) SetBody(v *QueryAgentInfoResponseBody) *QueryAgentInfoResponse {
	s.Body = v
	return s
}

func (s *QueryAgentInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
