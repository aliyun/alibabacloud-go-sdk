// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserBatchQuitGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UserBatchQuitGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UserBatchQuitGroupResponse
	GetStatusCode() *int32
	SetBody(v *UserBatchQuitGroupResponseBody) *UserBatchQuitGroupResponse
	GetBody() *UserBatchQuitGroupResponseBody
}

type UserBatchQuitGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UserBatchQuitGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UserBatchQuitGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UserBatchQuitGroupResponse) GoString() string {
	return s.String()
}

func (s *UserBatchQuitGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UserBatchQuitGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UserBatchQuitGroupResponse) GetBody() *UserBatchQuitGroupResponseBody {
	return s.Body
}

func (s *UserBatchQuitGroupResponse) SetHeaders(v map[string]*string) *UserBatchQuitGroupResponse {
	s.Headers = v
	return s
}

func (s *UserBatchQuitGroupResponse) SetStatusCode(v int32) *UserBatchQuitGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UserBatchQuitGroupResponse) SetBody(v *UserBatchQuitGroupResponseBody) *UserBatchQuitGroupResponse {
	s.Body = v
	return s
}

func (s *UserBatchQuitGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
