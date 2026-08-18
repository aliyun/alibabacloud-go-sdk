// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodePoolComponentInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNodePoolComponentInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNodePoolComponentInstancesResponse
	GetStatusCode() *int32
	SetBody(v *CreateNodePoolComponentInstancesResponseBody) *CreateNodePoolComponentInstancesResponse
	GetBody() *CreateNodePoolComponentInstancesResponseBody
}

type CreateNodePoolComponentInstancesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNodePoolComponentInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNodePoolComponentInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNodePoolComponentInstancesResponse) GoString() string {
	return s.String()
}

func (s *CreateNodePoolComponentInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNodePoolComponentInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNodePoolComponentInstancesResponse) GetBody() *CreateNodePoolComponentInstancesResponseBody {
	return s.Body
}

func (s *CreateNodePoolComponentInstancesResponse) SetHeaders(v map[string]*string) *CreateNodePoolComponentInstancesResponse {
	s.Headers = v
	return s
}

func (s *CreateNodePoolComponentInstancesResponse) SetStatusCode(v int32) *CreateNodePoolComponentInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNodePoolComponentInstancesResponse) SetBody(v *CreateNodePoolComponentInstancesResponseBody) *CreateNodePoolComponentInstancesResponse {
	s.Body = v
	return s
}

func (s *CreateNodePoolComponentInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
