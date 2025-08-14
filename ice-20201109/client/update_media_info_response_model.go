// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMediaInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMediaInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMediaInfoResponseBody) *UpdateMediaInfoResponse
	GetBody() *UpdateMediaInfoResponseBody
}

type UpdateMediaInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMediaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMediaInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateMediaInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMediaInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMediaInfoResponse) GetBody() *UpdateMediaInfoResponseBody {
	return s.Body
}

func (s *UpdateMediaInfoResponse) SetHeaders(v map[string]*string) *UpdateMediaInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateMediaInfoResponse) SetStatusCode(v int32) *UpdateMediaInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMediaInfoResponse) SetBody(v *UpdateMediaInfoResponseBody) *UpdateMediaInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateMediaInfoResponse) Validate() error {
	return dara.Validate(s)
}
