// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBeebotIntentUserSayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBeebotIntentUserSayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBeebotIntentUserSayResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBeebotIntentUserSayResponseBody) *DeleteBeebotIntentUserSayResponse
	GetBody() *DeleteBeebotIntentUserSayResponseBody
}

type DeleteBeebotIntentUserSayResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBeebotIntentUserSayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBeebotIntentUserSayResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBeebotIntentUserSayResponse) GoString() string {
	return s.String()
}

func (s *DeleteBeebotIntentUserSayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBeebotIntentUserSayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBeebotIntentUserSayResponse) GetBody() *DeleteBeebotIntentUserSayResponseBody {
	return s.Body
}

func (s *DeleteBeebotIntentUserSayResponse) SetHeaders(v map[string]*string) *DeleteBeebotIntentUserSayResponse {
	s.Headers = v
	return s
}

func (s *DeleteBeebotIntentUserSayResponse) SetStatusCode(v int32) *DeleteBeebotIntentUserSayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBeebotIntentUserSayResponse) SetBody(v *DeleteBeebotIntentUserSayResponseBody) *DeleteBeebotIntentUserSayResponse {
	s.Body = v
	return s
}

func (s *DeleteBeebotIntentUserSayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
