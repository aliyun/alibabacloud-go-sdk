// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentParseTestApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DocumentParseTestApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DocumentParseTestApiResponse
	GetStatusCode() *int32
	SetBody(v *DocumentParseTestApiResponseBody) *DocumentParseTestApiResponse
	GetBody() *DocumentParseTestApiResponseBody
}

type DocumentParseTestApiResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DocumentParseTestApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DocumentParseTestApiResponse) String() string {
	return dara.Prettify(s)
}

func (s DocumentParseTestApiResponse) GoString() string {
	return s.String()
}

func (s *DocumentParseTestApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DocumentParseTestApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DocumentParseTestApiResponse) GetBody() *DocumentParseTestApiResponseBody {
	return s.Body
}

func (s *DocumentParseTestApiResponse) SetHeaders(v map[string]*string) *DocumentParseTestApiResponse {
	s.Headers = v
	return s
}

func (s *DocumentParseTestApiResponse) SetStatusCode(v int32) *DocumentParseTestApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DocumentParseTestApiResponse) SetBody(v *DocumentParseTestApiResponseBody) *DocumentParseTestApiResponse {
	s.Body = v
	return s
}

func (s *DocumentParseTestApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
