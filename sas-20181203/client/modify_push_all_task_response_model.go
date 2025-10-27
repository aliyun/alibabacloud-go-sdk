// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPushAllTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPushAllTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPushAllTaskResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPushAllTaskResponseBody) *ModifyPushAllTaskResponse
	GetBody() *ModifyPushAllTaskResponseBody
}

type ModifyPushAllTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPushAllTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPushAllTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPushAllTaskResponse) GoString() string {
	return s.String()
}

func (s *ModifyPushAllTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPushAllTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPushAllTaskResponse) GetBody() *ModifyPushAllTaskResponseBody {
	return s.Body
}

func (s *ModifyPushAllTaskResponse) SetHeaders(v map[string]*string) *ModifyPushAllTaskResponse {
	s.Headers = v
	return s
}

func (s *ModifyPushAllTaskResponse) SetStatusCode(v int32) *ModifyPushAllTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPushAllTaskResponse) SetBody(v *ModifyPushAllTaskResponseBody) *ModifyPushAllTaskResponse {
	s.Body = v
	return s
}

func (s *ModifyPushAllTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
