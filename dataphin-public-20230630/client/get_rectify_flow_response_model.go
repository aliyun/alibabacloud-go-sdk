// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRectifyFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRectifyFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRectifyFlowResponse
	GetStatusCode() *int32
	SetBody(v *GetRectifyFlowResponseBody) *GetRectifyFlowResponse
	GetBody() *GetRectifyFlowResponseBody
}

type GetRectifyFlowResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRectifyFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRectifyFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRectifyFlowResponse) GoString() string {
	return s.String()
}

func (s *GetRectifyFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRectifyFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRectifyFlowResponse) GetBody() *GetRectifyFlowResponseBody {
	return s.Body
}

func (s *GetRectifyFlowResponse) SetHeaders(v map[string]*string) *GetRectifyFlowResponse {
	s.Headers = v
	return s
}

func (s *GetRectifyFlowResponse) SetStatusCode(v int32) *GetRectifyFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRectifyFlowResponse) SetBody(v *GetRectifyFlowResponseBody) *GetRectifyFlowResponse {
	s.Body = v
	return s
}

func (s *GetRectifyFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
