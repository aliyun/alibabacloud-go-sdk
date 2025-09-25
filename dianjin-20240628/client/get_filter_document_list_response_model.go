// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFilterDocumentListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFilterDocumentListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFilterDocumentListResponse
	GetStatusCode() *int32
	SetBody(v *GetFilterDocumentListResponseBody) *GetFilterDocumentListResponse
	GetBody() *GetFilterDocumentListResponseBody
}

type GetFilterDocumentListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFilterDocumentListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFilterDocumentListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFilterDocumentListResponse) GoString() string {
	return s.String()
}

func (s *GetFilterDocumentListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFilterDocumentListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFilterDocumentListResponse) GetBody() *GetFilterDocumentListResponseBody {
	return s.Body
}

func (s *GetFilterDocumentListResponse) SetHeaders(v map[string]*string) *GetFilterDocumentListResponse {
	s.Headers = v
	return s
}

func (s *GetFilterDocumentListResponse) SetStatusCode(v int32) *GetFilterDocumentListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFilterDocumentListResponse) SetBody(v *GetFilterDocumentListResponseBody) *GetFilterDocumentListResponse {
	s.Body = v
	return s
}

func (s *GetFilterDocumentListResponse) Validate() error {
	return dara.Validate(s)
}
