// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceGroupSourceFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReplaceGroupSourceFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReplaceGroupSourceFileResponse
	GetStatusCode() *int32
	SetBody(v *ReplaceGroupSourceFileResponseBody) *ReplaceGroupSourceFileResponse
	GetBody() *ReplaceGroupSourceFileResponseBody
}

type ReplaceGroupSourceFileResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplaceGroupSourceFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplaceGroupSourceFileResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplaceGroupSourceFileResponse) GoString() string {
	return s.String()
}

func (s *ReplaceGroupSourceFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReplaceGroupSourceFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReplaceGroupSourceFileResponse) GetBody() *ReplaceGroupSourceFileResponseBody {
	return s.Body
}

func (s *ReplaceGroupSourceFileResponse) SetHeaders(v map[string]*string) *ReplaceGroupSourceFileResponse {
	s.Headers = v
	return s
}

func (s *ReplaceGroupSourceFileResponse) SetStatusCode(v int32) *ReplaceGroupSourceFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceGroupSourceFileResponse) SetBody(v *ReplaceGroupSourceFileResponseBody) *ReplaceGroupSourceFileResponse {
	s.Body = v
	return s
}

func (s *ReplaceGroupSourceFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
