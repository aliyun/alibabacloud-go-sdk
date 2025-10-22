// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataLakeFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataLakeFunctionResponse
	GetStatusCode() *int32
	SetBody(v *ListDataLakeFunctionResponseBody) *ListDataLakeFunctionResponse
	GetBody() *ListDataLakeFunctionResponseBody
}

type ListDataLakeFunctionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataLakeFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataLakeFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeFunctionResponse) GoString() string {
	return s.String()
}

func (s *ListDataLakeFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataLakeFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataLakeFunctionResponse) GetBody() *ListDataLakeFunctionResponseBody {
	return s.Body
}

func (s *ListDataLakeFunctionResponse) SetHeaders(v map[string]*string) *ListDataLakeFunctionResponse {
	s.Headers = v
	return s
}

func (s *ListDataLakeFunctionResponse) SetStatusCode(v int32) *ListDataLakeFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataLakeFunctionResponse) SetBody(v *ListDataLakeFunctionResponseBody) *ListDataLakeFunctionResponse {
	s.Body = v
	return s
}

func (s *ListDataLakeFunctionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
