// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlOptimizeAdviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSqlOptimizeAdviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSqlOptimizeAdviceResponse
	GetStatusCode() *int32
	SetBody(v *GetSqlOptimizeAdviceResponseBody) *GetSqlOptimizeAdviceResponse
	GetBody() *GetSqlOptimizeAdviceResponseBody
}

type GetSqlOptimizeAdviceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSqlOptimizeAdviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSqlOptimizeAdviceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSqlOptimizeAdviceResponse) GoString() string {
	return s.String()
}

func (s *GetSqlOptimizeAdviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSqlOptimizeAdviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSqlOptimizeAdviceResponse) GetBody() *GetSqlOptimizeAdviceResponseBody {
	return s.Body
}

func (s *GetSqlOptimizeAdviceResponse) SetHeaders(v map[string]*string) *GetSqlOptimizeAdviceResponse {
	s.Headers = v
	return s
}

func (s *GetSqlOptimizeAdviceResponse) SetStatusCode(v int32) *GetSqlOptimizeAdviceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSqlOptimizeAdviceResponse) SetBody(v *GetSqlOptimizeAdviceResponseBody) *GetSqlOptimizeAdviceResponse {
	s.Body = v
	return s
}

func (s *GetSqlOptimizeAdviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
