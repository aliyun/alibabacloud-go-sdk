// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeprecateFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeprecateFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeprecateFlowResponse
	GetStatusCode() *int32
	SetBody(v *DeprecateFlowResponseBody) *DeprecateFlowResponse
	GetBody() *DeprecateFlowResponseBody
}

type DeprecateFlowResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeprecateFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeprecateFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s DeprecateFlowResponse) GoString() string {
	return s.String()
}

func (s *DeprecateFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeprecateFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeprecateFlowResponse) GetBody() *DeprecateFlowResponseBody {
	return s.Body
}

func (s *DeprecateFlowResponse) SetHeaders(v map[string]*string) *DeprecateFlowResponse {
	s.Headers = v
	return s
}

func (s *DeprecateFlowResponse) SetStatusCode(v int32) *DeprecateFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeprecateFlowResponse) SetBody(v *DeprecateFlowResponseBody) *DeprecateFlowResponse {
	s.Body = v
	return s
}

func (s *DeprecateFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
