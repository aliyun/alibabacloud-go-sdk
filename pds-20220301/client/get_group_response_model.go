// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGroupResponse
	GetStatusCode() *int32
	SetBody(v *Group) *GetGroupResponse
	GetBody() *Group
}

type GetGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Group             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGroupResponse) GoString() string {
	return s.String()
}

func (s *GetGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGroupResponse) GetBody() *Group {
	return s.Body
}

func (s *GetGroupResponse) SetHeaders(v map[string]*string) *GetGroupResponse {
	s.Headers = v
	return s
}

func (s *GetGroupResponse) SetStatusCode(v int32) *GetGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupResponse) SetBody(v *Group) *GetGroupResponse {
	s.Body = v
	return s
}

func (s *GetGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
