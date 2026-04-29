// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawMCPServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarClawMCPServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarClawMCPServersResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarClawMCPServersResponseBody) *DescribePolarClawMCPServersResponse
	GetBody() *DescribePolarClawMCPServersResponseBody
}

type DescribePolarClawMCPServersResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarClawMCPServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarClawMCPServersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawMCPServersResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarClawMCPServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarClawMCPServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarClawMCPServersResponse) GetBody() *DescribePolarClawMCPServersResponseBody {
	return s.Body
}

func (s *DescribePolarClawMCPServersResponse) SetHeaders(v map[string]*string) *DescribePolarClawMCPServersResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarClawMCPServersResponse) SetStatusCode(v int32) *DescribePolarClawMCPServersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarClawMCPServersResponse) SetBody(v *DescribePolarClawMCPServersResponseBody) *DescribePolarClawMCPServersResponse {
	s.Body = v
	return s
}

func (s *DescribePolarClawMCPServersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
