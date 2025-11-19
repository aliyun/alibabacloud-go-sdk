// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAIJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAIJobResponse
	GetStatusCode() *int32
	SetBody(v *ListAIJobResponseBody) *ListAIJobResponse
	GetBody() *ListAIJobResponseBody
}

type ListAIJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAIJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAIJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAIJobResponse) GoString() string {
	return s.String()
}

func (s *ListAIJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAIJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAIJobResponse) GetBody() *ListAIJobResponseBody {
	return s.Body
}

func (s *ListAIJobResponse) SetHeaders(v map[string]*string) *ListAIJobResponse {
	s.Headers = v
	return s
}

func (s *ListAIJobResponse) SetStatusCode(v int32) *ListAIJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAIJobResponse) SetBody(v *ListAIJobResponseBody) *ListAIJobResponse {
	s.Body = v
	return s
}

func (s *ListAIJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
