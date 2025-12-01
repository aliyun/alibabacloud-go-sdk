// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsoleScoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConsoleScoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConsoleScoreResponse
	GetStatusCode() *int32
	SetBody(v *GetConsoleScoreResponseBody) *GetConsoleScoreResponse
	GetBody() *GetConsoleScoreResponseBody
}

type GetConsoleScoreResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConsoleScoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConsoleScoreResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConsoleScoreResponse) GoString() string {
	return s.String()
}

func (s *GetConsoleScoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConsoleScoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConsoleScoreResponse) GetBody() *GetConsoleScoreResponseBody {
	return s.Body
}

func (s *GetConsoleScoreResponse) SetHeaders(v map[string]*string) *GetConsoleScoreResponse {
	s.Headers = v
	return s
}

func (s *GetConsoleScoreResponse) SetStatusCode(v int32) *GetConsoleScoreResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsoleScoreResponse) SetBody(v *GetConsoleScoreResponseBody) *GetConsoleScoreResponse {
	s.Body = v
	return s
}

func (s *GetConsoleScoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
