// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskReplicaGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiskReplicaGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiskReplicaGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiskReplicaGroupsResponseBody) *DescribeDiskReplicaGroupsResponse
	GetBody() *DescribeDiskReplicaGroupsResponseBody
}

type DescribeDiskReplicaGroupsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiskReplicaGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiskReplicaGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskReplicaGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiskReplicaGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiskReplicaGroupsResponse) GetBody() *DescribeDiskReplicaGroupsResponseBody {
	return s.Body
}

func (s *DescribeDiskReplicaGroupsResponse) SetHeaders(v map[string]*string) *DescribeDiskReplicaGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskReplicaGroupsResponse) SetStatusCode(v int32) *DescribeDiskReplicaGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiskReplicaGroupsResponse) SetBody(v *DescribeDiskReplicaGroupsResponseBody) *DescribeDiskReplicaGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeDiskReplicaGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
