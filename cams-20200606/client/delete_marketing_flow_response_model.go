// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMarketingFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMarketingFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMarketingFlowResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMarketingFlowResponseBody) *DeleteMarketingFlowResponse
	GetBody() *DeleteMarketingFlowResponseBody
}

type DeleteMarketingFlowResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMarketingFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMarketingFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMarketingFlowResponse) GoString() string {
	return s.String()
}

func (s *DeleteMarketingFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMarketingFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMarketingFlowResponse) GetBody() *DeleteMarketingFlowResponseBody {
	return s.Body
}

func (s *DeleteMarketingFlowResponse) SetHeaders(v map[string]*string) *DeleteMarketingFlowResponse {
	s.Headers = v
	return s
}

func (s *DeleteMarketingFlowResponse) SetStatusCode(v int32) *DeleteMarketingFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMarketingFlowResponse) SetBody(v *DeleteMarketingFlowResponseBody) *DeleteMarketingFlowResponse {
	s.Body = v
	return s
}

func (s *DeleteMarketingFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
