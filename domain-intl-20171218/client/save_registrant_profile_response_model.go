// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveRegistrantProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveRegistrantProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveRegistrantProfileResponse
	GetStatusCode() *int32
	SetBody(v *SaveRegistrantProfileResponseBody) *SaveRegistrantProfileResponse
	GetBody() *SaveRegistrantProfileResponseBody
}

type SaveRegistrantProfileResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveRegistrantProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveRegistrantProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveRegistrantProfileResponse) GoString() string {
	return s.String()
}

func (s *SaveRegistrantProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveRegistrantProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveRegistrantProfileResponse) GetBody() *SaveRegistrantProfileResponseBody {
	return s.Body
}

func (s *SaveRegistrantProfileResponse) SetHeaders(v map[string]*string) *SaveRegistrantProfileResponse {
	s.Headers = v
	return s
}

func (s *SaveRegistrantProfileResponse) SetStatusCode(v int32) *SaveRegistrantProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveRegistrantProfileResponse) SetBody(v *SaveRegistrantProfileResponseBody) *SaveRegistrantProfileResponse {
	s.Body = v
	return s
}

func (s *SaveRegistrantProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
