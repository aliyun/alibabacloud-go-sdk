// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListABTestFixedFlowDividersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListABTestFixedFlowDividersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListABTestFixedFlowDividersResponse
	GetStatusCode() *int32
	SetBody(v *ListABTestFixedFlowDividersResponseBody) *ListABTestFixedFlowDividersResponse
	GetBody() *ListABTestFixedFlowDividersResponseBody
}

type ListABTestFixedFlowDividersResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListABTestFixedFlowDividersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListABTestFixedFlowDividersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListABTestFixedFlowDividersResponse) GoString() string {
	return s.String()
}

func (s *ListABTestFixedFlowDividersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListABTestFixedFlowDividersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListABTestFixedFlowDividersResponse) GetBody() *ListABTestFixedFlowDividersResponseBody {
	return s.Body
}

func (s *ListABTestFixedFlowDividersResponse) SetHeaders(v map[string]*string) *ListABTestFixedFlowDividersResponse {
	s.Headers = v
	return s
}

func (s *ListABTestFixedFlowDividersResponse) SetStatusCode(v int32) *ListABTestFixedFlowDividersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListABTestFixedFlowDividersResponse) SetBody(v *ListABTestFixedFlowDividersResponseBody) *ListABTestFixedFlowDividersResponse {
	s.Body = v
	return s
}

func (s *ListABTestFixedFlowDividersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
