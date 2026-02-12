// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsInstanceCreateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsInstanceCreateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsInstanceCreateResponse
	GetStatusCode() *int32
	SetBody(v *OnsInstanceCreateResponseBody) *OnsInstanceCreateResponse
	GetBody() *OnsInstanceCreateResponseBody
}

type OnsInstanceCreateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsInstanceCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsInstanceCreateResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceCreateResponse) GoString() string {
	return s.String()
}

func (s *OnsInstanceCreateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsInstanceCreateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsInstanceCreateResponse) GetBody() *OnsInstanceCreateResponseBody {
	return s.Body
}

func (s *OnsInstanceCreateResponse) SetHeaders(v map[string]*string) *OnsInstanceCreateResponse {
	s.Headers = v
	return s
}

func (s *OnsInstanceCreateResponse) SetStatusCode(v int32) *OnsInstanceCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsInstanceCreateResponse) SetBody(v *OnsInstanceCreateResponseBody) *OnsInstanceCreateResponse {
	s.Body = v
	return s
}

func (s *OnsInstanceCreateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
