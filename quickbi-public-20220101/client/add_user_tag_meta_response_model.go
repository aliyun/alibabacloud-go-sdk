// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserTagMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddUserTagMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddUserTagMetaResponse
	GetStatusCode() *int32
	SetBody(v *AddUserTagMetaResponseBody) *AddUserTagMetaResponse
	GetBody() *AddUserTagMetaResponseBody
}

type AddUserTagMetaResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserTagMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserTagMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s AddUserTagMetaResponse) GoString() string {
	return s.String()
}

func (s *AddUserTagMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddUserTagMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddUserTagMetaResponse) GetBody() *AddUserTagMetaResponseBody {
	return s.Body
}

func (s *AddUserTagMetaResponse) SetHeaders(v map[string]*string) *AddUserTagMetaResponse {
	s.Headers = v
	return s
}

func (s *AddUserTagMetaResponse) SetStatusCode(v int32) *AddUserTagMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserTagMetaResponse) SetBody(v *AddUserTagMetaResponseBody) *AddUserTagMetaResponse {
	s.Body = v
	return s
}

func (s *AddUserTagMetaResponse) Validate() error {
	return dara.Validate(s)
}
