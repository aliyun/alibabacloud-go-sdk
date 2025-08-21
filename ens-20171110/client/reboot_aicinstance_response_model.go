// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootAICInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebootAICInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebootAICInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RebootAICInstanceResponseBody) *RebootAICInstanceResponse
	GetBody() *RebootAICInstanceResponseBody
}

type RebootAICInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootAICInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootAICInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RebootAICInstanceResponse) GoString() string {
	return s.String()
}

func (s *RebootAICInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebootAICInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebootAICInstanceResponse) GetBody() *RebootAICInstanceResponseBody {
	return s.Body
}

func (s *RebootAICInstanceResponse) SetHeaders(v map[string]*string) *RebootAICInstanceResponse {
	s.Headers = v
	return s
}

func (s *RebootAICInstanceResponse) SetStatusCode(v int32) *RebootAICInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootAICInstanceResponse) SetBody(v *RebootAICInstanceResponseBody) *RebootAICInstanceResponse {
	s.Body = v
	return s
}

func (s *RebootAICInstanceResponse) Validate() error {
	return dara.Validate(s)
}
