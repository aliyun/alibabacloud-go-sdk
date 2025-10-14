// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsOperateLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePdnsOperateLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePdnsOperateLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePdnsOperateLogsResponseBody) *DescribePdnsOperateLogsResponse
	GetBody() *DescribePdnsOperateLogsResponseBody
}

type DescribePdnsOperateLogsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePdnsOperateLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePdnsOperateLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsOperateLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribePdnsOperateLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePdnsOperateLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePdnsOperateLogsResponse) GetBody() *DescribePdnsOperateLogsResponseBody {
	return s.Body
}

func (s *DescribePdnsOperateLogsResponse) SetHeaders(v map[string]*string) *DescribePdnsOperateLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribePdnsOperateLogsResponse) SetStatusCode(v int32) *DescribePdnsOperateLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePdnsOperateLogsResponse) SetBody(v *DescribePdnsOperateLogsResponseBody) *DescribePdnsOperateLogsResponse {
	s.Body = v
	return s
}

func (s *DescribePdnsOperateLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
