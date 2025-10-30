// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAXBCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindAXBCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindAXBCallResponse
	GetStatusCode() *int32
	SetBody(v *BindAXBCallResponseBody) *BindAXBCallResponse
	GetBody() *BindAXBCallResponseBody
}

type BindAXBCallResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAXBCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAXBCallResponse) String() string {
	return dara.Prettify(s)
}

func (s BindAXBCallResponse) GoString() string {
	return s.String()
}

func (s *BindAXBCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindAXBCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindAXBCallResponse) GetBody() *BindAXBCallResponseBody {
	return s.Body
}

func (s *BindAXBCallResponse) SetHeaders(v map[string]*string) *BindAXBCallResponse {
	s.Headers = v
	return s
}

func (s *BindAXBCallResponse) SetStatusCode(v int32) *BindAXBCallResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAXBCallResponse) SetBody(v *BindAXBCallResponseBody) *BindAXBCallResponse {
	s.Body = v
	return s
}

func (s *BindAXBCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
