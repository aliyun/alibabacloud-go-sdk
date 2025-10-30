// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnBindXBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnBindXBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnBindXBResponse
	GetStatusCode() *int32
	SetBody(v *UnBindXBResponseBody) *UnBindXBResponse
	GetBody() *UnBindXBResponseBody
}

type UnBindXBResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnBindXBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnBindXBResponse) String() string {
	return dara.Prettify(s)
}

func (s UnBindXBResponse) GoString() string {
	return s.String()
}

func (s *UnBindXBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnBindXBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnBindXBResponse) GetBody() *UnBindXBResponseBody {
	return s.Body
}

func (s *UnBindXBResponse) SetHeaders(v map[string]*string) *UnBindXBResponse {
	s.Headers = v
	return s
}

func (s *UnBindXBResponse) SetStatusCode(v int32) *UnBindXBResponse {
	s.StatusCode = &v
	return s
}

func (s *UnBindXBResponse) SetBody(v *UnBindXBResponseBody) *UnBindXBResponse {
	s.Body = v
	return s
}

func (s *UnBindXBResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
