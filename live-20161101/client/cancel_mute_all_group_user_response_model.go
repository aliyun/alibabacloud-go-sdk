// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelMuteAllGroupUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelMuteAllGroupUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelMuteAllGroupUserResponse
	GetStatusCode() *int32
	SetBody(v *CancelMuteAllGroupUserResponseBody) *CancelMuteAllGroupUserResponse
	GetBody() *CancelMuteAllGroupUserResponseBody
}

type CancelMuteAllGroupUserResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelMuteAllGroupUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelMuteAllGroupUserResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelMuteAllGroupUserResponse) GoString() string {
	return s.String()
}

func (s *CancelMuteAllGroupUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelMuteAllGroupUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelMuteAllGroupUserResponse) GetBody() *CancelMuteAllGroupUserResponseBody {
	return s.Body
}

func (s *CancelMuteAllGroupUserResponse) SetHeaders(v map[string]*string) *CancelMuteAllGroupUserResponse {
	s.Headers = v
	return s
}

func (s *CancelMuteAllGroupUserResponse) SetStatusCode(v int32) *CancelMuteAllGroupUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelMuteAllGroupUserResponse) SetBody(v *CancelMuteAllGroupUserResponseBody) *CancelMuteAllGroupUserResponse {
	s.Body = v
	return s
}

func (s *CancelMuteAllGroupUserResponse) Validate() error {
	return dara.Validate(s)
}
