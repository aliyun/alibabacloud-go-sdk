// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlConversionProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSqlConversionProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSqlConversionProgressResponse
	GetStatusCode() *int32
	SetBody(v *GetSqlConversionProgressResponseBody) *GetSqlConversionProgressResponse
	GetBody() *GetSqlConversionProgressResponseBody
}

type GetSqlConversionProgressResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSqlConversionProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSqlConversionProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConversionProgressResponse) GoString() string {
	return s.String()
}

func (s *GetSqlConversionProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSqlConversionProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSqlConversionProgressResponse) GetBody() *GetSqlConversionProgressResponseBody {
	return s.Body
}

func (s *GetSqlConversionProgressResponse) SetHeaders(v map[string]*string) *GetSqlConversionProgressResponse {
	s.Headers = v
	return s
}

func (s *GetSqlConversionProgressResponse) SetStatusCode(v int32) *GetSqlConversionProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSqlConversionProgressResponse) SetBody(v *GetSqlConversionProgressResponseBody) *GetSqlConversionProgressResponse {
	s.Body = v
	return s
}

func (s *GetSqlConversionProgressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
