// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadCredentialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUploadCredentialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUploadCredentialsResponse
	GetStatusCode() *int32
	SetBody(v *GetUploadCredentialsResponseBody) *GetUploadCredentialsResponse
	GetBody() *GetUploadCredentialsResponseBody
}

type GetUploadCredentialsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUploadCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUploadCredentialsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUploadCredentialsResponse) GoString() string {
	return s.String()
}

func (s *GetUploadCredentialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUploadCredentialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUploadCredentialsResponse) GetBody() *GetUploadCredentialsResponseBody {
	return s.Body
}

func (s *GetUploadCredentialsResponse) SetHeaders(v map[string]*string) *GetUploadCredentialsResponse {
	s.Headers = v
	return s
}

func (s *GetUploadCredentialsResponse) SetStatusCode(v int32) *GetUploadCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUploadCredentialsResponse) SetBody(v *GetUploadCredentialsResponseBody) *GetUploadCredentialsResponse {
	s.Body = v
	return s
}

func (s *GetUploadCredentialsResponse) Validate() error {
	return dara.Validate(s)
}
