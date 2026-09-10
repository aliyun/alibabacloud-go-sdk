// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostInnerUploadConvertPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PostInnerUploadConvertPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PostInnerUploadConvertPackageResponse
	GetStatusCode() *int32
	SetBody(v *PostInnerUploadConvertPackageResponseBody) *PostInnerUploadConvertPackageResponse
	GetBody() *PostInnerUploadConvertPackageResponseBody
}

type PostInnerUploadConvertPackageResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PostInnerUploadConvertPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PostInnerUploadConvertPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s PostInnerUploadConvertPackageResponse) GoString() string {
	return s.String()
}

func (s *PostInnerUploadConvertPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PostInnerUploadConvertPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PostInnerUploadConvertPackageResponse) GetBody() *PostInnerUploadConvertPackageResponseBody {
	return s.Body
}

func (s *PostInnerUploadConvertPackageResponse) SetHeaders(v map[string]*string) *PostInnerUploadConvertPackageResponse {
	s.Headers = v
	return s
}

func (s *PostInnerUploadConvertPackageResponse) SetStatusCode(v int32) *PostInnerUploadConvertPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *PostInnerUploadConvertPackageResponse) SetBody(v *PostInnerUploadConvertPackageResponseBody) *PostInnerUploadConvertPackageResponse {
	s.Body = v
	return s
}

func (s *PostInnerUploadConvertPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
