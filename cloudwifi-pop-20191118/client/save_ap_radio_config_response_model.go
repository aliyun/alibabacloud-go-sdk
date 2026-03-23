// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveApRadioConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveApRadioConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveApRadioConfigResponse
	GetStatusCode() *int32
	SetBody(v *SaveApRadioConfigResponseBody) *SaveApRadioConfigResponse
	GetBody() *SaveApRadioConfigResponseBody
}

type SaveApRadioConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveApRadioConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveApRadioConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveApRadioConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveApRadioConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveApRadioConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveApRadioConfigResponse) GetBody() *SaveApRadioConfigResponseBody {
	return s.Body
}

func (s *SaveApRadioConfigResponse) SetHeaders(v map[string]*string) *SaveApRadioConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveApRadioConfigResponse) SetStatusCode(v int32) *SaveApRadioConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveApRadioConfigResponse) SetBody(v *SaveApRadioConfigResponseBody) *SaveApRadioConfigResponse {
	s.Body = v
	return s
}

func (s *SaveApRadioConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
