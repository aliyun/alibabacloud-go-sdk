// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIMBotsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIMBotsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIMBotsResponse
	GetStatusCode() *int32
	SetBody(v *ListIMBotsResult) *ListIMBotsResponse
	GetBody() *ListIMBotsResult
}

type ListIMBotsResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIMBotsResult  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIMBotsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIMBotsResponse) GoString() string {
	return s.String()
}

func (s *ListIMBotsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIMBotsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIMBotsResponse) GetBody() *ListIMBotsResult {
	return s.Body
}

func (s *ListIMBotsResponse) SetHeaders(v map[string]*string) *ListIMBotsResponse {
	s.Headers = v
	return s
}

func (s *ListIMBotsResponse) SetStatusCode(v int32) *ListIMBotsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIMBotsResponse) SetBody(v *ListIMBotsResult) *ListIMBotsResponse {
	s.Body = v
	return s
}

func (s *ListIMBotsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
