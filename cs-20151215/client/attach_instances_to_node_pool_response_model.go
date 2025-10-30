// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachInstancesToNodePoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachInstancesToNodePoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachInstancesToNodePoolResponse
	GetStatusCode() *int32
	SetBody(v *AttachInstancesToNodePoolResponseBody) *AttachInstancesToNodePoolResponse
	GetBody() *AttachInstancesToNodePoolResponseBody
}

type AttachInstancesToNodePoolResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachInstancesToNodePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachInstancesToNodePoolResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachInstancesToNodePoolResponse) GoString() string {
	return s.String()
}

func (s *AttachInstancesToNodePoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachInstancesToNodePoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachInstancesToNodePoolResponse) GetBody() *AttachInstancesToNodePoolResponseBody {
	return s.Body
}

func (s *AttachInstancesToNodePoolResponse) SetHeaders(v map[string]*string) *AttachInstancesToNodePoolResponse {
	s.Headers = v
	return s
}

func (s *AttachInstancesToNodePoolResponse) SetStatusCode(v int32) *AttachInstancesToNodePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachInstancesToNodePoolResponse) SetBody(v *AttachInstancesToNodePoolResponseBody) *AttachInstancesToNodePoolResponse {
	s.Body = v
	return s
}

func (s *AttachInstancesToNodePoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
