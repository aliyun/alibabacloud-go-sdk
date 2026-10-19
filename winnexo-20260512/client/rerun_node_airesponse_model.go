// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerunNodeAIResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RerunNodeAIResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RerunNodeAIResponse
	GetStatusCode() *int32
	SetBody(v *RerunNodeAIResponseBody) *RerunNodeAIResponse
	GetBody() *RerunNodeAIResponseBody
}

type RerunNodeAIResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RerunNodeAIResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RerunNodeAIResponse) String() string {
	return dara.Prettify(s)
}

func (s RerunNodeAIResponse) GoString() string {
	return s.String()
}

func (s *RerunNodeAIResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RerunNodeAIResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RerunNodeAIResponse) GetBody() *RerunNodeAIResponseBody {
	return s.Body
}

func (s *RerunNodeAIResponse) SetHeaders(v map[string]*string) *RerunNodeAIResponse {
	s.Headers = v
	return s
}

func (s *RerunNodeAIResponse) SetStatusCode(v int32) *RerunNodeAIResponse {
	s.StatusCode = &v
	return s
}

func (s *RerunNodeAIResponse) SetBody(v *RerunNodeAIResponseBody) *RerunNodeAIResponse {
	s.Body = v
	return s
}

func (s *RerunNodeAIResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
