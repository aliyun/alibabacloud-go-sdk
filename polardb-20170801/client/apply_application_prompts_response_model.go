// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyApplicationPromptsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyApplicationPromptsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyApplicationPromptsResponse
	GetStatusCode() *int32
	SetBody(v *ApplyApplicationPromptsResponseBody) *ApplyApplicationPromptsResponse
	GetBody() *ApplyApplicationPromptsResponseBody
}

type ApplyApplicationPromptsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyApplicationPromptsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyApplicationPromptsResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyApplicationPromptsResponse) GoString() string {
	return s.String()
}

func (s *ApplyApplicationPromptsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyApplicationPromptsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyApplicationPromptsResponse) GetBody() *ApplyApplicationPromptsResponseBody {
	return s.Body
}

func (s *ApplyApplicationPromptsResponse) SetHeaders(v map[string]*string) *ApplyApplicationPromptsResponse {
	s.Headers = v
	return s
}

func (s *ApplyApplicationPromptsResponse) SetStatusCode(v int32) *ApplyApplicationPromptsResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyApplicationPromptsResponse) SetBody(v *ApplyApplicationPromptsResponseBody) *ApplyApplicationPromptsResponse {
	s.Body = v
	return s
}

func (s *ApplyApplicationPromptsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
