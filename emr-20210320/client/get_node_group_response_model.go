// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNodeGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNodeGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetNodeGroupResponseBody) *GetNodeGroupResponse
	GetBody() *GetNodeGroupResponseBody
}

type GetNodeGroupResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *GetNodeGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNodeGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNodeGroupResponse) GetBody() *GetNodeGroupResponseBody {
	return s.Body
}

func (s *GetNodeGroupResponse) SetHeaders(v map[string]*string) *GetNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *GetNodeGroupResponse) SetStatusCode(v int32) *GetNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeGroupResponse) SetBody(v *GetNodeGroupResponseBody) *GetNodeGroupResponse {
	s.Body = v
	return s
}

func (s *GetNodeGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
