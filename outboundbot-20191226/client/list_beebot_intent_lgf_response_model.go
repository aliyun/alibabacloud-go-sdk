// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBeebotIntentLgfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBeebotIntentLgfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBeebotIntentLgfResponse
	GetStatusCode() *int32
	SetBody(v *ListBeebotIntentLgfResponseBody) *ListBeebotIntentLgfResponse
	GetBody() *ListBeebotIntentLgfResponseBody
}

type ListBeebotIntentLgfResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBeebotIntentLgfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBeebotIntentLgfResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBeebotIntentLgfResponse) GoString() string {
	return s.String()
}

func (s *ListBeebotIntentLgfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBeebotIntentLgfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBeebotIntentLgfResponse) GetBody() *ListBeebotIntentLgfResponseBody {
	return s.Body
}

func (s *ListBeebotIntentLgfResponse) SetHeaders(v map[string]*string) *ListBeebotIntentLgfResponse {
	s.Headers = v
	return s
}

func (s *ListBeebotIntentLgfResponse) SetStatusCode(v int32) *ListBeebotIntentLgfResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBeebotIntentLgfResponse) SetBody(v *ListBeebotIntentLgfResponseBody) *ListBeebotIntentLgfResponse {
	s.Body = v
	return s
}

func (s *ListBeebotIntentLgfResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
