// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLdpsComputeGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLdpsComputeGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLdpsComputeGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateLdpsComputeGroupResponseBody) *CreateLdpsComputeGroupResponse
	GetBody() *CreateLdpsComputeGroupResponseBody
}

type CreateLdpsComputeGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLdpsComputeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLdpsComputeGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLdpsComputeGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateLdpsComputeGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLdpsComputeGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLdpsComputeGroupResponse) GetBody() *CreateLdpsComputeGroupResponseBody {
	return s.Body
}

func (s *CreateLdpsComputeGroupResponse) SetHeaders(v map[string]*string) *CreateLdpsComputeGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateLdpsComputeGroupResponse) SetStatusCode(v int32) *CreateLdpsComputeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLdpsComputeGroupResponse) SetBody(v *CreateLdpsComputeGroupResponseBody) *CreateLdpsComputeGroupResponse {
	s.Body = v
	return s
}

func (s *CreateLdpsComputeGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
