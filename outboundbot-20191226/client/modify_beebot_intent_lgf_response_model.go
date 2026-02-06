// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBeebotIntentLgfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBeebotIntentLgfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBeebotIntentLgfResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBeebotIntentLgfResponseBody) *ModifyBeebotIntentLgfResponse
	GetBody() *ModifyBeebotIntentLgfResponseBody
}

type ModifyBeebotIntentLgfResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBeebotIntentLgfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBeebotIntentLgfResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBeebotIntentLgfResponse) GoString() string {
	return s.String()
}

func (s *ModifyBeebotIntentLgfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBeebotIntentLgfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBeebotIntentLgfResponse) GetBody() *ModifyBeebotIntentLgfResponseBody {
	return s.Body
}

func (s *ModifyBeebotIntentLgfResponse) SetHeaders(v map[string]*string) *ModifyBeebotIntentLgfResponse {
	s.Headers = v
	return s
}

func (s *ModifyBeebotIntentLgfResponse) SetStatusCode(v int32) *ModifyBeebotIntentLgfResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBeebotIntentLgfResponse) SetBody(v *ModifyBeebotIntentLgfResponseBody) *ModifyBeebotIntentLgfResponse {
	s.Body = v
	return s
}

func (s *ModifyBeebotIntentLgfResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
