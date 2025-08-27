// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGroupUserSaveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GroupUserSaveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GroupUserSaveResponse
	GetStatusCode() *int32
	SetBody(v *GroupUserSaveResponseBody) *GroupUserSaveResponse
	GetBody() *GroupUserSaveResponseBody
}

type GroupUserSaveResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GroupUserSaveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GroupUserSaveResponse) String() string {
	return dara.Prettify(s)
}

func (s GroupUserSaveResponse) GoString() string {
	return s.String()
}

func (s *GroupUserSaveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GroupUserSaveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GroupUserSaveResponse) GetBody() *GroupUserSaveResponseBody {
	return s.Body
}

func (s *GroupUserSaveResponse) SetHeaders(v map[string]*string) *GroupUserSaveResponse {
	s.Headers = v
	return s
}

func (s *GroupUserSaveResponse) SetStatusCode(v int32) *GroupUserSaveResponse {
	s.StatusCode = &v
	return s
}

func (s *GroupUserSaveResponse) SetBody(v *GroupUserSaveResponseBody) *GroupUserSaveResponse {
	s.Body = v
	return s
}

func (s *GroupUserSaveResponse) Validate() error {
	return dara.Validate(s)
}
