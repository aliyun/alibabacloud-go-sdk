// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetADBSparkNecessaryRAMPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetADBSparkNecessaryRAMPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetADBSparkNecessaryRAMPermissionsResponse
	GetStatusCode() *int32
	SetBody(v *GetADBSparkNecessaryRAMPermissionsResponseBody) *GetADBSparkNecessaryRAMPermissionsResponse
	GetBody() *GetADBSparkNecessaryRAMPermissionsResponseBody
}

type GetADBSparkNecessaryRAMPermissionsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetADBSparkNecessaryRAMPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetADBSparkNecessaryRAMPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetADBSparkNecessaryRAMPermissionsResponse) GoString() string {
	return s.String()
}

func (s *GetADBSparkNecessaryRAMPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetADBSparkNecessaryRAMPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetADBSparkNecessaryRAMPermissionsResponse) GetBody() *GetADBSparkNecessaryRAMPermissionsResponseBody {
	return s.Body
}

func (s *GetADBSparkNecessaryRAMPermissionsResponse) SetHeaders(v map[string]*string) *GetADBSparkNecessaryRAMPermissionsResponse {
	s.Headers = v
	return s
}

func (s *GetADBSparkNecessaryRAMPermissionsResponse) SetStatusCode(v int32) *GetADBSparkNecessaryRAMPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetADBSparkNecessaryRAMPermissionsResponse) SetBody(v *GetADBSparkNecessaryRAMPermissionsResponseBody) *GetADBSparkNecessaryRAMPermissionsResponse {
	s.Body = v
	return s
}

func (s *GetADBSparkNecessaryRAMPermissionsResponse) Validate() error {
	return dara.Validate(s)
}
