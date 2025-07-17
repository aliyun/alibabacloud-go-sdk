// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStandardGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStandardGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetStandardGroupResponseBody) *GetStandardGroupResponse
	GetBody() *GetStandardGroupResponseBody
}

type GetStandardGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStandardGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStandardGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStandardGroupResponse) GoString() string {
	return s.String()
}

func (s *GetStandardGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStandardGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStandardGroupResponse) GetBody() *GetStandardGroupResponseBody {
	return s.Body
}

func (s *GetStandardGroupResponse) SetHeaders(v map[string]*string) *GetStandardGroupResponse {
	s.Headers = v
	return s
}

func (s *GetStandardGroupResponse) SetStatusCode(v int32) *GetStandardGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStandardGroupResponse) SetBody(v *GetStandardGroupResponseBody) *GetStandardGroupResponse {
	s.Body = v
	return s
}

func (s *GetStandardGroupResponse) Validate() error {
	return dara.Validate(s)
}
