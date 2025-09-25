// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLibraryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLibraryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLibraryResponse
	GetStatusCode() *int32
	SetBody(v *CreateLibraryResponseBody) *CreateLibraryResponse
	GetBody() *CreateLibraryResponseBody
}

type CreateLibraryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLibraryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLibraryResponse) GoString() string {
	return s.String()
}

func (s *CreateLibraryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLibraryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLibraryResponse) GetBody() *CreateLibraryResponseBody {
	return s.Body
}

func (s *CreateLibraryResponse) SetHeaders(v map[string]*string) *CreateLibraryResponse {
	s.Headers = v
	return s
}

func (s *CreateLibraryResponse) SetStatusCode(v int32) *CreateLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLibraryResponse) SetBody(v *CreateLibraryResponseBody) *CreateLibraryResponse {
	s.Body = v
	return s
}

func (s *CreateLibraryResponse) Validate() error {
	return dara.Validate(s)
}
