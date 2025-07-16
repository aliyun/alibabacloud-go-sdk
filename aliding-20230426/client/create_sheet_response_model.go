// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSheetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSheetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSheetResponse
	GetStatusCode() *int32
	SetBody(v *CreateSheetResponseBody) *CreateSheetResponse
	GetBody() *CreateSheetResponseBody
}

type CreateSheetResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSheetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSheetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSheetResponse) GoString() string {
	return s.String()
}

func (s *CreateSheetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSheetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSheetResponse) GetBody() *CreateSheetResponseBody {
	return s.Body
}

func (s *CreateSheetResponse) SetHeaders(v map[string]*string) *CreateSheetResponse {
	s.Headers = v
	return s
}

func (s *CreateSheetResponse) SetStatusCode(v int32) *CreateSheetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSheetResponse) SetBody(v *CreateSheetResponseBody) *CreateSheetResponse {
	s.Body = v
	return s
}

func (s *CreateSheetResponse) Validate() error {
	return dara.Validate(s)
}
