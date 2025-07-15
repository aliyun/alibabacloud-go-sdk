// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearIntervenesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClearIntervenesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClearIntervenesResponse
	GetStatusCode() *int32
	SetBody(v *ClearIntervenesResponseBody) *ClearIntervenesResponse
	GetBody() *ClearIntervenesResponseBody
}

type ClearIntervenesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearIntervenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearIntervenesResponse) String() string {
	return dara.Prettify(s)
}

func (s ClearIntervenesResponse) GoString() string {
	return s.String()
}

func (s *ClearIntervenesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClearIntervenesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClearIntervenesResponse) GetBody() *ClearIntervenesResponseBody {
	return s.Body
}

func (s *ClearIntervenesResponse) SetHeaders(v map[string]*string) *ClearIntervenesResponse {
	s.Headers = v
	return s
}

func (s *ClearIntervenesResponse) SetStatusCode(v int32) *ClearIntervenesResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearIntervenesResponse) SetBody(v *ClearIntervenesResponseBody) *ClearIntervenesResponse {
	s.Body = v
	return s
}

func (s *ClearIntervenesResponse) Validate() error {
	return dara.Validate(s)
}
