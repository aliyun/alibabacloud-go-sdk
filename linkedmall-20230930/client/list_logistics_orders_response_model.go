// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogisticsOrdersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLogisticsOrdersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLogisticsOrdersResponse
	GetStatusCode() *int32
	SetBody(v *LogisticsOrderListResult) *ListLogisticsOrdersResponse
	GetBody() *LogisticsOrderListResult
}

type ListLogisticsOrdersResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LogisticsOrderListResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLogisticsOrdersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLogisticsOrdersResponse) GoString() string {
	return s.String()
}

func (s *ListLogisticsOrdersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLogisticsOrdersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLogisticsOrdersResponse) GetBody() *LogisticsOrderListResult {
	return s.Body
}

func (s *ListLogisticsOrdersResponse) SetHeaders(v map[string]*string) *ListLogisticsOrdersResponse {
	s.Headers = v
	return s
}

func (s *ListLogisticsOrdersResponse) SetStatusCode(v int32) *ListLogisticsOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogisticsOrdersResponse) SetBody(v *LogisticsOrderListResult) *ListLogisticsOrdersResponse {
	s.Body = v
	return s
}

func (s *ListLogisticsOrdersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
