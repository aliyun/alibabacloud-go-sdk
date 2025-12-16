// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateABTestFixedFlowDividersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateABTestFixedFlowDividersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateABTestFixedFlowDividersResponse
	GetStatusCode() *int32
	SetBody(v *UpdateABTestFixedFlowDividersResponseBody) *UpdateABTestFixedFlowDividersResponse
	GetBody() *UpdateABTestFixedFlowDividersResponseBody
}

type UpdateABTestFixedFlowDividersResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateABTestFixedFlowDividersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateABTestFixedFlowDividersResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateABTestFixedFlowDividersResponse) GoString() string {
	return s.String()
}

func (s *UpdateABTestFixedFlowDividersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateABTestFixedFlowDividersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateABTestFixedFlowDividersResponse) GetBody() *UpdateABTestFixedFlowDividersResponseBody {
	return s.Body
}

func (s *UpdateABTestFixedFlowDividersResponse) SetHeaders(v map[string]*string) *UpdateABTestFixedFlowDividersResponse {
	s.Headers = v
	return s
}

func (s *UpdateABTestFixedFlowDividersResponse) SetStatusCode(v int32) *UpdateABTestFixedFlowDividersResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateABTestFixedFlowDividersResponse) SetBody(v *UpdateABTestFixedFlowDividersResponseBody) *UpdateABTestFixedFlowDividersResponse {
	s.Body = v
	return s
}

func (s *UpdateABTestFixedFlowDividersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
