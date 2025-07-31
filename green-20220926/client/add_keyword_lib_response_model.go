// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKeywordLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddKeywordLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddKeywordLibResponse
	GetStatusCode() *int32
	SetBody(v *AddKeywordLibResponseBody) *AddKeywordLibResponse
	GetBody() *AddKeywordLibResponseBody
}

type AddKeywordLibResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddKeywordLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddKeywordLibResponse) String() string {
	return dara.Prettify(s)
}

func (s AddKeywordLibResponse) GoString() string {
	return s.String()
}

func (s *AddKeywordLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddKeywordLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddKeywordLibResponse) GetBody() *AddKeywordLibResponseBody {
	return s.Body
}

func (s *AddKeywordLibResponse) SetHeaders(v map[string]*string) *AddKeywordLibResponse {
	s.Headers = v
	return s
}

func (s *AddKeywordLibResponse) SetStatusCode(v int32) *AddKeywordLibResponse {
	s.StatusCode = &v
	return s
}

func (s *AddKeywordLibResponse) SetBody(v *AddKeywordLibResponseBody) *AddKeywordLibResponse {
	s.Body = v
	return s
}

func (s *AddKeywordLibResponse) Validate() error {
	return dara.Validate(s)
}
