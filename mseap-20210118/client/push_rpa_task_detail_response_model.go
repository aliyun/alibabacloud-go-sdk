// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushRpaTaskDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushRpaTaskDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushRpaTaskDetailResponse
	GetStatusCode() *int32
	SetBody(v *PushRpaTaskDetailResponseBody) *PushRpaTaskDetailResponse
	GetBody() *PushRpaTaskDetailResponseBody
}

type PushRpaTaskDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushRpaTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushRpaTaskDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s PushRpaTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *PushRpaTaskDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushRpaTaskDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushRpaTaskDetailResponse) GetBody() *PushRpaTaskDetailResponseBody {
	return s.Body
}

func (s *PushRpaTaskDetailResponse) SetHeaders(v map[string]*string) *PushRpaTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *PushRpaTaskDetailResponse) SetStatusCode(v int32) *PushRpaTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *PushRpaTaskDetailResponse) SetBody(v *PushRpaTaskDetailResponseBody) *PushRpaTaskDetailResponse {
	s.Body = v
	return s
}

func (s *PushRpaTaskDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
