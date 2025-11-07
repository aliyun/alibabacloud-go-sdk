// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkReachableAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNetworkReachableAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNetworkReachableAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *CreateNetworkReachableAnalysisResponseBody) *CreateNetworkReachableAnalysisResponse
	GetBody() *CreateNetworkReachableAnalysisResponseBody
}

type CreateNetworkReachableAnalysisResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetworkReachableAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetworkReachableAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkReachableAnalysisResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkReachableAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNetworkReachableAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNetworkReachableAnalysisResponse) GetBody() *CreateNetworkReachableAnalysisResponseBody {
	return s.Body
}

func (s *CreateNetworkReachableAnalysisResponse) SetHeaders(v map[string]*string) *CreateNetworkReachableAnalysisResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkReachableAnalysisResponse) SetStatusCode(v int32) *CreateNetworkReachableAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetworkReachableAnalysisResponse) SetBody(v *CreateNetworkReachableAnalysisResponseBody) *CreateNetworkReachableAnalysisResponse {
	s.Body = v
	return s
}

func (s *CreateNetworkReachableAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
