// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetImageBuildRiskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetImageBuildRiskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetImageBuildRiskStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetImageBuildRiskStatusResponseBody) *SetImageBuildRiskStatusResponse
	GetBody() *SetImageBuildRiskStatusResponseBody
}

type SetImageBuildRiskStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetImageBuildRiskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetImageBuildRiskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetImageBuildRiskStatusResponse) GoString() string {
	return s.String()
}

func (s *SetImageBuildRiskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetImageBuildRiskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetImageBuildRiskStatusResponse) GetBody() *SetImageBuildRiskStatusResponseBody {
	return s.Body
}

func (s *SetImageBuildRiskStatusResponse) SetHeaders(v map[string]*string) *SetImageBuildRiskStatusResponse {
	s.Headers = v
	return s
}

func (s *SetImageBuildRiskStatusResponse) SetStatusCode(v int32) *SetImageBuildRiskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetImageBuildRiskStatusResponse) SetBody(v *SetImageBuildRiskStatusResponseBody) *SetImageBuildRiskStatusResponse {
	s.Body = v
	return s
}

func (s *SetImageBuildRiskStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
