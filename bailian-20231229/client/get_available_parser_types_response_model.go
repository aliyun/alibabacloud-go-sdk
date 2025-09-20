// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAvailableParserTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAvailableParserTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAvailableParserTypesResponse
	GetStatusCode() *int32
	SetBody(v *GetAvailableParserTypesResponseBody) *GetAvailableParserTypesResponse
	GetBody() *GetAvailableParserTypesResponseBody
}

type GetAvailableParserTypesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAvailableParserTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAvailableParserTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAvailableParserTypesResponse) GoString() string {
	return s.String()
}

func (s *GetAvailableParserTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAvailableParserTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAvailableParserTypesResponse) GetBody() *GetAvailableParserTypesResponseBody {
	return s.Body
}

func (s *GetAvailableParserTypesResponse) SetHeaders(v map[string]*string) *GetAvailableParserTypesResponse {
	s.Headers = v
	return s
}

func (s *GetAvailableParserTypesResponse) SetStatusCode(v int32) *GetAvailableParserTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAvailableParserTypesResponse) SetBody(v *GetAvailableParserTypesResponseBody) *GetAvailableParserTypesResponse {
	s.Body = v
	return s
}

func (s *GetAvailableParserTypesResponse) Validate() error {
	return dara.Validate(s)
}
