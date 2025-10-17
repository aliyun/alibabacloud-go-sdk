// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOssSubDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOssSubDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOssSubDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateOssSubDirectoryResponseBody) *CreateOssSubDirectoryResponse
	GetBody() *CreateOssSubDirectoryResponseBody
}

type CreateOssSubDirectoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOssSubDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOssSubDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOssSubDirectoryResponse) GoString() string {
	return s.String()
}

func (s *CreateOssSubDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOssSubDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOssSubDirectoryResponse) GetBody() *CreateOssSubDirectoryResponseBody {
	return s.Body
}

func (s *CreateOssSubDirectoryResponse) SetHeaders(v map[string]*string) *CreateOssSubDirectoryResponse {
	s.Headers = v
	return s
}

func (s *CreateOssSubDirectoryResponse) SetStatusCode(v int32) *CreateOssSubDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOssSubDirectoryResponse) SetBody(v *CreateOssSubDirectoryResponseBody) *CreateOssSubDirectoryResponse {
	s.Body = v
	return s
}

func (s *CreateOssSubDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
