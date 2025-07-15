// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverPhysicalConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecoverPhysicalConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecoverPhysicalConnectionResponse
	GetStatusCode() *int32
	SetBody(v *RecoverPhysicalConnectionResponseBody) *RecoverPhysicalConnectionResponse
	GetBody() *RecoverPhysicalConnectionResponseBody
}

type RecoverPhysicalConnectionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecoverPhysicalConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecoverPhysicalConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s RecoverPhysicalConnectionResponse) GoString() string {
	return s.String()
}

func (s *RecoverPhysicalConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecoverPhysicalConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecoverPhysicalConnectionResponse) GetBody() *RecoverPhysicalConnectionResponseBody {
	return s.Body
}

func (s *RecoverPhysicalConnectionResponse) SetHeaders(v map[string]*string) *RecoverPhysicalConnectionResponse {
	s.Headers = v
	return s
}

func (s *RecoverPhysicalConnectionResponse) SetStatusCode(v int32) *RecoverPhysicalConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoverPhysicalConnectionResponse) SetBody(v *RecoverPhysicalConnectionResponseBody) *RecoverPhysicalConnectionResponse {
	s.Body = v
	return s
}

func (s *RecoverPhysicalConnectionResponse) Validate() error {
	return dara.Validate(s)
}
