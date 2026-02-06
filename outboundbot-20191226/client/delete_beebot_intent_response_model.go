// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBeebotIntentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBeebotIntentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBeebotIntentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBeebotIntentResponseBody) *DeleteBeebotIntentResponse
	GetBody() *DeleteBeebotIntentResponseBody
}

type DeleteBeebotIntentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBeebotIntentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBeebotIntentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBeebotIntentResponse) GoString() string {
	return s.String()
}

func (s *DeleteBeebotIntentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBeebotIntentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBeebotIntentResponse) GetBody() *DeleteBeebotIntentResponseBody {
	return s.Body
}

func (s *DeleteBeebotIntentResponse) SetHeaders(v map[string]*string) *DeleteBeebotIntentResponse {
	s.Headers = v
	return s
}

func (s *DeleteBeebotIntentResponse) SetStatusCode(v int32) *DeleteBeebotIntentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBeebotIntentResponse) SetBody(v *DeleteBeebotIntentResponseBody) *DeleteBeebotIntentResponse {
	s.Body = v
	return s
}

func (s *DeleteBeebotIntentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
