// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndAnalyzeNetworkPathResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAndAnalyzeNetworkPathResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAndAnalyzeNetworkPathResponse
	GetStatusCode() *int32
	SetBody(v *CreateAndAnalyzeNetworkPathResponseBody) *CreateAndAnalyzeNetworkPathResponse
	GetBody() *CreateAndAnalyzeNetworkPathResponseBody
}

type CreateAndAnalyzeNetworkPathResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAndAnalyzeNetworkPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAndAnalyzeNetworkPathResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAndAnalyzeNetworkPathResponse) GoString() string {
	return s.String()
}

func (s *CreateAndAnalyzeNetworkPathResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAndAnalyzeNetworkPathResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAndAnalyzeNetworkPathResponse) GetBody() *CreateAndAnalyzeNetworkPathResponseBody {
	return s.Body
}

func (s *CreateAndAnalyzeNetworkPathResponse) SetHeaders(v map[string]*string) *CreateAndAnalyzeNetworkPathResponse {
	s.Headers = v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponse) SetStatusCode(v int32) *CreateAndAnalyzeNetworkPathResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponse) SetBody(v *CreateAndAnalyzeNetworkPathResponseBody) *CreateAndAnalyzeNetworkPathResponse {
	s.Body = v
	return s
}

func (s *CreateAndAnalyzeNetworkPathResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
