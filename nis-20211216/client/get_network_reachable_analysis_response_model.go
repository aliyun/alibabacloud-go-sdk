// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkReachableAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNetworkReachableAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNetworkReachableAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *GetNetworkReachableAnalysisResponseBody) *GetNetworkReachableAnalysisResponse
	GetBody() *GetNetworkReachableAnalysisResponseBody
}

type GetNetworkReachableAnalysisResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNetworkReachableAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNetworkReachableAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkReachableAnalysisResponse) GoString() string {
	return s.String()
}

func (s *GetNetworkReachableAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNetworkReachableAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNetworkReachableAnalysisResponse) GetBody() *GetNetworkReachableAnalysisResponseBody {
	return s.Body
}

func (s *GetNetworkReachableAnalysisResponse) SetHeaders(v map[string]*string) *GetNetworkReachableAnalysisResponse {
	s.Headers = v
	return s
}

func (s *GetNetworkReachableAnalysisResponse) SetStatusCode(v int32) *GetNetworkReachableAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNetworkReachableAnalysisResponse) SetBody(v *GetNetworkReachableAnalysisResponseBody) *GetNetworkReachableAnalysisResponse {
	s.Body = v
	return s
}

func (s *GetNetworkReachableAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
