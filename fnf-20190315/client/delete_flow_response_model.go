// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFlowResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFlowResponseBody) *DeleteFlowResponse
	GetBody() *DeleteFlowResponseBody
}

type DeleteFlowResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFlowResponse) GetBody() *DeleteFlowResponseBody {
	return s.Body
}

func (s *DeleteFlowResponse) SetHeaders(v map[string]*string) *DeleteFlowResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowResponse) SetStatusCode(v int32) *DeleteFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFlowResponse) SetBody(v *DeleteFlowResponseBody) *DeleteFlowResponse {
	s.Body = v
	return s
}

func (s *DeleteFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
