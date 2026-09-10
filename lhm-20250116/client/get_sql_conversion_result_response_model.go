// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlConversionResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSqlConversionResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSqlConversionResultResponse
	GetStatusCode() *int32
	SetBody(v *GetSqlConversionResultResponseBody) *GetSqlConversionResultResponse
	GetBody() *GetSqlConversionResultResponseBody
}

type GetSqlConversionResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSqlConversionResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSqlConversionResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConversionResultResponse) GoString() string {
	return s.String()
}

func (s *GetSqlConversionResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSqlConversionResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSqlConversionResultResponse) GetBody() *GetSqlConversionResultResponseBody {
	return s.Body
}

func (s *GetSqlConversionResultResponse) SetHeaders(v map[string]*string) *GetSqlConversionResultResponse {
	s.Headers = v
	return s
}

func (s *GetSqlConversionResultResponse) SetStatusCode(v int32) *GetSqlConversionResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSqlConversionResultResponse) SetBody(v *GetSqlConversionResultResponseBody) *GetSqlConversionResultResponse {
	s.Body = v
	return s
}

func (s *GetSqlConversionResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
