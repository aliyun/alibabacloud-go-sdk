// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRotateClusterManagedCertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RotateClusterManagedCertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RotateClusterManagedCertResponse
	GetStatusCode() *int32
	SetBody(v *RotateClusterManagedCertResponseBody) *RotateClusterManagedCertResponse
	GetBody() *RotateClusterManagedCertResponseBody
}

type RotateClusterManagedCertResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RotateClusterManagedCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RotateClusterManagedCertResponse) String() string {
	return dara.Prettify(s)
}

func (s RotateClusterManagedCertResponse) GoString() string {
	return s.String()
}

func (s *RotateClusterManagedCertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RotateClusterManagedCertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RotateClusterManagedCertResponse) GetBody() *RotateClusterManagedCertResponseBody {
	return s.Body
}

func (s *RotateClusterManagedCertResponse) SetHeaders(v map[string]*string) *RotateClusterManagedCertResponse {
	s.Headers = v
	return s
}

func (s *RotateClusterManagedCertResponse) SetStatusCode(v int32) *RotateClusterManagedCertResponse {
	s.StatusCode = &v
	return s
}

func (s *RotateClusterManagedCertResponse) SetBody(v *RotateClusterManagedCertResponseBody) *RotateClusterManagedCertResponse {
	s.Body = v
	return s
}

func (s *RotateClusterManagedCertResponse) Validate() error {
	return dara.Validate(s)
}
