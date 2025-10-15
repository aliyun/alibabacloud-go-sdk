// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLibraryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLibraryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLibraryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLibraryResponseBody) *UpdateLibraryResponse
	GetBody() *UpdateLibraryResponseBody
}

type UpdateLibraryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLibraryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLibraryResponse) GoString() string {
	return s.String()
}

func (s *UpdateLibraryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLibraryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLibraryResponse) GetBody() *UpdateLibraryResponseBody {
	return s.Body
}

func (s *UpdateLibraryResponse) SetHeaders(v map[string]*string) *UpdateLibraryResponse {
	s.Headers = v
	return s
}

func (s *UpdateLibraryResponse) SetStatusCode(v int32) *UpdateLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLibraryResponse) SetBody(v *UpdateLibraryResponseBody) *UpdateLibraryResponse {
	s.Body = v
	return s
}

func (s *UpdateLibraryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
