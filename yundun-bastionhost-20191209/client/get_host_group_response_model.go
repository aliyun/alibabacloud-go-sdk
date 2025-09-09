// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHostGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHostGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHostGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetHostGroupResponseBody) *GetHostGroupResponse
	GetBody() *GetHostGroupResponseBody
}

type GetHostGroupResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHostGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHostGroupResponse) GoString() string {
	return s.String()
}

func (s *GetHostGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHostGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHostGroupResponse) GetBody() *GetHostGroupResponseBody {
	return s.Body
}

func (s *GetHostGroupResponse) SetHeaders(v map[string]*string) *GetHostGroupResponse {
	s.Headers = v
	return s
}

func (s *GetHostGroupResponse) SetStatusCode(v int32) *GetHostGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHostGroupResponse) SetBody(v *GetHostGroupResponseBody) *GetHostGroupResponse {
	s.Body = v
	return s
}

func (s *GetHostGroupResponse) Validate() error {
	return dara.Validate(s)
}
