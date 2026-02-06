// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBeebotIntentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBeebotIntentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBeebotIntentResponse
	GetStatusCode() *int32
	SetBody(v *CreateBeebotIntentResponseBody) *CreateBeebotIntentResponse
	GetBody() *CreateBeebotIntentResponseBody
}

type CreateBeebotIntentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBeebotIntentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBeebotIntentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBeebotIntentResponse) GoString() string {
	return s.String()
}

func (s *CreateBeebotIntentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBeebotIntentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBeebotIntentResponse) GetBody() *CreateBeebotIntentResponseBody {
	return s.Body
}

func (s *CreateBeebotIntentResponse) SetHeaders(v map[string]*string) *CreateBeebotIntentResponse {
	s.Headers = v
	return s
}

func (s *CreateBeebotIntentResponse) SetStatusCode(v int32) *CreateBeebotIntentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBeebotIntentResponse) SetBody(v *CreateBeebotIntentResponseBody) *CreateBeebotIntentResponse {
	s.Body = v
	return s
}

func (s *CreateBeebotIntentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
