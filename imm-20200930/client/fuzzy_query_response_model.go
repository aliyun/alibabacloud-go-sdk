// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFuzzyQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FuzzyQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FuzzyQueryResponse
	GetStatusCode() *int32
	SetBody(v *FuzzyQueryResponseBody) *FuzzyQueryResponse
	GetBody() *FuzzyQueryResponseBody
}

type FuzzyQueryResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FuzzyQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FuzzyQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s FuzzyQueryResponse) GoString() string {
	return s.String()
}

func (s *FuzzyQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FuzzyQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FuzzyQueryResponse) GetBody() *FuzzyQueryResponseBody {
	return s.Body
}

func (s *FuzzyQueryResponse) SetHeaders(v map[string]*string) *FuzzyQueryResponse {
	s.Headers = v
	return s
}

func (s *FuzzyQueryResponse) SetStatusCode(v int32) *FuzzyQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *FuzzyQueryResponse) SetBody(v *FuzzyQueryResponseBody) *FuzzyQueryResponse {
	s.Body = v
	return s
}

func (s *FuzzyQueryResponse) Validate() error {
	return dara.Validate(s)
}
