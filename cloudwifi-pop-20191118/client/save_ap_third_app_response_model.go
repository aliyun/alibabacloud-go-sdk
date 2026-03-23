// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApThirdAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveApThirdAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveApThirdAppResponse
	GetStatusCode() *int32
	SetBody(v *SaveApThirdAppResponseBody) *SaveApThirdAppResponse
	GetBody() *SaveApThirdAppResponseBody
}

type SaveApThirdAppResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveApThirdAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveApThirdAppResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveApThirdAppResponse) GoString() string {
	return s.String()
}

func (s *SaveApThirdAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveApThirdAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveApThirdAppResponse) GetBody() *SaveApThirdAppResponseBody {
	return s.Body
}

func (s *SaveApThirdAppResponse) SetHeaders(v map[string]*string) *SaveApThirdAppResponse {
	s.Headers = v
	return s
}

func (s *SaveApThirdAppResponse) SetStatusCode(v int32) *SaveApThirdAppResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveApThirdAppResponse) SetBody(v *SaveApThirdAppResponseBody) *SaveApThirdAppResponse {
	s.Body = v
	return s
}

func (s *SaveApThirdAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
