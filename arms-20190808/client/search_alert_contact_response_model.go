// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAlertContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchAlertContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchAlertContactResponse
	GetStatusCode() *int32
	SetBody(v *SearchAlertContactResponseBody) *SearchAlertContactResponse
	GetBody() *SearchAlertContactResponseBody
}

type SearchAlertContactResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchAlertContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchAlertContactResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertContactResponse) GoString() string {
	return s.String()
}

func (s *SearchAlertContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchAlertContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchAlertContactResponse) GetBody() *SearchAlertContactResponseBody {
	return s.Body
}

func (s *SearchAlertContactResponse) SetHeaders(v map[string]*string) *SearchAlertContactResponse {
	s.Headers = v
	return s
}

func (s *SearchAlertContactResponse) SetStatusCode(v int32) *SearchAlertContactResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAlertContactResponse) SetBody(v *SearchAlertContactResponseBody) *SearchAlertContactResponse {
	s.Body = v
	return s
}

func (s *SearchAlertContactResponse) Validate() error {
	return dara.Validate(s)
}
