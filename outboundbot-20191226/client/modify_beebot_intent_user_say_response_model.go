// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBeebotIntentUserSayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBeebotIntentUserSayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBeebotIntentUserSayResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBeebotIntentUserSayResponseBody) *ModifyBeebotIntentUserSayResponse
	GetBody() *ModifyBeebotIntentUserSayResponseBody
}

type ModifyBeebotIntentUserSayResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBeebotIntentUserSayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBeebotIntentUserSayResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBeebotIntentUserSayResponse) GoString() string {
	return s.String()
}

func (s *ModifyBeebotIntentUserSayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBeebotIntentUserSayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBeebotIntentUserSayResponse) GetBody() *ModifyBeebotIntentUserSayResponseBody {
	return s.Body
}

func (s *ModifyBeebotIntentUserSayResponse) SetHeaders(v map[string]*string) *ModifyBeebotIntentUserSayResponse {
	s.Headers = v
	return s
}

func (s *ModifyBeebotIntentUserSayResponse) SetStatusCode(v int32) *ModifyBeebotIntentUserSayResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBeebotIntentUserSayResponse) SetBody(v *ModifyBeebotIntentUserSayResponseBody) *ModifyBeebotIntentUserSayResponse {
	s.Body = v
	return s
}

func (s *ModifyBeebotIntentUserSayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
