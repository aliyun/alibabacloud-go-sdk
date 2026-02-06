// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBeebotIntentUserSayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBeebotIntentUserSayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBeebotIntentUserSayResponse
	GetStatusCode() *int32
	SetBody(v *ListBeebotIntentUserSayResponseBody) *ListBeebotIntentUserSayResponse
	GetBody() *ListBeebotIntentUserSayResponseBody
}

type ListBeebotIntentUserSayResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBeebotIntentUserSayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBeebotIntentUserSayResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBeebotIntentUserSayResponse) GoString() string {
	return s.String()
}

func (s *ListBeebotIntentUserSayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBeebotIntentUserSayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBeebotIntentUserSayResponse) GetBody() *ListBeebotIntentUserSayResponseBody {
	return s.Body
}

func (s *ListBeebotIntentUserSayResponse) SetHeaders(v map[string]*string) *ListBeebotIntentUserSayResponse {
	s.Headers = v
	return s
}

func (s *ListBeebotIntentUserSayResponse) SetStatusCode(v int32) *ListBeebotIntentUserSayResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBeebotIntentUserSayResponse) SetBody(v *ListBeebotIntentUserSayResponseBody) *ListBeebotIntentUserSayResponse {
	s.Body = v
	return s
}

func (s *ListBeebotIntentUserSayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
