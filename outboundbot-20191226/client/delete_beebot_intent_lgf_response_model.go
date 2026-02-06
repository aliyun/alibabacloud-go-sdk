// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBeebotIntentLgfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBeebotIntentLgfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBeebotIntentLgfResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBeebotIntentLgfResponseBody) *DeleteBeebotIntentLgfResponse
	GetBody() *DeleteBeebotIntentLgfResponseBody
}

type DeleteBeebotIntentLgfResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBeebotIntentLgfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBeebotIntentLgfResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBeebotIntentLgfResponse) GoString() string {
	return s.String()
}

func (s *DeleteBeebotIntentLgfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBeebotIntentLgfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBeebotIntentLgfResponse) GetBody() *DeleteBeebotIntentLgfResponseBody {
	return s.Body
}

func (s *DeleteBeebotIntentLgfResponse) SetHeaders(v map[string]*string) *DeleteBeebotIntentLgfResponse {
	s.Headers = v
	return s
}

func (s *DeleteBeebotIntentLgfResponse) SetStatusCode(v int32) *DeleteBeebotIntentLgfResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBeebotIntentLgfResponse) SetBody(v *DeleteBeebotIntentLgfResponseBody) *DeleteBeebotIntentLgfResponse {
	s.Body = v
	return s
}

func (s *DeleteBeebotIntentLgfResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
