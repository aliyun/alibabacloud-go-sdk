// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMailAddressByParamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMailAddressByParamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMailAddressByParamResponse
	GetStatusCode() *int32
	SetBody(v *QueryMailAddressByParamResponseBody) *QueryMailAddressByParamResponse
	GetBody() *QueryMailAddressByParamResponseBody
}

type QueryMailAddressByParamResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMailAddressByParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMailAddressByParamResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMailAddressByParamResponse) GoString() string {
	return s.String()
}

func (s *QueryMailAddressByParamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMailAddressByParamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMailAddressByParamResponse) GetBody() *QueryMailAddressByParamResponseBody {
	return s.Body
}

func (s *QueryMailAddressByParamResponse) SetHeaders(v map[string]*string) *QueryMailAddressByParamResponse {
	s.Headers = v
	return s
}

func (s *QueryMailAddressByParamResponse) SetStatusCode(v int32) *QueryMailAddressByParamResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMailAddressByParamResponse) SetBody(v *QueryMailAddressByParamResponseBody) *QueryMailAddressByParamResponse {
	s.Body = v
	return s
}

func (s *QueryMailAddressByParamResponse) Validate() error {
	return dara.Validate(s)
}
