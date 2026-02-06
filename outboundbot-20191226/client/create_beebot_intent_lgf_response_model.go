// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBeebotIntentLgfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBeebotIntentLgfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBeebotIntentLgfResponse
	GetStatusCode() *int32
	SetBody(v *CreateBeebotIntentLgfResponseBody) *CreateBeebotIntentLgfResponse
	GetBody() *CreateBeebotIntentLgfResponseBody
}

type CreateBeebotIntentLgfResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBeebotIntentLgfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBeebotIntentLgfResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBeebotIntentLgfResponse) GoString() string {
	return s.String()
}

func (s *CreateBeebotIntentLgfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBeebotIntentLgfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBeebotIntentLgfResponse) GetBody() *CreateBeebotIntentLgfResponseBody {
	return s.Body
}

func (s *CreateBeebotIntentLgfResponse) SetHeaders(v map[string]*string) *CreateBeebotIntentLgfResponse {
	s.Headers = v
	return s
}

func (s *CreateBeebotIntentLgfResponse) SetStatusCode(v int32) *CreateBeebotIntentLgfResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBeebotIntentLgfResponse) SetBody(v *CreateBeebotIntentLgfResponseBody) *CreateBeebotIntentLgfResponse {
	s.Body = v
	return s
}

func (s *CreateBeebotIntentLgfResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
