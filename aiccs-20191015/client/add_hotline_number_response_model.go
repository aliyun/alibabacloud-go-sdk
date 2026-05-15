// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddHotlineNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddHotlineNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddHotlineNumberResponse
	GetStatusCode() *int32
	SetBody(v *AddHotlineNumberResponseBody) *AddHotlineNumberResponse
	GetBody() *AddHotlineNumberResponseBody
}

type AddHotlineNumberResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddHotlineNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddHotlineNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s AddHotlineNumberResponse) GoString() string {
	return s.String()
}

func (s *AddHotlineNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddHotlineNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddHotlineNumberResponse) GetBody() *AddHotlineNumberResponseBody {
	return s.Body
}

func (s *AddHotlineNumberResponse) SetHeaders(v map[string]*string) *AddHotlineNumberResponse {
	s.Headers = v
	return s
}

func (s *AddHotlineNumberResponse) SetStatusCode(v int32) *AddHotlineNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *AddHotlineNumberResponse) SetBody(v *AddHotlineNumberResponseBody) *AddHotlineNumberResponse {
	s.Body = v
	return s
}

func (s *AddHotlineNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
