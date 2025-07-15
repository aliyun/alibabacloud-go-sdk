// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoNotCallFileUploadParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDoNotCallFileUploadParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDoNotCallFileUploadParametersResponse
	GetStatusCode() *int32
	SetBody(v *GetDoNotCallFileUploadParametersResponseBody) *GetDoNotCallFileUploadParametersResponse
	GetBody() *GetDoNotCallFileUploadParametersResponseBody
}

type GetDoNotCallFileUploadParametersResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDoNotCallFileUploadParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDoNotCallFileUploadParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDoNotCallFileUploadParametersResponse) GoString() string {
	return s.String()
}

func (s *GetDoNotCallFileUploadParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDoNotCallFileUploadParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDoNotCallFileUploadParametersResponse) GetBody() *GetDoNotCallFileUploadParametersResponseBody {
	return s.Body
}

func (s *GetDoNotCallFileUploadParametersResponse) SetHeaders(v map[string]*string) *GetDoNotCallFileUploadParametersResponse {
	s.Headers = v
	return s
}

func (s *GetDoNotCallFileUploadParametersResponse) SetStatusCode(v int32) *GetDoNotCallFileUploadParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDoNotCallFileUploadParametersResponse) SetBody(v *GetDoNotCallFileUploadParametersResponseBody) *GetDoNotCallFileUploadParametersResponse {
	s.Body = v
	return s
}

func (s *GetDoNotCallFileUploadParametersResponse) Validate() error {
	return dara.Validate(s)
}
