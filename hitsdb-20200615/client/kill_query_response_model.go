// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *KillQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *KillQueryResponse
	GetStatusCode() *int32
	SetBody(v *KillQueryResponseBody) *KillQueryResponse
	GetBody() *KillQueryResponseBody
}

type KillQueryResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KillQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KillQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s KillQueryResponse) GoString() string {
	return s.String()
}

func (s *KillQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *KillQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *KillQueryResponse) GetBody() *KillQueryResponseBody {
	return s.Body
}

func (s *KillQueryResponse) SetHeaders(v map[string]*string) *KillQueryResponse {
	s.Headers = v
	return s
}

func (s *KillQueryResponse) SetStatusCode(v int32) *KillQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *KillQueryResponse) SetBody(v *KillQueryResponseBody) *KillQueryResponse {
	s.Body = v
	return s
}

func (s *KillQueryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
