// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBeebotIntentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBeebotIntentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBeebotIntentResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBeebotIntentResponseBody) *ModifyBeebotIntentResponse
	GetBody() *ModifyBeebotIntentResponseBody
}

type ModifyBeebotIntentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBeebotIntentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBeebotIntentResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBeebotIntentResponse) GoString() string {
	return s.String()
}

func (s *ModifyBeebotIntentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBeebotIntentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBeebotIntentResponse) GetBody() *ModifyBeebotIntentResponseBody {
	return s.Body
}

func (s *ModifyBeebotIntentResponse) SetHeaders(v map[string]*string) *ModifyBeebotIntentResponse {
	s.Headers = v
	return s
}

func (s *ModifyBeebotIntentResponse) SetStatusCode(v int32) *ModifyBeebotIntentResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBeebotIntentResponse) SetBody(v *ModifyBeebotIntentResponseBody) *ModifyBeebotIntentResponse {
	s.Body = v
	return s
}

func (s *ModifyBeebotIntentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
