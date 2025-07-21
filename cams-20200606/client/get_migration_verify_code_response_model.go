// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMigrationVerifyCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMigrationVerifyCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMigrationVerifyCodeResponse
	GetStatusCode() *int32
	SetBody(v *GetMigrationVerifyCodeResponseBody) *GetMigrationVerifyCodeResponse
	GetBody() *GetMigrationVerifyCodeResponseBody
}

type GetMigrationVerifyCodeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMigrationVerifyCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMigrationVerifyCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationVerifyCodeResponse) GoString() string {
	return s.String()
}

func (s *GetMigrationVerifyCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMigrationVerifyCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMigrationVerifyCodeResponse) GetBody() *GetMigrationVerifyCodeResponseBody {
	return s.Body
}

func (s *GetMigrationVerifyCodeResponse) SetHeaders(v map[string]*string) *GetMigrationVerifyCodeResponse {
	s.Headers = v
	return s
}

func (s *GetMigrationVerifyCodeResponse) SetStatusCode(v int32) *GetMigrationVerifyCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMigrationVerifyCodeResponse) SetBody(v *GetMigrationVerifyCodeResponseBody) *GetMigrationVerifyCodeResponse {
	s.Body = v
	return s
}

func (s *GetMigrationVerifyCodeResponse) Validate() error {
	return dara.Validate(s)
}
