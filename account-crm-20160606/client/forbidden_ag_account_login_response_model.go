// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForbiddenAgAccountLoginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ForbiddenAgAccountLoginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ForbiddenAgAccountLoginResponse
	GetStatusCode() *int32
	SetBody(v *ForbiddenAgAccountLoginResponseBody) *ForbiddenAgAccountLoginResponse
	GetBody() *ForbiddenAgAccountLoginResponseBody
}

type ForbiddenAgAccountLoginResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ForbiddenAgAccountLoginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ForbiddenAgAccountLoginResponse) String() string {
	return dara.Prettify(s)
}

func (s ForbiddenAgAccountLoginResponse) GoString() string {
	return s.String()
}

func (s *ForbiddenAgAccountLoginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ForbiddenAgAccountLoginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ForbiddenAgAccountLoginResponse) GetBody() *ForbiddenAgAccountLoginResponseBody {
	return s.Body
}

func (s *ForbiddenAgAccountLoginResponse) SetHeaders(v map[string]*string) *ForbiddenAgAccountLoginResponse {
	s.Headers = v
	return s
}

func (s *ForbiddenAgAccountLoginResponse) SetStatusCode(v int32) *ForbiddenAgAccountLoginResponse {
	s.StatusCode = &v
	return s
}

func (s *ForbiddenAgAccountLoginResponse) SetBody(v *ForbiddenAgAccountLoginResponseBody) *ForbiddenAgAccountLoginResponse {
	s.Body = v
	return s
}

func (s *ForbiddenAgAccountLoginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
