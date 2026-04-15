// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertFlowDSLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConvertFlowDSLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConvertFlowDSLResponse
	GetStatusCode() *int32
	SetBody(v *ConvertFlowDSLResult) *ConvertFlowDSLResponse
	GetBody() *ConvertFlowDSLResult
}

type ConvertFlowDSLResponse struct {
	Headers    map[string]*string    `json:"headers" xml:"headers"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertFlowDSLResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertFlowDSLResponse) String() string {
	return dara.Prettify(s)
}

func (s ConvertFlowDSLResponse) GoString() string {
	return s.String()
}

func (s *ConvertFlowDSLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConvertFlowDSLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConvertFlowDSLResponse) GetBody() *ConvertFlowDSLResult {
	return s.Body
}

func (s *ConvertFlowDSLResponse) SetHeaders(v map[string]*string) *ConvertFlowDSLResponse {
	s.Headers = v
	return s
}

func (s *ConvertFlowDSLResponse) SetStatusCode(v int32) *ConvertFlowDSLResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertFlowDSLResponse) SetBody(v *ConvertFlowDSLResult) *ConvertFlowDSLResponse {
	s.Body = v
	return s
}

func (s *ConvertFlowDSLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
