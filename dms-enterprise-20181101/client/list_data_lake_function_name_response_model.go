// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeFunctionNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataLakeFunctionNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataLakeFunctionNameResponse
	GetStatusCode() *int32
	SetBody(v *ListDataLakeFunctionNameResponseBody) *ListDataLakeFunctionNameResponse
	GetBody() *ListDataLakeFunctionNameResponseBody
}

type ListDataLakeFunctionNameResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataLakeFunctionNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataLakeFunctionNameResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeFunctionNameResponse) GoString() string {
	return s.String()
}

func (s *ListDataLakeFunctionNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataLakeFunctionNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataLakeFunctionNameResponse) GetBody() *ListDataLakeFunctionNameResponseBody {
	return s.Body
}

func (s *ListDataLakeFunctionNameResponse) SetHeaders(v map[string]*string) *ListDataLakeFunctionNameResponse {
	s.Headers = v
	return s
}

func (s *ListDataLakeFunctionNameResponse) SetStatusCode(v int32) *ListDataLakeFunctionNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataLakeFunctionNameResponse) SetBody(v *ListDataLakeFunctionNameResponseBody) *ListDataLakeFunctionNameResponse {
	s.Body = v
	return s
}

func (s *ListDataLakeFunctionNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
