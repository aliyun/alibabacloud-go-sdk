// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultOlapComputeGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDefaultOlapComputeGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDefaultOlapComputeGroupResponse
	GetStatusCode() *int32
	SetBody(v *SetDefaultOlapComputeGroupResponseBody) *SetDefaultOlapComputeGroupResponse
	GetBody() *SetDefaultOlapComputeGroupResponseBody
}

type SetDefaultOlapComputeGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDefaultOlapComputeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDefaultOlapComputeGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultOlapComputeGroupResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultOlapComputeGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDefaultOlapComputeGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDefaultOlapComputeGroupResponse) GetBody() *SetDefaultOlapComputeGroupResponseBody {
	return s.Body
}

func (s *SetDefaultOlapComputeGroupResponse) SetHeaders(v map[string]*string) *SetDefaultOlapComputeGroupResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultOlapComputeGroupResponse) SetStatusCode(v int32) *SetDefaultOlapComputeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDefaultOlapComputeGroupResponse) SetBody(v *SetDefaultOlapComputeGroupResponseBody) *SetDefaultOlapComputeGroupResponse {
	s.Body = v
	return s
}

func (s *SetDefaultOlapComputeGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
