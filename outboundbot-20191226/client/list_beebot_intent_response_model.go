// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBeebotIntentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBeebotIntentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBeebotIntentResponse
	GetStatusCode() *int32
	SetBody(v *ListBeebotIntentResponseBody) *ListBeebotIntentResponse
	GetBody() *ListBeebotIntentResponseBody
}

type ListBeebotIntentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBeebotIntentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBeebotIntentResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBeebotIntentResponse) GoString() string {
	return s.String()
}

func (s *ListBeebotIntentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBeebotIntentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBeebotIntentResponse) GetBody() *ListBeebotIntentResponseBody {
	return s.Body
}

func (s *ListBeebotIntentResponse) SetHeaders(v map[string]*string) *ListBeebotIntentResponse {
	s.Headers = v
	return s
}

func (s *ListBeebotIntentResponse) SetStatusCode(v int32) *ListBeebotIntentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBeebotIntentResponse) SetBody(v *ListBeebotIntentResponseBody) *ListBeebotIntentResponse {
	s.Body = v
	return s
}

func (s *ListBeebotIntentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
