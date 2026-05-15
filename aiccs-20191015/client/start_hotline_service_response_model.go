// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartHotlineServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartHotlineServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartHotlineServiceResponse
	GetStatusCode() *int32
	SetBody(v *StartHotlineServiceResponseBody) *StartHotlineServiceResponse
	GetBody() *StartHotlineServiceResponseBody
}

type StartHotlineServiceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartHotlineServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartHotlineServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartHotlineServiceResponse) GoString() string {
	return s.String()
}

func (s *StartHotlineServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartHotlineServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartHotlineServiceResponse) GetBody() *StartHotlineServiceResponseBody {
	return s.Body
}

func (s *StartHotlineServiceResponse) SetHeaders(v map[string]*string) *StartHotlineServiceResponse {
	s.Headers = v
	return s
}

func (s *StartHotlineServiceResponse) SetStatusCode(v int32) *StartHotlineServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartHotlineServiceResponse) SetBody(v *StartHotlineServiceResponseBody) *StartHotlineServiceResponse {
	s.Body = v
	return s
}

func (s *StartHotlineServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
