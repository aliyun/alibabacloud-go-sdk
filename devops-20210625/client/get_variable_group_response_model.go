// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVariableGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVariableGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVariableGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetVariableGroupResponseBody) *GetVariableGroupResponse
	GetBody() *GetVariableGroupResponseBody
}

type GetVariableGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVariableGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVariableGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVariableGroupResponse) GoString() string {
	return s.String()
}

func (s *GetVariableGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVariableGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVariableGroupResponse) GetBody() *GetVariableGroupResponseBody {
	return s.Body
}

func (s *GetVariableGroupResponse) SetHeaders(v map[string]*string) *GetVariableGroupResponse {
	s.Headers = v
	return s
}

func (s *GetVariableGroupResponse) SetStatusCode(v int32) *GetVariableGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVariableGroupResponse) SetBody(v *GetVariableGroupResponseBody) *GetVariableGroupResponse {
	s.Body = v
	return s
}

func (s *GetVariableGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
