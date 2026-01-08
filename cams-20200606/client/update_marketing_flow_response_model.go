// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMarketingFLowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMarketingFLowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMarketingFLowResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMarketingFLowResponseBody) *UpdateMarketingFLowResponse
	GetBody() *UpdateMarketingFLowResponseBody
}

type UpdateMarketingFLowResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMarketingFLowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMarketingFLowResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMarketingFLowResponse) GoString() string {
	return s.String()
}

func (s *UpdateMarketingFLowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMarketingFLowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMarketingFLowResponse) GetBody() *UpdateMarketingFLowResponseBody {
	return s.Body
}

func (s *UpdateMarketingFLowResponse) SetHeaders(v map[string]*string) *UpdateMarketingFLowResponse {
	s.Headers = v
	return s
}

func (s *UpdateMarketingFLowResponse) SetStatusCode(v int32) *UpdateMarketingFLowResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMarketingFLowResponse) SetBody(v *UpdateMarketingFLowResponseBody) *UpdateMarketingFLowResponse {
	s.Body = v
	return s
}

func (s *UpdateMarketingFLowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
