// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelMuteGroupUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelMuteGroupUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelMuteGroupUserResponse
	GetStatusCode() *int32
	SetBody(v *CancelMuteGroupUserResponseBody) *CancelMuteGroupUserResponse
	GetBody() *CancelMuteGroupUserResponseBody
}

type CancelMuteGroupUserResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelMuteGroupUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelMuteGroupUserResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelMuteGroupUserResponse) GoString() string {
	return s.String()
}

func (s *CancelMuteGroupUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelMuteGroupUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelMuteGroupUserResponse) GetBody() *CancelMuteGroupUserResponseBody {
	return s.Body
}

func (s *CancelMuteGroupUserResponse) SetHeaders(v map[string]*string) *CancelMuteGroupUserResponse {
	s.Headers = v
	return s
}

func (s *CancelMuteGroupUserResponse) SetStatusCode(v int32) *CancelMuteGroupUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelMuteGroupUserResponse) SetBody(v *CancelMuteGroupUserResponseBody) *CancelMuteGroupUserResponse {
	s.Body = v
	return s
}

func (s *CancelMuteGroupUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
