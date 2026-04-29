// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawChannelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarClawChannelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarClawChannelsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarClawChannelsResponseBody) *DescribePolarClawChannelsResponse
	GetBody() *DescribePolarClawChannelsResponseBody
}

type DescribePolarClawChannelsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarClawChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarClawChannelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawChannelsResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarClawChannelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarClawChannelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarClawChannelsResponse) GetBody() *DescribePolarClawChannelsResponseBody {
	return s.Body
}

func (s *DescribePolarClawChannelsResponse) SetHeaders(v map[string]*string) *DescribePolarClawChannelsResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarClawChannelsResponse) SetStatusCode(v int32) *DescribePolarClawChannelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarClawChannelsResponse) SetBody(v *DescribePolarClawChannelsResponseBody) *DescribePolarClawChannelsResponse {
	s.Body = v
	return s
}

func (s *DescribePolarClawChannelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
