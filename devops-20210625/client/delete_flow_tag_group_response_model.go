// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowTagGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFlowTagGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFlowTagGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFlowTagGroupResponseBody) *DeleteFlowTagGroupResponse
	GetBody() *DeleteFlowTagGroupResponseBody
}

type DeleteFlowTagGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFlowTagGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFlowTagGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowTagGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowTagGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFlowTagGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFlowTagGroupResponse) GetBody() *DeleteFlowTagGroupResponseBody {
	return s.Body
}

func (s *DeleteFlowTagGroupResponse) SetHeaders(v map[string]*string) *DeleteFlowTagGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowTagGroupResponse) SetStatusCode(v int32) *DeleteFlowTagGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFlowTagGroupResponse) SetBody(v *DeleteFlowTagGroupResponseBody) *DeleteFlowTagGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteFlowTagGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
