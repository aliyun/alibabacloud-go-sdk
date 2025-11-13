// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterScannerYamlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClusterScannerYamlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClusterScannerYamlResponse
	GetStatusCode() *int32
	SetBody(v *GetClusterScannerYamlResponseBody) *GetClusterScannerYamlResponse
	GetBody() *GetClusterScannerYamlResponseBody
}

type GetClusterScannerYamlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterScannerYamlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterScannerYamlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClusterScannerYamlResponse) GoString() string {
	return s.String()
}

func (s *GetClusterScannerYamlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClusterScannerYamlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClusterScannerYamlResponse) GetBody() *GetClusterScannerYamlResponseBody {
	return s.Body
}

func (s *GetClusterScannerYamlResponse) SetHeaders(v map[string]*string) *GetClusterScannerYamlResponse {
	s.Headers = v
	return s
}

func (s *GetClusterScannerYamlResponse) SetStatusCode(v int32) *GetClusterScannerYamlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterScannerYamlResponse) SetBody(v *GetClusterScannerYamlResponseBody) *GetClusterScannerYamlResponse {
	s.Body = v
	return s
}

func (s *GetClusterScannerYamlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
