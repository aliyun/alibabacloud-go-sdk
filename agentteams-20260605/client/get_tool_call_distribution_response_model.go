// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetToolCallDistributionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetToolCallDistributionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetToolCallDistributionResponse
	GetStatusCode() *int32
	SetBody(v *GetToolCallDistributionResponseBody) *GetToolCallDistributionResponse
	GetBody() *GetToolCallDistributionResponseBody
}

type GetToolCallDistributionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetToolCallDistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetToolCallDistributionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetToolCallDistributionResponse) GoString() string {
	return s.String()
}

func (s *GetToolCallDistributionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetToolCallDistributionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetToolCallDistributionResponse) GetBody() *GetToolCallDistributionResponseBody {
	return s.Body
}

func (s *GetToolCallDistributionResponse) SetHeaders(v map[string]*string) *GetToolCallDistributionResponse {
	s.Headers = v
	return s
}

func (s *GetToolCallDistributionResponse) SetStatusCode(v int32) *GetToolCallDistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetToolCallDistributionResponse) SetBody(v *GetToolCallDistributionResponseBody) *GetToolCallDistributionResponse {
	s.Body = v
	return s
}

func (s *GetToolCallDistributionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
