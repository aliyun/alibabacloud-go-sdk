// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMasterSlaveServerGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMasterSlaveServerGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMasterSlaveServerGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMasterSlaveServerGroupsResponseBody) *DescribeMasterSlaveServerGroupsResponse
	GetBody() *DescribeMasterSlaveServerGroupsResponseBody
}

type DescribeMasterSlaveServerGroupsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMasterSlaveServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMasterSlaveServerGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMasterSlaveServerGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMasterSlaveServerGroupsResponse) GetBody() *DescribeMasterSlaveServerGroupsResponseBody {
	return s.Body
}

func (s *DescribeMasterSlaveServerGroupsResponse) SetHeaders(v map[string]*string) *DescribeMasterSlaveServerGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponse) SetStatusCode(v int32) *DescribeMasterSlaveServerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponse) SetBody(v *DescribeMasterSlaveServerGroupsResponseBody) *DescribeMasterSlaveServerGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeMasterSlaveServerGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
