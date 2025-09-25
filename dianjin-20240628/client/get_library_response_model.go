// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLibraryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLibraryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLibraryResponse
	GetStatusCode() *int32
	SetBody(v *GetLibraryResponseBody) *GetLibraryResponse
	GetBody() *GetLibraryResponseBody
}

type GetLibraryResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLibraryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLibraryResponse) GoString() string {
	return s.String()
}

func (s *GetLibraryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLibraryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLibraryResponse) GetBody() *GetLibraryResponseBody {
	return s.Body
}

func (s *GetLibraryResponse) SetHeaders(v map[string]*string) *GetLibraryResponse {
	s.Headers = v
	return s
}

func (s *GetLibraryResponse) SetStatusCode(v int32) *GetLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLibraryResponse) SetBody(v *GetLibraryResponseBody) *GetLibraryResponse {
	s.Body = v
	return s
}

func (s *GetLibraryResponse) Validate() error {
	return dara.Validate(s)
}
