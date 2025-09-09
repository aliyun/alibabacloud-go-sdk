// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBroadcastTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBroadcastTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBroadcastTablesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBroadcastTablesResponseBody) *DescribeBroadcastTablesResponse
	GetBody() *DescribeBroadcastTablesResponseBody
}

type DescribeBroadcastTablesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBroadcastTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBroadcastTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBroadcastTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBroadcastTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBroadcastTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBroadcastTablesResponse) GetBody() *DescribeBroadcastTablesResponseBody {
	return s.Body
}

func (s *DescribeBroadcastTablesResponse) SetHeaders(v map[string]*string) *DescribeBroadcastTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeBroadcastTablesResponse) SetStatusCode(v int32) *DescribeBroadcastTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBroadcastTablesResponse) SetBody(v *DescribeBroadcastTablesResponseBody) *DescribeBroadcastTablesResponse {
	s.Body = v
	return s
}

func (s *DescribeBroadcastTablesResponse) Validate() error {
	return dara.Validate(s)
}
