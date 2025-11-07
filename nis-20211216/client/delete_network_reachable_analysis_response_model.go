// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkReachableAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNetworkReachableAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNetworkReachableAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNetworkReachableAnalysisResponseBody) *DeleteNetworkReachableAnalysisResponse
	GetBody() *DeleteNetworkReachableAnalysisResponseBody
}

type DeleteNetworkReachableAnalysisResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetworkReachableAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetworkReachableAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkReachableAnalysisResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkReachableAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNetworkReachableAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNetworkReachableAnalysisResponse) GetBody() *DeleteNetworkReachableAnalysisResponseBody {
	return s.Body
}

func (s *DeleteNetworkReachableAnalysisResponse) SetHeaders(v map[string]*string) *DeleteNetworkReachableAnalysisResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkReachableAnalysisResponse) SetStatusCode(v int32) *DeleteNetworkReachableAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetworkReachableAnalysisResponse) SetBody(v *DeleteNetworkReachableAnalysisResponseBody) *DeleteNetworkReachableAnalysisResponse {
	s.Body = v
	return s
}

func (s *DeleteNetworkReachableAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
