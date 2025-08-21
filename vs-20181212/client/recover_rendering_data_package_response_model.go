// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverRenderingDataPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecoverRenderingDataPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecoverRenderingDataPackageResponse
	GetStatusCode() *int32
	SetBody(v *RecoverRenderingDataPackageResponseBody) *RecoverRenderingDataPackageResponse
	GetBody() *RecoverRenderingDataPackageResponseBody
}

type RecoverRenderingDataPackageResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecoverRenderingDataPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecoverRenderingDataPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s RecoverRenderingDataPackageResponse) GoString() string {
	return s.String()
}

func (s *RecoverRenderingDataPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecoverRenderingDataPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecoverRenderingDataPackageResponse) GetBody() *RecoverRenderingDataPackageResponseBody {
	return s.Body
}

func (s *RecoverRenderingDataPackageResponse) SetHeaders(v map[string]*string) *RecoverRenderingDataPackageResponse {
	s.Headers = v
	return s
}

func (s *RecoverRenderingDataPackageResponse) SetStatusCode(v int32) *RecoverRenderingDataPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoverRenderingDataPackageResponse) SetBody(v *RecoverRenderingDataPackageResponseBody) *RecoverRenderingDataPackageResponse {
	s.Body = v
	return s
}

func (s *RecoverRenderingDataPackageResponse) Validate() error {
	return dara.Validate(s)
}
