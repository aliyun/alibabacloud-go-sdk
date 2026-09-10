// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyntaxCheckAndTransformSqlConversionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyntaxCheckAndTransformSqlConversionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyntaxCheckAndTransformSqlConversionTaskResponse
	GetStatusCode() *int32
	SetBody(v *SyntaxCheckAndTransformSqlConversionTaskResponseBody) *SyntaxCheckAndTransformSqlConversionTaskResponse
	GetBody() *SyntaxCheckAndTransformSqlConversionTaskResponseBody
}

type SyntaxCheckAndTransformSqlConversionTaskResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyntaxCheckAndTransformSqlConversionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyntaxCheckAndTransformSqlConversionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s SyntaxCheckAndTransformSqlConversionTaskResponse) GoString() string {
	return s.String()
}

func (s *SyntaxCheckAndTransformSqlConversionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyntaxCheckAndTransformSqlConversionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyntaxCheckAndTransformSqlConversionTaskResponse) GetBody() *SyntaxCheckAndTransformSqlConversionTaskResponseBody {
	return s.Body
}

func (s *SyntaxCheckAndTransformSqlConversionTaskResponse) SetHeaders(v map[string]*string) *SyntaxCheckAndTransformSqlConversionTaskResponse {
	s.Headers = v
	return s
}

func (s *SyntaxCheckAndTransformSqlConversionTaskResponse) SetStatusCode(v int32) *SyntaxCheckAndTransformSqlConversionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SyntaxCheckAndTransformSqlConversionTaskResponse) SetBody(v *SyntaxCheckAndTransformSqlConversionTaskResponseBody) *SyntaxCheckAndTransformSqlConversionTaskResponse {
	s.Body = v
	return s
}

func (s *SyntaxCheckAndTransformSqlConversionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
