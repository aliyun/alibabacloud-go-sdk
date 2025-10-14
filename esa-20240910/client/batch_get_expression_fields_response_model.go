// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetExpressionFieldsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchGetExpressionFieldsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchGetExpressionFieldsResponse
	GetStatusCode() *int32
	SetBody(v *BatchGetExpressionFieldsResponseBody) *BatchGetExpressionFieldsResponse
	GetBody() *BatchGetExpressionFieldsResponseBody
}

type BatchGetExpressionFieldsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetExpressionFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetExpressionFieldsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchGetExpressionFieldsResponse) GoString() string {
	return s.String()
}

func (s *BatchGetExpressionFieldsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchGetExpressionFieldsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchGetExpressionFieldsResponse) GetBody() *BatchGetExpressionFieldsResponseBody {
	return s.Body
}

func (s *BatchGetExpressionFieldsResponse) SetHeaders(v map[string]*string) *BatchGetExpressionFieldsResponse {
	s.Headers = v
	return s
}

func (s *BatchGetExpressionFieldsResponse) SetStatusCode(v int32) *BatchGetExpressionFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetExpressionFieldsResponse) SetBody(v *BatchGetExpressionFieldsResponseBody) *BatchGetExpressionFieldsResponse {
	s.Body = v
	return s
}

func (s *BatchGetExpressionFieldsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
