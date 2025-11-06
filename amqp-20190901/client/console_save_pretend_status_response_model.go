// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConsoleSavePretendStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConsoleSavePretendStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConsoleSavePretendStatusResponse
	GetStatusCode() *int32
	SetBody(v *ConsoleSavePretendStatusResponseBody) *ConsoleSavePretendStatusResponse
	GetBody() *ConsoleSavePretendStatusResponseBody
}

type ConsoleSavePretendStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConsoleSavePretendStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConsoleSavePretendStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ConsoleSavePretendStatusResponse) GoString() string {
	return s.String()
}

func (s *ConsoleSavePretendStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConsoleSavePretendStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConsoleSavePretendStatusResponse) GetBody() *ConsoleSavePretendStatusResponseBody {
	return s.Body
}

func (s *ConsoleSavePretendStatusResponse) SetHeaders(v map[string]*string) *ConsoleSavePretendStatusResponse {
	s.Headers = v
	return s
}

func (s *ConsoleSavePretendStatusResponse) SetStatusCode(v int32) *ConsoleSavePretendStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ConsoleSavePretendStatusResponse) SetBody(v *ConsoleSavePretendStatusResponseBody) *ConsoleSavePretendStatusResponse {
	s.Body = v
	return s
}

func (s *ConsoleSavePretendStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
