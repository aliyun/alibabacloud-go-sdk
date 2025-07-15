// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportDoNotCallNumbersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportDoNotCallNumbersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportDoNotCallNumbersResponse
	GetStatusCode() *int32
	SetBody(v *ImportDoNotCallNumbersResponseBody) *ImportDoNotCallNumbersResponse
	GetBody() *ImportDoNotCallNumbersResponseBody
}

type ImportDoNotCallNumbersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportDoNotCallNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportDoNotCallNumbersResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportDoNotCallNumbersResponse) GoString() string {
	return s.String()
}

func (s *ImportDoNotCallNumbersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportDoNotCallNumbersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportDoNotCallNumbersResponse) GetBody() *ImportDoNotCallNumbersResponseBody {
	return s.Body
}

func (s *ImportDoNotCallNumbersResponse) SetHeaders(v map[string]*string) *ImportDoNotCallNumbersResponse {
	s.Headers = v
	return s
}

func (s *ImportDoNotCallNumbersResponse) SetStatusCode(v int32) *ImportDoNotCallNumbersResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportDoNotCallNumbersResponse) SetBody(v *ImportDoNotCallNumbersResponseBody) *ImportDoNotCallNumbersResponse {
	s.Body = v
	return s
}

func (s *ImportDoNotCallNumbersResponse) Validate() error {
	return dara.Validate(s)
}
