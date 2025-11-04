// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeMediaConnectFlowOutputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeMediaConnectFlowOutputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeMediaConnectFlowOutputResponse
	GetStatusCode() *int32
	SetBody(v *ResumeMediaConnectFlowOutputResponseBody) *ResumeMediaConnectFlowOutputResponse
	GetBody() *ResumeMediaConnectFlowOutputResponseBody
}

type ResumeMediaConnectFlowOutputResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeMediaConnectFlowOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeMediaConnectFlowOutputResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeMediaConnectFlowOutputResponse) GoString() string {
	return s.String()
}

func (s *ResumeMediaConnectFlowOutputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeMediaConnectFlowOutputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeMediaConnectFlowOutputResponse) GetBody() *ResumeMediaConnectFlowOutputResponseBody {
	return s.Body
}

func (s *ResumeMediaConnectFlowOutputResponse) SetHeaders(v map[string]*string) *ResumeMediaConnectFlowOutputResponse {
	s.Headers = v
	return s
}

func (s *ResumeMediaConnectFlowOutputResponse) SetStatusCode(v int32) *ResumeMediaConnectFlowOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeMediaConnectFlowOutputResponse) SetBody(v *ResumeMediaConnectFlowOutputResponseBody) *ResumeMediaConnectFlowOutputResponse {
	s.Body = v
	return s
}

func (s *ResumeMediaConnectFlowOutputResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
