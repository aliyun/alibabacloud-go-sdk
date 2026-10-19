// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseWorkItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PauseWorkItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PauseWorkItemResponse
	GetStatusCode() *int32
	SetBody(v *PauseWorkItemResponseBody) *PauseWorkItemResponse
	GetBody() *PauseWorkItemResponseBody
}

type PauseWorkItemResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PauseWorkItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PauseWorkItemResponse) String() string {
	return dara.Prettify(s)
}

func (s PauseWorkItemResponse) GoString() string {
	return s.String()
}

func (s *PauseWorkItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PauseWorkItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PauseWorkItemResponse) GetBody() *PauseWorkItemResponseBody {
	return s.Body
}

func (s *PauseWorkItemResponse) SetHeaders(v map[string]*string) *PauseWorkItemResponse {
	s.Headers = v
	return s
}

func (s *PauseWorkItemResponse) SetStatusCode(v int32) *PauseWorkItemResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseWorkItemResponse) SetBody(v *PauseWorkItemResponseBody) *PauseWorkItemResponse {
	s.Body = v
	return s
}

func (s *PauseWorkItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
