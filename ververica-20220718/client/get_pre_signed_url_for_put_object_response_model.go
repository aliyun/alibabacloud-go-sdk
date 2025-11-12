// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPreSignedUrlForPutObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPreSignedUrlForPutObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPreSignedUrlForPutObjectResponse
	GetStatusCode() *int32
	SetBody(v *GetPreSignedUrlForPutObjectResponseBody) *GetPreSignedUrlForPutObjectResponse
	GetBody() *GetPreSignedUrlForPutObjectResponseBody
}

type GetPreSignedUrlForPutObjectResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPreSignedUrlForPutObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPreSignedUrlForPutObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPreSignedUrlForPutObjectResponse) GoString() string {
	return s.String()
}

func (s *GetPreSignedUrlForPutObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPreSignedUrlForPutObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPreSignedUrlForPutObjectResponse) GetBody() *GetPreSignedUrlForPutObjectResponseBody {
	return s.Body
}

func (s *GetPreSignedUrlForPutObjectResponse) SetHeaders(v map[string]*string) *GetPreSignedUrlForPutObjectResponse {
	s.Headers = v
	return s
}

func (s *GetPreSignedUrlForPutObjectResponse) SetStatusCode(v int32) *GetPreSignedUrlForPutObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPreSignedUrlForPutObjectResponse) SetBody(v *GetPreSignedUrlForPutObjectResponseBody) *GetPreSignedUrlForPutObjectResponse {
	s.Body = v
	return s
}

func (s *GetPreSignedUrlForPutObjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
