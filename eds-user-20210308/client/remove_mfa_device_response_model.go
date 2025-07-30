// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveMfaDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveMfaDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveMfaDeviceResponse
	GetStatusCode() *int32
	SetBody(v *RemoveMfaDeviceResponseBody) *RemoveMfaDeviceResponse
	GetBody() *RemoveMfaDeviceResponseBody
}

type RemoveMfaDeviceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveMfaDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveMfaDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveMfaDeviceResponse) GoString() string {
	return s.String()
}

func (s *RemoveMfaDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveMfaDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveMfaDeviceResponse) GetBody() *RemoveMfaDeviceResponseBody {
	return s.Body
}

func (s *RemoveMfaDeviceResponse) SetHeaders(v map[string]*string) *RemoveMfaDeviceResponse {
	s.Headers = v
	return s
}

func (s *RemoveMfaDeviceResponse) SetStatusCode(v int32) *RemoveMfaDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveMfaDeviceResponse) SetBody(v *RemoveMfaDeviceResponseBody) *RemoveMfaDeviceResponse {
	s.Body = v
	return s
}

func (s *RemoveMfaDeviceResponse) Validate() error {
	return dara.Validate(s)
}
