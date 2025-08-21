// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCnameReuseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCnameReuseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCnameReuseResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCnameReuseResponseBody) *ModifyCnameReuseResponse
	GetBody() *ModifyCnameReuseResponseBody
}

type ModifyCnameReuseResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCnameReuseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCnameReuseResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCnameReuseResponse) GoString() string {
	return s.String()
}

func (s *ModifyCnameReuseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCnameReuseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCnameReuseResponse) GetBody() *ModifyCnameReuseResponseBody {
	return s.Body
}

func (s *ModifyCnameReuseResponse) SetHeaders(v map[string]*string) *ModifyCnameReuseResponse {
	s.Headers = v
	return s
}

func (s *ModifyCnameReuseResponse) SetStatusCode(v int32) *ModifyCnameReuseResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCnameReuseResponse) SetBody(v *ModifyCnameReuseResponseBody) *ModifyCnameReuseResponse {
	s.Body = v
	return s
}

func (s *ModifyCnameReuseResponse) Validate() error {
	return dara.Validate(s)
}
