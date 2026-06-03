// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSessionDistributionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSessionDistributionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSessionDistributionResponse
	GetStatusCode() *int32
	SetBody(v *GetSessionDistributionResponseBody) *GetSessionDistributionResponse
	GetBody() *GetSessionDistributionResponseBody
}

type GetSessionDistributionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSessionDistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSessionDistributionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSessionDistributionResponse) GoString() string {
	return s.String()
}

func (s *GetSessionDistributionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSessionDistributionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSessionDistributionResponse) GetBody() *GetSessionDistributionResponseBody {
	return s.Body
}

func (s *GetSessionDistributionResponse) SetHeaders(v map[string]*string) *GetSessionDistributionResponse {
	s.Headers = v
	return s
}

func (s *GetSessionDistributionResponse) SetStatusCode(v int32) *GetSessionDistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSessionDistributionResponse) SetBody(v *GetSessionDistributionResponseBody) *GetSessionDistributionResponse {
	s.Body = v
	return s
}

func (s *GetSessionDistributionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
