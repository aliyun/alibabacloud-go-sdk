// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateAvatarProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateAvatarProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateAvatarProjectResponse
	GetStatusCode() *int32
	SetBody(v *OperateAvatarProjectResponseBody) *OperateAvatarProjectResponse
	GetBody() *OperateAvatarProjectResponseBody
}

type OperateAvatarProjectResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateAvatarProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateAvatarProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateAvatarProjectResponse) GoString() string {
	return s.String()
}

func (s *OperateAvatarProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateAvatarProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateAvatarProjectResponse) GetBody() *OperateAvatarProjectResponseBody {
	return s.Body
}

func (s *OperateAvatarProjectResponse) SetHeaders(v map[string]*string) *OperateAvatarProjectResponse {
	s.Headers = v
	return s
}

func (s *OperateAvatarProjectResponse) SetStatusCode(v int32) *OperateAvatarProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateAvatarProjectResponse) SetBody(v *OperateAvatarProjectResponseBody) *OperateAvatarProjectResponse {
	s.Body = v
	return s
}

func (s *OperateAvatarProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
