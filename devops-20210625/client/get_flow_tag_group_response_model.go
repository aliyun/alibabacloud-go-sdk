// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowTagGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFlowTagGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFlowTagGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetFlowTagGroupResponseBody) *GetFlowTagGroupResponse
	GetBody() *GetFlowTagGroupResponseBody
}

type GetFlowTagGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFlowTagGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFlowTagGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFlowTagGroupResponse) GoString() string {
	return s.String()
}

func (s *GetFlowTagGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFlowTagGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFlowTagGroupResponse) GetBody() *GetFlowTagGroupResponseBody {
	return s.Body
}

func (s *GetFlowTagGroupResponse) SetHeaders(v map[string]*string) *GetFlowTagGroupResponse {
	s.Headers = v
	return s
}

func (s *GetFlowTagGroupResponse) SetStatusCode(v int32) *GetFlowTagGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFlowTagGroupResponse) SetBody(v *GetFlowTagGroupResponseBody) *GetFlowTagGroupResponse {
	s.Body = v
	return s
}

func (s *GetFlowTagGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
