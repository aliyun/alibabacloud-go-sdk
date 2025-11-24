// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeshMultiClusterNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMeshMultiClusterNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMeshMultiClusterNetworkResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMeshMultiClusterNetworkResponseBody) *DescribeMeshMultiClusterNetworkResponse
	GetBody() *DescribeMeshMultiClusterNetworkResponseBody
}

type DescribeMeshMultiClusterNetworkResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMeshMultiClusterNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMeshMultiClusterNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeshMultiClusterNetworkResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeshMultiClusterNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMeshMultiClusterNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMeshMultiClusterNetworkResponse) GetBody() *DescribeMeshMultiClusterNetworkResponseBody {
	return s.Body
}

func (s *DescribeMeshMultiClusterNetworkResponse) SetHeaders(v map[string]*string) *DescribeMeshMultiClusterNetworkResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeshMultiClusterNetworkResponse) SetStatusCode(v int32) *DescribeMeshMultiClusterNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMeshMultiClusterNetworkResponse) SetBody(v *DescribeMeshMultiClusterNetworkResponseBody) *DescribeMeshMultiClusterNetworkResponse {
	s.Body = v
	return s
}

func (s *DescribeMeshMultiClusterNetworkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
