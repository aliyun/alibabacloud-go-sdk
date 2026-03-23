// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApBasicConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveApBasicConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveApBasicConfigResponse
	GetStatusCode() *int32
	SetBody(v *SaveApBasicConfigResponseBody) *SaveApBasicConfigResponse
	GetBody() *SaveApBasicConfigResponseBody
}

type SaveApBasicConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveApBasicConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveApBasicConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveApBasicConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveApBasicConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveApBasicConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveApBasicConfigResponse) GetBody() *SaveApBasicConfigResponseBody {
	return s.Body
}

func (s *SaveApBasicConfigResponse) SetHeaders(v map[string]*string) *SaveApBasicConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveApBasicConfigResponse) SetStatusCode(v int32) *SaveApBasicConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveApBasicConfigResponse) SetBody(v *SaveApBasicConfigResponseBody) *SaveApBasicConfigResponse {
	s.Body = v
	return s
}

func (s *SaveApBasicConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
