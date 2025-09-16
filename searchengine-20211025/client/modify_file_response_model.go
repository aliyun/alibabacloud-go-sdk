// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyFileResponse
	GetStatusCode() *int32
	SetBody(v *ModifyFileResponseBody) *ModifyFileResponse
	GetBody() *ModifyFileResponseBody
}

type ModifyFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFileResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyFileResponse) GoString() string {
	return s.String()
}

func (s *ModifyFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyFileResponse) GetBody() *ModifyFileResponseBody {
	return s.Body
}

func (s *ModifyFileResponse) SetHeaders(v map[string]*string) *ModifyFileResponse {
	s.Headers = v
	return s
}

func (s *ModifyFileResponse) SetStatusCode(v int32) *ModifyFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFileResponse) SetBody(v *ModifyFileResponseBody) *ModifyFileResponse {
	s.Body = v
	return s
}

func (s *ModifyFileResponse) Validate() error {
	return dara.Validate(s)
}
