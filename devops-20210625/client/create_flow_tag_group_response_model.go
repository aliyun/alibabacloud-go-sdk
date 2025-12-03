// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowTagGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFlowTagGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFlowTagGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateFlowTagGroupResponseBody) *CreateFlowTagGroupResponse
	GetBody() *CreateFlowTagGroupResponseBody
}

type CreateFlowTagGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFlowTagGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFlowTagGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowTagGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowTagGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFlowTagGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFlowTagGroupResponse) GetBody() *CreateFlowTagGroupResponseBody {
	return s.Body
}

func (s *CreateFlowTagGroupResponse) SetHeaders(v map[string]*string) *CreateFlowTagGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowTagGroupResponse) SetStatusCode(v int32) *CreateFlowTagGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFlowTagGroupResponse) SetBody(v *CreateFlowTagGroupResponseBody) *CreateFlowTagGroupResponse {
	s.Body = v
	return s
}

func (s *CreateFlowTagGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
