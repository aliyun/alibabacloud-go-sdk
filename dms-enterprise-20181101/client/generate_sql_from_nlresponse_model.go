// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateSqlFromNLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateSqlFromNLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateSqlFromNLResponse
	GetStatusCode() *int32
	SetBody(v *GenerateSqlFromNLResponseBody) *GenerateSqlFromNLResponse
	GetBody() *GenerateSqlFromNLResponseBody
}

type GenerateSqlFromNLResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateSqlFromNLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateSqlFromNLResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateSqlFromNLResponse) GoString() string {
	return s.String()
}

func (s *GenerateSqlFromNLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateSqlFromNLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateSqlFromNLResponse) GetBody() *GenerateSqlFromNLResponseBody {
	return s.Body
}

func (s *GenerateSqlFromNLResponse) SetHeaders(v map[string]*string) *GenerateSqlFromNLResponse {
	s.Headers = v
	return s
}

func (s *GenerateSqlFromNLResponse) SetStatusCode(v int32) *GenerateSqlFromNLResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateSqlFromNLResponse) SetBody(v *GenerateSqlFromNLResponseBody) *GenerateSqlFromNLResponse {
	s.Body = v
	return s
}

func (s *GenerateSqlFromNLResponse) Validate() error {
	return dara.Validate(s)
}
