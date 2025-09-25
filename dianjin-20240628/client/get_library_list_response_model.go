// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLibraryListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLibraryListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLibraryListResponse
	GetStatusCode() *int32
	SetBody(v *GetLibraryListResponseBody) *GetLibraryListResponse
	GetBody() *GetLibraryListResponseBody
}

type GetLibraryListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLibraryListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLibraryListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLibraryListResponse) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLibraryListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLibraryListResponse) GetBody() *GetLibraryListResponseBody {
	return s.Body
}

func (s *GetLibraryListResponse) SetHeaders(v map[string]*string) *GetLibraryListResponse {
	s.Headers = v
	return s
}

func (s *GetLibraryListResponse) SetStatusCode(v int32) *GetLibraryListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLibraryListResponse) SetBody(v *GetLibraryListResponseBody) *GetLibraryListResponse {
	s.Body = v
	return s
}

func (s *GetLibraryListResponse) Validate() error {
	return dara.Validate(s)
}
