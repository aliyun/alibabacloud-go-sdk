// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCuResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCuResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCuResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCuResponseBody) *ModifyCuResponse
	GetBody() *ModifyCuResponseBody
}

type ModifyCuResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCuResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCuResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCuResponse) GoString() string {
	return s.String()
}

func (s *ModifyCuResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCuResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCuResponse) GetBody() *ModifyCuResponseBody {
	return s.Body
}

func (s *ModifyCuResponse) SetHeaders(v map[string]*string) *ModifyCuResponse {
	s.Headers = v
	return s
}

func (s *ModifyCuResponse) SetStatusCode(v int32) *ModifyCuResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCuResponse) SetBody(v *ModifyCuResponseBody) *ModifyCuResponse {
	s.Body = v
	return s
}

func (s *ModifyCuResponse) Validate() error {
	return dara.Validate(s)
}
