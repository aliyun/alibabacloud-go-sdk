// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConsoleClearPretendStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConsoleClearPretendStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConsoleClearPretendStatusResponse
	GetStatusCode() *int32
	SetBody(v *ConsoleClearPretendStatusResponseBody) *ConsoleClearPretendStatusResponse
	GetBody() *ConsoleClearPretendStatusResponseBody
}

type ConsoleClearPretendStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConsoleClearPretendStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConsoleClearPretendStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ConsoleClearPretendStatusResponse) GoString() string {
	return s.String()
}

func (s *ConsoleClearPretendStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConsoleClearPretendStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConsoleClearPretendStatusResponse) GetBody() *ConsoleClearPretendStatusResponseBody {
	return s.Body
}

func (s *ConsoleClearPretendStatusResponse) SetHeaders(v map[string]*string) *ConsoleClearPretendStatusResponse {
	s.Headers = v
	return s
}

func (s *ConsoleClearPretendStatusResponse) SetStatusCode(v int32) *ConsoleClearPretendStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ConsoleClearPretendStatusResponse) SetBody(v *ConsoleClearPretendStatusResponseBody) *ConsoleClearPretendStatusResponse {
	s.Body = v
	return s
}

func (s *ConsoleClearPretendStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
