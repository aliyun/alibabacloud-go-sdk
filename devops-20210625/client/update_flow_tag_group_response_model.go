// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowTagGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFlowTagGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFlowTagGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFlowTagGroupResponseBody) *UpdateFlowTagGroupResponse
	GetBody() *UpdateFlowTagGroupResponseBody
}

type UpdateFlowTagGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFlowTagGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFlowTagGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowTagGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlowTagGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFlowTagGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFlowTagGroupResponse) GetBody() *UpdateFlowTagGroupResponseBody {
	return s.Body
}

func (s *UpdateFlowTagGroupResponse) SetHeaders(v map[string]*string) *UpdateFlowTagGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlowTagGroupResponse) SetStatusCode(v int32) *UpdateFlowTagGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFlowTagGroupResponse) SetBody(v *UpdateFlowTagGroupResponseBody) *UpdateFlowTagGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateFlowTagGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
