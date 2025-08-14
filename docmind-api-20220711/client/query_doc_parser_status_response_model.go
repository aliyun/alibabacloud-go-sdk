// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDocParserStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDocParserStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDocParserStatusResponse
	GetStatusCode() *int32
	SetBody(v *QueryDocParserStatusResponseBody) *QueryDocParserStatusResponse
	GetBody() *QueryDocParserStatusResponseBody
}

type QueryDocParserStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDocParserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDocParserStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDocParserStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryDocParserStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDocParserStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDocParserStatusResponse) GetBody() *QueryDocParserStatusResponseBody {
	return s.Body
}

func (s *QueryDocParserStatusResponse) SetHeaders(v map[string]*string) *QueryDocParserStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryDocParserStatusResponse) SetStatusCode(v int32) *QueryDocParserStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDocParserStatusResponse) SetBody(v *QueryDocParserStatusResponseBody) *QueryDocParserStatusResponse {
	s.Body = v
	return s
}

func (s *QueryDocParserStatusResponse) Validate() error {
	return dara.Validate(s)
}
