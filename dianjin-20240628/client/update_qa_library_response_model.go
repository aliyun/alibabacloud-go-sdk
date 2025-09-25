// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQaLibraryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateQaLibraryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateQaLibraryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateQaLibraryResponseBody) *UpdateQaLibraryResponse
	GetBody() *UpdateQaLibraryResponseBody
}

type UpdateQaLibraryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQaLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQaLibraryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateQaLibraryResponse) GoString() string {
	return s.String()
}

func (s *UpdateQaLibraryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateQaLibraryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateQaLibraryResponse) GetBody() *UpdateQaLibraryResponseBody {
	return s.Body
}

func (s *UpdateQaLibraryResponse) SetHeaders(v map[string]*string) *UpdateQaLibraryResponse {
	s.Headers = v
	return s
}

func (s *UpdateQaLibraryResponse) SetStatusCode(v int32) *UpdateQaLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQaLibraryResponse) SetBody(v *UpdateQaLibraryResponseBody) *UpdateQaLibraryResponse {
	s.Body = v
	return s
}

func (s *UpdateQaLibraryResponse) Validate() error {
	return dara.Validate(s)
}
