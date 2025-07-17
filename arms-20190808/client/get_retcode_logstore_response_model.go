// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRetcodeLogstoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRetcodeLogstoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRetcodeLogstoreResponse
	GetStatusCode() *int32
	SetBody(v *GetRetcodeLogstoreResponseBody) *GetRetcodeLogstoreResponse
	GetBody() *GetRetcodeLogstoreResponseBody
}

type GetRetcodeLogstoreResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRetcodeLogstoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRetcodeLogstoreResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRetcodeLogstoreResponse) GoString() string {
	return s.String()
}

func (s *GetRetcodeLogstoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRetcodeLogstoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRetcodeLogstoreResponse) GetBody() *GetRetcodeLogstoreResponseBody {
	return s.Body
}

func (s *GetRetcodeLogstoreResponse) SetHeaders(v map[string]*string) *GetRetcodeLogstoreResponse {
	s.Headers = v
	return s
}

func (s *GetRetcodeLogstoreResponse) SetStatusCode(v int32) *GetRetcodeLogstoreResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRetcodeLogstoreResponse) SetBody(v *GetRetcodeLogstoreResponseBody) *GetRetcodeLogstoreResponse {
	s.Body = v
	return s
}

func (s *GetRetcodeLogstoreResponse) Validate() error {
	return dara.Validate(s)
}
