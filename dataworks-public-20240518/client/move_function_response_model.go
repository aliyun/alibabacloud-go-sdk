// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MoveFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MoveFunctionResponse
	GetStatusCode() *int32
	SetBody(v *MoveFunctionResponseBody) *MoveFunctionResponse
	GetBody() *MoveFunctionResponseBody
}

type MoveFunctionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s MoveFunctionResponse) GoString() string {
	return s.String()
}

func (s *MoveFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MoveFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MoveFunctionResponse) GetBody() *MoveFunctionResponseBody {
	return s.Body
}

func (s *MoveFunctionResponse) SetHeaders(v map[string]*string) *MoveFunctionResponse {
	s.Headers = v
	return s
}

func (s *MoveFunctionResponse) SetStatusCode(v int32) *MoveFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveFunctionResponse) SetBody(v *MoveFunctionResponseBody) *MoveFunctionResponse {
	s.Body = v
	return s
}

func (s *MoveFunctionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
