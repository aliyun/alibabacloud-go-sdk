// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHttp2EnableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHttp2EnableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHttp2EnableResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHttp2EnableResponseBody) *ModifyHttp2EnableResponse
	GetBody() *ModifyHttp2EnableResponseBody
}

type ModifyHttp2EnableResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHttp2EnableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHttp2EnableResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHttp2EnableResponse) GoString() string {
	return s.String()
}

func (s *ModifyHttp2EnableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHttp2EnableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHttp2EnableResponse) GetBody() *ModifyHttp2EnableResponseBody {
	return s.Body
}

func (s *ModifyHttp2EnableResponse) SetHeaders(v map[string]*string) *ModifyHttp2EnableResponse {
	s.Headers = v
	return s
}

func (s *ModifyHttp2EnableResponse) SetStatusCode(v int32) *ModifyHttp2EnableResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHttp2EnableResponse) SetBody(v *ModifyHttp2EnableResponseBody) *ModifyHttp2EnableResponse {
	s.Body = v
	return s
}

func (s *ModifyHttp2EnableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
