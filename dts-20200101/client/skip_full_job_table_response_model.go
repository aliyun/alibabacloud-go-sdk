// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkipFullJobTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SkipFullJobTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SkipFullJobTableResponse
	GetStatusCode() *int32
	SetBody(v *SkipFullJobTableResponseBody) *SkipFullJobTableResponse
	GetBody() *SkipFullJobTableResponseBody
}

type SkipFullJobTableResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SkipFullJobTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SkipFullJobTableResponse) String() string {
	return dara.Prettify(s)
}

func (s SkipFullJobTableResponse) GoString() string {
	return s.String()
}

func (s *SkipFullJobTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SkipFullJobTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SkipFullJobTableResponse) GetBody() *SkipFullJobTableResponseBody {
	return s.Body
}

func (s *SkipFullJobTableResponse) SetHeaders(v map[string]*string) *SkipFullJobTableResponse {
	s.Headers = v
	return s
}

func (s *SkipFullJobTableResponse) SetStatusCode(v int32) *SkipFullJobTableResponse {
	s.StatusCode = &v
	return s
}

func (s *SkipFullJobTableResponse) SetBody(v *SkipFullJobTableResponseBody) *SkipFullJobTableResponse {
	s.Body = v
	return s
}

func (s *SkipFullJobTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
