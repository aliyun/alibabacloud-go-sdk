// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBeebotIntentUserSayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBeebotIntentUserSayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBeebotIntentUserSayResponse
	GetStatusCode() *int32
	SetBody(v *CreateBeebotIntentUserSayResponseBody) *CreateBeebotIntentUserSayResponse
	GetBody() *CreateBeebotIntentUserSayResponseBody
}

type CreateBeebotIntentUserSayResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBeebotIntentUserSayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBeebotIntentUserSayResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBeebotIntentUserSayResponse) GoString() string {
	return s.String()
}

func (s *CreateBeebotIntentUserSayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBeebotIntentUserSayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBeebotIntentUserSayResponse) GetBody() *CreateBeebotIntentUserSayResponseBody {
	return s.Body
}

func (s *CreateBeebotIntentUserSayResponse) SetHeaders(v map[string]*string) *CreateBeebotIntentUserSayResponse {
	s.Headers = v
	return s
}

func (s *CreateBeebotIntentUserSayResponse) SetStatusCode(v int32) *CreateBeebotIntentUserSayResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBeebotIntentUserSayResponse) SetBody(v *CreateBeebotIntentUserSayResponseBody) *CreateBeebotIntentUserSayResponse {
	s.Body = v
	return s
}

func (s *CreateBeebotIntentUserSayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
