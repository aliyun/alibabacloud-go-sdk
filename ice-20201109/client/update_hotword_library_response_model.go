// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotwordLibraryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHotwordLibraryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHotwordLibraryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHotwordLibraryResponseBody) *UpdateHotwordLibraryResponse
	GetBody() *UpdateHotwordLibraryResponseBody
}

type UpdateHotwordLibraryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHotwordLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHotwordLibraryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotwordLibraryResponse) GoString() string {
	return s.String()
}

func (s *UpdateHotwordLibraryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHotwordLibraryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHotwordLibraryResponse) GetBody() *UpdateHotwordLibraryResponseBody {
	return s.Body
}

func (s *UpdateHotwordLibraryResponse) SetHeaders(v map[string]*string) *UpdateHotwordLibraryResponse {
	s.Headers = v
	return s
}

func (s *UpdateHotwordLibraryResponse) SetStatusCode(v int32) *UpdateHotwordLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHotwordLibraryResponse) SetBody(v *UpdateHotwordLibraryResponseBody) *UpdateHotwordLibraryResponse {
	s.Body = v
	return s
}

func (s *UpdateHotwordLibraryResponse) Validate() error {
	return dara.Validate(s)
}
