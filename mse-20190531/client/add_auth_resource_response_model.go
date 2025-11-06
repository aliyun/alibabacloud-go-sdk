// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAuthResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAuthResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAuthResourceResponse
	GetStatusCode() *int32
	SetBody(v *AddAuthResourceResponseBody) *AddAuthResourceResponse
	GetBody() *AddAuthResourceResponseBody
}

type AddAuthResourceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAuthResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAuthResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAuthResourceResponse) GoString() string {
	return s.String()
}

func (s *AddAuthResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAuthResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAuthResourceResponse) GetBody() *AddAuthResourceResponseBody {
	return s.Body
}

func (s *AddAuthResourceResponse) SetHeaders(v map[string]*string) *AddAuthResourceResponse {
	s.Headers = v
	return s
}

func (s *AddAuthResourceResponse) SetStatusCode(v int32) *AddAuthResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAuthResourceResponse) SetBody(v *AddAuthResourceResponseBody) *AddAuthResourceResponse {
	s.Body = v
	return s
}

func (s *AddAuthResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
