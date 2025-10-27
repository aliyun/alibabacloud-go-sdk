// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFlowResponse
	GetStatusCode() *int32
	SetBody(v *CreateFlowResponseBody) *CreateFlowResponse
	GetBody() *CreateFlowResponseBody
}

type CreateFlowResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFlowResponse) GetBody() *CreateFlowResponseBody {
	return s.Body
}

func (s *CreateFlowResponse) SetHeaders(v map[string]*string) *CreateFlowResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowResponse) SetStatusCode(v int32) *CreateFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFlowResponse) SetBody(v *CreateFlowResponseBody) *CreateFlowResponse {
	s.Body = v
	return s
}

func (s *CreateFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
