// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMasterSlaveServerGroupAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMasterSlaveServerGroupAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMasterSlaveServerGroupAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMasterSlaveServerGroupAttributeResponseBody) *DescribeMasterSlaveServerGroupAttributeResponse
	GetBody() *DescribeMasterSlaveServerGroupAttributeResponseBody
}

type DescribeMasterSlaveServerGroupAttributeResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMasterSlaveServerGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMasterSlaveServerGroupAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMasterSlaveServerGroupAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMasterSlaveServerGroupAttributeResponse) GetBody() *DescribeMasterSlaveServerGroupAttributeResponseBody {
	return s.Body
}

func (s *DescribeMasterSlaveServerGroupAttributeResponse) SetHeaders(v map[string]*string) *DescribeMasterSlaveServerGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponse) SetStatusCode(v int32) *DescribeMasterSlaveServerGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponse) SetBody(v *DescribeMasterSlaveServerGroupAttributeResponseBody) *DescribeMasterSlaveServerGroupAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
