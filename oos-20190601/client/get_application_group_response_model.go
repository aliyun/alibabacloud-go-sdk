// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApplicationGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApplicationGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetApplicationGroupResponseBody) *GetApplicationGroupResponse
	GetBody() *GetApplicationGroupResponseBody
}

type GetApplicationGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationGroupResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApplicationGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApplicationGroupResponse) GetBody() *GetApplicationGroupResponseBody {
	return s.Body
}

func (s *GetApplicationGroupResponse) SetHeaders(v map[string]*string) *GetApplicationGroupResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationGroupResponse) SetStatusCode(v int32) *GetApplicationGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationGroupResponse) SetBody(v *GetApplicationGroupResponseBody) *GetApplicationGroupResponse {
	s.Body = v
	return s
}

func (s *GetApplicationGroupResponse) Validate() error {
	return dara.Validate(s)
}
