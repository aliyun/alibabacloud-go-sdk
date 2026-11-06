// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCorpNumbersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCorpNumbersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCorpNumbersResponse
	GetStatusCode() *int32
	SetBody(v *ListCorpNumbersResponseBody) *ListCorpNumbersResponse
	GetBody() *ListCorpNumbersResponseBody
}

type ListCorpNumbersResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCorpNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCorpNumbersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCorpNumbersResponse) GoString() string {
	return s.String()
}

func (s *ListCorpNumbersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCorpNumbersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCorpNumbersResponse) GetBody() *ListCorpNumbersResponseBody {
	return s.Body
}

func (s *ListCorpNumbersResponse) SetHeaders(v map[string]*string) *ListCorpNumbersResponse {
	s.Headers = v
	return s
}

func (s *ListCorpNumbersResponse) SetStatusCode(v int32) *ListCorpNumbersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCorpNumbersResponse) SetBody(v *ListCorpNumbersResponseBody) *ListCorpNumbersResponse {
	s.Body = v
	return s
}

func (s *ListCorpNumbersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
