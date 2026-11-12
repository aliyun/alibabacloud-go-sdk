// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesTrafficPackagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstancesTrafficPackagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstancesTrafficPackagesResponse
	GetStatusCode() *int32
	SetBody(v *ListInstancesTrafficPackagesResponseBody) *ListInstancesTrafficPackagesResponse
	GetBody() *ListInstancesTrafficPackagesResponseBody
}

type ListInstancesTrafficPackagesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesTrafficPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesTrafficPackagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesTrafficPackagesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesTrafficPackagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstancesTrafficPackagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstancesTrafficPackagesResponse) GetBody() *ListInstancesTrafficPackagesResponseBody {
	return s.Body
}

func (s *ListInstancesTrafficPackagesResponse) SetHeaders(v map[string]*string) *ListInstancesTrafficPackagesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesTrafficPackagesResponse) SetStatusCode(v int32) *ListInstancesTrafficPackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesTrafficPackagesResponse) SetBody(v *ListInstancesTrafficPackagesResponseBody) *ListInstancesTrafficPackagesResponse {
	s.Body = v
	return s
}

func (s *ListInstancesTrafficPackagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
