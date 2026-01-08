// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseMarketingFLowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PauseMarketingFLowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PauseMarketingFLowResponse
	GetStatusCode() *int32
	SetBody(v *PauseMarketingFLowResponseBody) *PauseMarketingFLowResponse
	GetBody() *PauseMarketingFLowResponseBody
}

type PauseMarketingFLowResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PauseMarketingFLowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PauseMarketingFLowResponse) String() string {
	return dara.Prettify(s)
}

func (s PauseMarketingFLowResponse) GoString() string {
	return s.String()
}

func (s *PauseMarketingFLowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PauseMarketingFLowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PauseMarketingFLowResponse) GetBody() *PauseMarketingFLowResponseBody {
	return s.Body
}

func (s *PauseMarketingFLowResponse) SetHeaders(v map[string]*string) *PauseMarketingFLowResponse {
	s.Headers = v
	return s
}

func (s *PauseMarketingFLowResponse) SetStatusCode(v int32) *PauseMarketingFLowResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseMarketingFLowResponse) SetBody(v *PauseMarketingFLowResponseBody) *PauseMarketingFLowResponse {
	s.Body = v
	return s
}

func (s *PauseMarketingFLowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
