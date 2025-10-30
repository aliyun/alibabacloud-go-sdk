// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnBindAXBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnBindAXBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnBindAXBResponse
	GetStatusCode() *int32
	SetBody(v *UnBindAXBResponseBody) *UnBindAXBResponse
	GetBody() *UnBindAXBResponseBody
}

type UnBindAXBResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnBindAXBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnBindAXBResponse) String() string {
	return dara.Prettify(s)
}

func (s UnBindAXBResponse) GoString() string {
	return s.String()
}

func (s *UnBindAXBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnBindAXBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnBindAXBResponse) GetBody() *UnBindAXBResponseBody {
	return s.Body
}

func (s *UnBindAXBResponse) SetHeaders(v map[string]*string) *UnBindAXBResponse {
	s.Headers = v
	return s
}

func (s *UnBindAXBResponse) SetStatusCode(v int32) *UnBindAXBResponse {
	s.StatusCode = &v
	return s
}

func (s *UnBindAXBResponse) SetBody(v *UnBindAXBResponseBody) *UnBindAXBResponse {
	s.Body = v
	return s
}

func (s *UnBindAXBResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
