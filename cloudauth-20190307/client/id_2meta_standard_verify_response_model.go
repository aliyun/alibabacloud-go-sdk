// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId2MetaStandardVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Id2MetaStandardVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Id2MetaStandardVerifyResponse
	GetStatusCode() *int32
	SetBody(v *Id2MetaStandardVerifyResponseBody) *Id2MetaStandardVerifyResponse
	GetBody() *Id2MetaStandardVerifyResponseBody
}

type Id2MetaStandardVerifyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Id2MetaStandardVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Id2MetaStandardVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaStandardVerifyResponse) GoString() string {
	return s.String()
}

func (s *Id2MetaStandardVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Id2MetaStandardVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Id2MetaStandardVerifyResponse) GetBody() *Id2MetaStandardVerifyResponseBody {
	return s.Body
}

func (s *Id2MetaStandardVerifyResponse) SetHeaders(v map[string]*string) *Id2MetaStandardVerifyResponse {
	s.Headers = v
	return s
}

func (s *Id2MetaStandardVerifyResponse) SetStatusCode(v int32) *Id2MetaStandardVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *Id2MetaStandardVerifyResponse) SetBody(v *Id2MetaStandardVerifyResponseBody) *Id2MetaStandardVerifyResponse {
	s.Body = v
	return s
}

func (s *Id2MetaStandardVerifyResponse) Validate() error {
	return dara.Validate(s)
}
