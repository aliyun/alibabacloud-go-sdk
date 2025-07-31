// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUploadLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUploadLinkResponse
	GetStatusCode() *int32
	SetBody(v *GetUploadLinkResponseBody) *GetUploadLinkResponse
	GetBody() *GetUploadLinkResponseBody
}

type GetUploadLinkResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUploadLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUploadLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUploadLinkResponse) GoString() string {
	return s.String()
}

func (s *GetUploadLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUploadLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUploadLinkResponse) GetBody() *GetUploadLinkResponseBody {
	return s.Body
}

func (s *GetUploadLinkResponse) SetHeaders(v map[string]*string) *GetUploadLinkResponse {
	s.Headers = v
	return s
}

func (s *GetUploadLinkResponse) SetStatusCode(v int32) *GetUploadLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUploadLinkResponse) SetBody(v *GetUploadLinkResponseBody) *GetUploadLinkResponse {
	s.Body = v
	return s
}

func (s *GetUploadLinkResponse) Validate() error {
	return dara.Validate(s)
}
