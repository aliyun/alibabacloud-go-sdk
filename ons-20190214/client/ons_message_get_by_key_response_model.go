// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsMessageGetByKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsMessageGetByKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsMessageGetByKeyResponse
	GetStatusCode() *int32
	SetBody(v *OnsMessageGetByKeyResponseBody) *OnsMessageGetByKeyResponse
	GetBody() *OnsMessageGetByKeyResponseBody
}

type OnsMessageGetByKeyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsMessageGetByKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsMessageGetByKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageGetByKeyResponse) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsMessageGetByKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsMessageGetByKeyResponse) GetBody() *OnsMessageGetByKeyResponseBody {
	return s.Body
}

func (s *OnsMessageGetByKeyResponse) SetHeaders(v map[string]*string) *OnsMessageGetByKeyResponse {
	s.Headers = v
	return s
}

func (s *OnsMessageGetByKeyResponse) SetStatusCode(v int32) *OnsMessageGetByKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsMessageGetByKeyResponse) SetBody(v *OnsMessageGetByKeyResponseBody) *OnsMessageGetByKeyResponse {
	s.Body = v
	return s
}

func (s *OnsMessageGetByKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
