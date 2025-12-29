// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateScreenSaverResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddOrUpdateScreenSaverResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddOrUpdateScreenSaverResponse
	GetStatusCode() *int32
	SetBody(v *AddOrUpdateScreenSaverResponseBody) *AddOrUpdateScreenSaverResponse
	GetBody() *AddOrUpdateScreenSaverResponseBody
}

type AddOrUpdateScreenSaverResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddOrUpdateScreenSaverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddOrUpdateScreenSaverResponse) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateScreenSaverResponse) GoString() string {
	return s.String()
}

func (s *AddOrUpdateScreenSaverResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddOrUpdateScreenSaverResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddOrUpdateScreenSaverResponse) GetBody() *AddOrUpdateScreenSaverResponseBody {
	return s.Body
}

func (s *AddOrUpdateScreenSaverResponse) SetHeaders(v map[string]*string) *AddOrUpdateScreenSaverResponse {
	s.Headers = v
	return s
}

func (s *AddOrUpdateScreenSaverResponse) SetStatusCode(v int32) *AddOrUpdateScreenSaverResponse {
	s.StatusCode = &v
	return s
}

func (s *AddOrUpdateScreenSaverResponse) SetBody(v *AddOrUpdateScreenSaverResponseBody) *AddOrUpdateScreenSaverResponse {
	s.Body = v
	return s
}

func (s *AddOrUpdateScreenSaverResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
