// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetListenerAccessControlStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetListenerAccessControlStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetListenerAccessControlStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetListenerAccessControlStatusResponseBody) *SetListenerAccessControlStatusResponse
	GetBody() *SetListenerAccessControlStatusResponseBody
}

type SetListenerAccessControlStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetListenerAccessControlStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetListenerAccessControlStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetListenerAccessControlStatusResponse) GoString() string {
	return s.String()
}

func (s *SetListenerAccessControlStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetListenerAccessControlStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetListenerAccessControlStatusResponse) GetBody() *SetListenerAccessControlStatusResponseBody {
	return s.Body
}

func (s *SetListenerAccessControlStatusResponse) SetHeaders(v map[string]*string) *SetListenerAccessControlStatusResponse {
	s.Headers = v
	return s
}

func (s *SetListenerAccessControlStatusResponse) SetStatusCode(v int32) *SetListenerAccessControlStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetListenerAccessControlStatusResponse) SetBody(v *SetListenerAccessControlStatusResponseBody) *SetListenerAccessControlStatusResponse {
	s.Body = v
	return s
}

func (s *SetListenerAccessControlStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
