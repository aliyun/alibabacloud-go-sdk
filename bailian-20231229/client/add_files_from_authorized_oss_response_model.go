// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFilesFromAuthorizedOssResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFilesFromAuthorizedOssResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFilesFromAuthorizedOssResponse
	GetStatusCode() *int32
	SetBody(v *AddFilesFromAuthorizedOssResponseBody) *AddFilesFromAuthorizedOssResponse
	GetBody() *AddFilesFromAuthorizedOssResponseBody
}

type AddFilesFromAuthorizedOssResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFilesFromAuthorizedOssResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFilesFromAuthorizedOssResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFilesFromAuthorizedOssResponse) GoString() string {
	return s.String()
}

func (s *AddFilesFromAuthorizedOssResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFilesFromAuthorizedOssResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFilesFromAuthorizedOssResponse) GetBody() *AddFilesFromAuthorizedOssResponseBody {
	return s.Body
}

func (s *AddFilesFromAuthorizedOssResponse) SetHeaders(v map[string]*string) *AddFilesFromAuthorizedOssResponse {
	s.Headers = v
	return s
}

func (s *AddFilesFromAuthorizedOssResponse) SetStatusCode(v int32) *AddFilesFromAuthorizedOssResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFilesFromAuthorizedOssResponse) SetBody(v *AddFilesFromAuthorizedOssResponseBody) *AddFilesFromAuthorizedOssResponse {
	s.Body = v
	return s
}

func (s *AddFilesFromAuthorizedOssResponse) Validate() error {
	return dara.Validate(s)
}
