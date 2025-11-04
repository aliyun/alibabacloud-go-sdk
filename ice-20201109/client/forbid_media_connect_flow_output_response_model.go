// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForbidMediaConnectFlowOutputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ForbidMediaConnectFlowOutputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ForbidMediaConnectFlowOutputResponse
	GetStatusCode() *int32
	SetBody(v *ForbidMediaConnectFlowOutputResponseBody) *ForbidMediaConnectFlowOutputResponse
	GetBody() *ForbidMediaConnectFlowOutputResponseBody
}

type ForbidMediaConnectFlowOutputResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ForbidMediaConnectFlowOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ForbidMediaConnectFlowOutputResponse) String() string {
	return dara.Prettify(s)
}

func (s ForbidMediaConnectFlowOutputResponse) GoString() string {
	return s.String()
}

func (s *ForbidMediaConnectFlowOutputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ForbidMediaConnectFlowOutputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ForbidMediaConnectFlowOutputResponse) GetBody() *ForbidMediaConnectFlowOutputResponseBody {
	return s.Body
}

func (s *ForbidMediaConnectFlowOutputResponse) SetHeaders(v map[string]*string) *ForbidMediaConnectFlowOutputResponse {
	s.Headers = v
	return s
}

func (s *ForbidMediaConnectFlowOutputResponse) SetStatusCode(v int32) *ForbidMediaConnectFlowOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *ForbidMediaConnectFlowOutputResponse) SetBody(v *ForbidMediaConnectFlowOutputResponseBody) *ForbidMediaConnectFlowOutputResponse {
	s.Body = v
	return s
}

func (s *ForbidMediaConnectFlowOutputResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
