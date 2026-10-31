// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRectifyFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRectifyFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRectifyFlowResponse
	GetStatusCode() *int32
	SetBody(v *ListRectifyFlowResponseBody) *ListRectifyFlowResponse
	GetBody() *ListRectifyFlowResponseBody
}

type ListRectifyFlowResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRectifyFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRectifyFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRectifyFlowResponse) GoString() string {
	return s.String()
}

func (s *ListRectifyFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRectifyFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRectifyFlowResponse) GetBody() *ListRectifyFlowResponseBody {
	return s.Body
}

func (s *ListRectifyFlowResponse) SetHeaders(v map[string]*string) *ListRectifyFlowResponse {
	s.Headers = v
	return s
}

func (s *ListRectifyFlowResponse) SetStatusCode(v int32) *ListRectifyFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRectifyFlowResponse) SetBody(v *ListRectifyFlowResponseBody) *ListRectifyFlowResponse {
	s.Body = v
	return s
}

func (s *ListRectifyFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
