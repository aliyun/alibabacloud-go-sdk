// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateSqlBySemanticSqlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateSqlBySemanticSqlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateSqlBySemanticSqlResponse
	GetStatusCode() *int32
	SetBody(v *GenerateSqlBySemanticSqlResponseBody) *GenerateSqlBySemanticSqlResponse
	GetBody() *GenerateSqlBySemanticSqlResponseBody
}

type GenerateSqlBySemanticSqlResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateSqlBySemanticSqlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateSqlBySemanticSqlResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateSqlBySemanticSqlResponse) GoString() string {
	return s.String()
}

func (s *GenerateSqlBySemanticSqlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateSqlBySemanticSqlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateSqlBySemanticSqlResponse) GetBody() *GenerateSqlBySemanticSqlResponseBody {
	return s.Body
}

func (s *GenerateSqlBySemanticSqlResponse) SetHeaders(v map[string]*string) *GenerateSqlBySemanticSqlResponse {
	s.Headers = v
	return s
}

func (s *GenerateSqlBySemanticSqlResponse) SetStatusCode(v int32) *GenerateSqlBySemanticSqlResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateSqlBySemanticSqlResponse) SetBody(v *GenerateSqlBySemanticSqlResponseBody) *GenerateSqlBySemanticSqlResponse {
	s.Body = v
	return s
}

func (s *GenerateSqlBySemanticSqlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
