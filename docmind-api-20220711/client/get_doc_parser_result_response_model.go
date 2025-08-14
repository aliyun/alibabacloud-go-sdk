// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocParserResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocParserResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocParserResultResponse
	GetStatusCode() *int32
	SetBody(v *GetDocParserResultResponseBody) *GetDocParserResultResponse
	GetBody() *GetDocParserResultResponseBody
}

type GetDocParserResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocParserResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocParserResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocParserResultResponse) GoString() string {
	return s.String()
}

func (s *GetDocParserResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocParserResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocParserResultResponse) GetBody() *GetDocParserResultResponseBody {
	return s.Body
}

func (s *GetDocParserResultResponse) SetHeaders(v map[string]*string) *GetDocParserResultResponse {
	s.Headers = v
	return s
}

func (s *GetDocParserResultResponse) SetStatusCode(v int32) *GetDocParserResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocParserResultResponse) SetBody(v *GetDocParserResultResponseBody) *GetDocParserResultResponse {
	s.Body = v
	return s
}

func (s *GetDocParserResultResponse) Validate() error {
	return dara.Validate(s)
}
