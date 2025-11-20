// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetConferenceHostsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetConferenceHostsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetConferenceHostsResponse
	GetStatusCode() *int32
	SetBody(v *SetConferenceHostsResponseBody) *SetConferenceHostsResponse
	GetBody() *SetConferenceHostsResponseBody
}

type SetConferenceHostsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetConferenceHostsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetConferenceHostsResponse) String() string {
	return dara.Prettify(s)
}

func (s SetConferenceHostsResponse) GoString() string {
	return s.String()
}

func (s *SetConferenceHostsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetConferenceHostsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetConferenceHostsResponse) GetBody() *SetConferenceHostsResponseBody {
	return s.Body
}

func (s *SetConferenceHostsResponse) SetHeaders(v map[string]*string) *SetConferenceHostsResponse {
	s.Headers = v
	return s
}

func (s *SetConferenceHostsResponse) SetStatusCode(v int32) *SetConferenceHostsResponse {
	s.StatusCode = &v
	return s
}

func (s *SetConferenceHostsResponse) SetBody(v *SetConferenceHostsResponseBody) *SetConferenceHostsResponse {
	s.Body = v
	return s
}

func (s *SetConferenceHostsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
