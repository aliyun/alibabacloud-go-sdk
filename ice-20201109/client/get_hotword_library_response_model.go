// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotwordLibraryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotwordLibraryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotwordLibraryResponse
	GetStatusCode() *int32
	SetBody(v *GetHotwordLibraryResponseBody) *GetHotwordLibraryResponse
	GetBody() *GetHotwordLibraryResponseBody
}

type GetHotwordLibraryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotwordLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotwordLibraryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotwordLibraryResponse) GoString() string {
	return s.String()
}

func (s *GetHotwordLibraryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotwordLibraryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotwordLibraryResponse) GetBody() *GetHotwordLibraryResponseBody {
	return s.Body
}

func (s *GetHotwordLibraryResponse) SetHeaders(v map[string]*string) *GetHotwordLibraryResponse {
	s.Headers = v
	return s
}

func (s *GetHotwordLibraryResponse) SetStatusCode(v int32) *GetHotwordLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotwordLibraryResponse) SetBody(v *GetHotwordLibraryResponseBody) *GetHotwordLibraryResponse {
	s.Body = v
	return s
}

func (s *GetHotwordLibraryResponse) Validate() error {
	return dara.Validate(s)
}
