// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunContractExtractResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunContractExtractResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunContractExtractResponse
	GetStatusCode() *int32
	SetBody(v *RunContractExtractResponseBody) *RunContractExtractResponse
	GetBody() *RunContractExtractResponseBody
}

type RunContractExtractResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunContractExtractResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunContractExtractResponse) String() string {
	return dara.Prettify(s)
}

func (s RunContractExtractResponse) GoString() string {
	return s.String()
}

func (s *RunContractExtractResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunContractExtractResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunContractExtractResponse) GetBody() *RunContractExtractResponseBody {
	return s.Body
}

func (s *RunContractExtractResponse) SetHeaders(v map[string]*string) *RunContractExtractResponse {
	s.Headers = v
	return s
}

func (s *RunContractExtractResponse) SetStatusCode(v int32) *RunContractExtractResponse {
	s.StatusCode = &v
	return s
}

func (s *RunContractExtractResponse) SetBody(v *RunContractExtractResponseBody) *RunContractExtractResponse {
	s.Body = v
	return s
}

func (s *RunContractExtractResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
