// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisconnectAndroidInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisconnectAndroidInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisconnectAndroidInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DisconnectAndroidInstanceResponseBody) *DisconnectAndroidInstanceResponse
	GetBody() *DisconnectAndroidInstanceResponseBody
}

type DisconnectAndroidInstanceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisconnectAndroidInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisconnectAndroidInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DisconnectAndroidInstanceResponse) GoString() string {
	return s.String()
}

func (s *DisconnectAndroidInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisconnectAndroidInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisconnectAndroidInstanceResponse) GetBody() *DisconnectAndroidInstanceResponseBody {
	return s.Body
}

func (s *DisconnectAndroidInstanceResponse) SetHeaders(v map[string]*string) *DisconnectAndroidInstanceResponse {
	s.Headers = v
	return s
}

func (s *DisconnectAndroidInstanceResponse) SetStatusCode(v int32) *DisconnectAndroidInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DisconnectAndroidInstanceResponse) SetBody(v *DisconnectAndroidInstanceResponseBody) *DisconnectAndroidInstanceResponse {
	s.Body = v
	return s
}

func (s *DisconnectAndroidInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
