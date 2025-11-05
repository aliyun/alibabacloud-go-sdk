// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFCTriggerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFCTriggerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFCTriggerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFCTriggerResponseBody) *DeleteFCTriggerResponse
	GetBody() *DeleteFCTriggerResponseBody
}

type DeleteFCTriggerResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFCTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFCTriggerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFCTriggerResponse) GoString() string {
	return s.String()
}

func (s *DeleteFCTriggerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFCTriggerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFCTriggerResponse) GetBody() *DeleteFCTriggerResponseBody {
	return s.Body
}

func (s *DeleteFCTriggerResponse) SetHeaders(v map[string]*string) *DeleteFCTriggerResponse {
	s.Headers = v
	return s
}

func (s *DeleteFCTriggerResponse) SetStatusCode(v int32) *DeleteFCTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFCTriggerResponse) SetBody(v *DeleteFCTriggerResponseBody) *DeleteFCTriggerResponse {
	s.Body = v
	return s
}

func (s *DeleteFCTriggerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
