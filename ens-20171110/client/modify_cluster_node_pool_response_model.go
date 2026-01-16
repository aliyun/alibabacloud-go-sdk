// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterNodePoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClusterNodePoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClusterNodePoolResponse
	GetStatusCode() *int32
	SetBody(v *ModifyClusterNodePoolResponseBody) *ModifyClusterNodePoolResponse
	GetBody() *ModifyClusterNodePoolResponseBody
}

type ModifyClusterNodePoolResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClusterNodePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClusterNodePoolResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNodePoolResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterNodePoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClusterNodePoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClusterNodePoolResponse) GetBody() *ModifyClusterNodePoolResponseBody {
	return s.Body
}

func (s *ModifyClusterNodePoolResponse) SetHeaders(v map[string]*string) *ModifyClusterNodePoolResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterNodePoolResponse) SetStatusCode(v int32) *ModifyClusterNodePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClusterNodePoolResponse) SetBody(v *ModifyClusterNodePoolResponseBody) *ModifyClusterNodePoolResponse {
	s.Body = v
	return s
}

func (s *ModifyClusterNodePoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
