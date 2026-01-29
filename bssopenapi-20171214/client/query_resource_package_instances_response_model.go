// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryResourcePackageInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryResourcePackageInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryResourcePackageInstancesResponse
	GetStatusCode() *int32
	SetBody(v *QueryResourcePackageInstancesResponseBody) *QueryResourcePackageInstancesResponse
	GetBody() *QueryResourcePackageInstancesResponseBody
}

type QueryResourcePackageInstancesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryResourcePackageInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryResourcePackageInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryResourcePackageInstancesResponse) GoString() string {
	return s.String()
}

func (s *QueryResourcePackageInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryResourcePackageInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryResourcePackageInstancesResponse) GetBody() *QueryResourcePackageInstancesResponseBody {
	return s.Body
}

func (s *QueryResourcePackageInstancesResponse) SetHeaders(v map[string]*string) *QueryResourcePackageInstancesResponse {
	s.Headers = v
	return s
}

func (s *QueryResourcePackageInstancesResponse) SetStatusCode(v int32) *QueryResourcePackageInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryResourcePackageInstancesResponse) SetBody(v *QueryResourcePackageInstancesResponseBody) *QueryResourcePackageInstancesResponse {
	s.Body = v
	return s
}

func (s *QueryResourcePackageInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
