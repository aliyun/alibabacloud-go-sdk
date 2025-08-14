// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHotwordLibraryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHotwordLibraryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHotwordLibraryResponse
	GetStatusCode() *int32
	SetBody(v *CreateHotwordLibraryResponseBody) *CreateHotwordLibraryResponse
	GetBody() *CreateHotwordLibraryResponseBody
}

type CreateHotwordLibraryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHotwordLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHotwordLibraryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHotwordLibraryResponse) GoString() string {
	return s.String()
}

func (s *CreateHotwordLibraryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHotwordLibraryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHotwordLibraryResponse) GetBody() *CreateHotwordLibraryResponseBody {
	return s.Body
}

func (s *CreateHotwordLibraryResponse) SetHeaders(v map[string]*string) *CreateHotwordLibraryResponse {
	s.Headers = v
	return s
}

func (s *CreateHotwordLibraryResponse) SetStatusCode(v int32) *CreateHotwordLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHotwordLibraryResponse) SetBody(v *CreateHotwordLibraryResponseBody) *CreateHotwordLibraryResponse {
	s.Body = v
	return s
}

func (s *CreateHotwordLibraryResponse) Validate() error {
	return dara.Validate(s)
}
