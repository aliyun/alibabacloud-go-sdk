// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStackGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStackGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetStackGroupResponseBody) *GetStackGroupResponse
	GetBody() *GetStackGroupResponseBody
}

type GetStackGroupResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStackGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStackGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStackGroupResponse) GoString() string {
	return s.String()
}

func (s *GetStackGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStackGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStackGroupResponse) GetBody() *GetStackGroupResponseBody {
	return s.Body
}

func (s *GetStackGroupResponse) SetHeaders(v map[string]*string) *GetStackGroupResponse {
	s.Headers = v
	return s
}

func (s *GetStackGroupResponse) SetStatusCode(v int32) *GetStackGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStackGroupResponse) SetBody(v *GetStackGroupResponseBody) *GetStackGroupResponse {
	s.Body = v
	return s
}

func (s *GetStackGroupResponse) Validate() error {
	return dara.Validate(s)
}
