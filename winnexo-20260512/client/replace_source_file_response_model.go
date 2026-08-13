// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceSourceFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReplaceSourceFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReplaceSourceFileResponse
	GetStatusCode() *int32
	SetBody(v *ReplaceSourceFileResponseBody) *ReplaceSourceFileResponse
	GetBody() *ReplaceSourceFileResponseBody
}

type ReplaceSourceFileResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplaceSourceFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplaceSourceFileResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplaceSourceFileResponse) GoString() string {
	return s.String()
}

func (s *ReplaceSourceFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReplaceSourceFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReplaceSourceFileResponse) GetBody() *ReplaceSourceFileResponseBody {
	return s.Body
}

func (s *ReplaceSourceFileResponse) SetHeaders(v map[string]*string) *ReplaceSourceFileResponse {
	s.Headers = v
	return s
}

func (s *ReplaceSourceFileResponse) SetStatusCode(v int32) *ReplaceSourceFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceSourceFileResponse) SetBody(v *ReplaceSourceFileResponseBody) *ReplaceSourceFileResponse {
	s.Body = v
	return s
}

func (s *ReplaceSourceFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
