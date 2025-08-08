// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMcubeMiniPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadMcubeMiniPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadMcubeMiniPackageResponse
	GetStatusCode() *int32
	SetBody(v *UploadMcubeMiniPackageResponseBody) *UploadMcubeMiniPackageResponse
	GetBody() *UploadMcubeMiniPackageResponseBody
}

type UploadMcubeMiniPackageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadMcubeMiniPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadMcubeMiniPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadMcubeMiniPackageResponse) GoString() string {
	return s.String()
}

func (s *UploadMcubeMiniPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadMcubeMiniPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadMcubeMiniPackageResponse) GetBody() *UploadMcubeMiniPackageResponseBody {
	return s.Body
}

func (s *UploadMcubeMiniPackageResponse) SetHeaders(v map[string]*string) *UploadMcubeMiniPackageResponse {
	s.Headers = v
	return s
}

func (s *UploadMcubeMiniPackageResponse) SetStatusCode(v int32) *UploadMcubeMiniPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadMcubeMiniPackageResponse) SetBody(v *UploadMcubeMiniPackageResponseBody) *UploadMcubeMiniPackageResponse {
	s.Body = v
	return s
}

func (s *UploadMcubeMiniPackageResponse) Validate() error {
	return dara.Validate(s)
}
