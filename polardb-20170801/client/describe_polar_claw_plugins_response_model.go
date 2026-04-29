// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawPluginsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarClawPluginsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarClawPluginsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarClawPluginsResponseBody) *DescribePolarClawPluginsResponse
	GetBody() *DescribePolarClawPluginsResponseBody
}

type DescribePolarClawPluginsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarClawPluginsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarClawPluginsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawPluginsResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarClawPluginsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarClawPluginsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarClawPluginsResponse) GetBody() *DescribePolarClawPluginsResponseBody {
	return s.Body
}

func (s *DescribePolarClawPluginsResponse) SetHeaders(v map[string]*string) *DescribePolarClawPluginsResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarClawPluginsResponse) SetStatusCode(v int32) *DescribePolarClawPluginsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarClawPluginsResponse) SetBody(v *DescribePolarClawPluginsResponseBody) *DescribePolarClawPluginsResponse {
	s.Body = v
	return s
}

func (s *DescribePolarClawPluginsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
