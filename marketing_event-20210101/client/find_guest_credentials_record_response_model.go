// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindGuestCredentialsRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FindGuestCredentialsRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FindGuestCredentialsRecordResponse
	GetStatusCode() *int32
	SetBody(v *FindGuestCredentialsRecordResponseBody) *FindGuestCredentialsRecordResponse
	GetBody() *FindGuestCredentialsRecordResponseBody
}

type FindGuestCredentialsRecordResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FindGuestCredentialsRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FindGuestCredentialsRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s FindGuestCredentialsRecordResponse) GoString() string {
	return s.String()
}

func (s *FindGuestCredentialsRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FindGuestCredentialsRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FindGuestCredentialsRecordResponse) GetBody() *FindGuestCredentialsRecordResponseBody {
	return s.Body
}

func (s *FindGuestCredentialsRecordResponse) SetHeaders(v map[string]*string) *FindGuestCredentialsRecordResponse {
	s.Headers = v
	return s
}

func (s *FindGuestCredentialsRecordResponse) SetStatusCode(v int32) *FindGuestCredentialsRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *FindGuestCredentialsRecordResponse) SetBody(v *FindGuestCredentialsRecordResponseBody) *FindGuestCredentialsRecordResponse {
	s.Body = v
	return s
}

func (s *FindGuestCredentialsRecordResponse) Validate() error {
	return dara.Validate(s)
}
