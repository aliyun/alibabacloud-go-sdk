// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogTypeDistributionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLogTypeDistributionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLogTypeDistributionResponse
	GetStatusCode() *int32
	SetBody(v *GetLogTypeDistributionResponseBody) *GetLogTypeDistributionResponse
	GetBody() *GetLogTypeDistributionResponseBody
}

type GetLogTypeDistributionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLogTypeDistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLogTypeDistributionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLogTypeDistributionResponse) GoString() string {
	return s.String()
}

func (s *GetLogTypeDistributionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLogTypeDistributionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLogTypeDistributionResponse) GetBody() *GetLogTypeDistributionResponseBody {
	return s.Body
}

func (s *GetLogTypeDistributionResponse) SetHeaders(v map[string]*string) *GetLogTypeDistributionResponse {
	s.Headers = v
	return s
}

func (s *GetLogTypeDistributionResponse) SetStatusCode(v int32) *GetLogTypeDistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogTypeDistributionResponse) SetBody(v *GetLogTypeDistributionResponseBody) *GetLogTypeDistributionResponse {
	s.Body = v
	return s
}

func (s *GetLogTypeDistributionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
