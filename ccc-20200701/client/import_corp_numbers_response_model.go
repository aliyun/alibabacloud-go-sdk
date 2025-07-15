// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportCorpNumbersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportCorpNumbersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportCorpNumbersResponse
	GetStatusCode() *int32
	SetBody(v *ImportCorpNumbersResponseBody) *ImportCorpNumbersResponse
	GetBody() *ImportCorpNumbersResponseBody
}

type ImportCorpNumbersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportCorpNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportCorpNumbersResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportCorpNumbersResponse) GoString() string {
	return s.String()
}

func (s *ImportCorpNumbersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportCorpNumbersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportCorpNumbersResponse) GetBody() *ImportCorpNumbersResponseBody {
	return s.Body
}

func (s *ImportCorpNumbersResponse) SetHeaders(v map[string]*string) *ImportCorpNumbersResponse {
	s.Headers = v
	return s
}

func (s *ImportCorpNumbersResponse) SetStatusCode(v int32) *ImportCorpNumbersResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportCorpNumbersResponse) SetBody(v *ImportCorpNumbersResponseBody) *ImportCorpNumbersResponse {
	s.Body = v
	return s
}

func (s *ImportCorpNumbersResponse) Validate() error {
	return dara.Validate(s)
}
