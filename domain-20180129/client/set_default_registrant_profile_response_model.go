// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultRegistrantProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDefaultRegistrantProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDefaultRegistrantProfileResponse
	GetStatusCode() *int32
	SetBody(v *SetDefaultRegistrantProfileResponseBody) *SetDefaultRegistrantProfileResponse
	GetBody() *SetDefaultRegistrantProfileResponseBody
}

type SetDefaultRegistrantProfileResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDefaultRegistrantProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDefaultRegistrantProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultRegistrantProfileResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultRegistrantProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDefaultRegistrantProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDefaultRegistrantProfileResponse) GetBody() *SetDefaultRegistrantProfileResponseBody {
	return s.Body
}

func (s *SetDefaultRegistrantProfileResponse) SetHeaders(v map[string]*string) *SetDefaultRegistrantProfileResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultRegistrantProfileResponse) SetStatusCode(v int32) *SetDefaultRegistrantProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDefaultRegistrantProfileResponse) SetBody(v *SetDefaultRegistrantProfileResponseBody) *SetDefaultRegistrantProfileResponse {
	s.Body = v
	return s
}

func (s *SetDefaultRegistrantProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
