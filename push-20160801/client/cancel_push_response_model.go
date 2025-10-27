// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPushResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelPushResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelPushResponse
	GetStatusCode() *int32
	SetBody(v *CancelPushResponseBody) *CancelPushResponse
	GetBody() *CancelPushResponseBody
}

type CancelPushResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelPushResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelPushResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelPushResponse) GoString() string {
	return s.String()
}

func (s *CancelPushResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelPushResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelPushResponse) GetBody() *CancelPushResponseBody {
	return s.Body
}

func (s *CancelPushResponse) SetHeaders(v map[string]*string) *CancelPushResponse {
	s.Headers = v
	return s
}

func (s *CancelPushResponse) SetStatusCode(v int32) *CancelPushResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelPushResponse) SetBody(v *CancelPushResponseBody) *CancelPushResponse {
	s.Body = v
	return s
}

func (s *CancelPushResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
