// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateCompliancePacksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAggregateCompliancePacksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAggregateCompliancePacksResponse
	GetStatusCode() *int32
	SetBody(v *ListAggregateCompliancePacksResponseBody) *ListAggregateCompliancePacksResponse
	GetBody() *ListAggregateCompliancePacksResponseBody
}

type ListAggregateCompliancePacksResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAggregateCompliancePacksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAggregateCompliancePacksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateCompliancePacksResponse) GoString() string {
	return s.String()
}

func (s *ListAggregateCompliancePacksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAggregateCompliancePacksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAggregateCompliancePacksResponse) GetBody() *ListAggregateCompliancePacksResponseBody {
	return s.Body
}

func (s *ListAggregateCompliancePacksResponse) SetHeaders(v map[string]*string) *ListAggregateCompliancePacksResponse {
	s.Headers = v
	return s
}

func (s *ListAggregateCompliancePacksResponse) SetStatusCode(v int32) *ListAggregateCompliancePacksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggregateCompliancePacksResponse) SetBody(v *ListAggregateCompliancePacksResponseBody) *ListAggregateCompliancePacksResponse {
	s.Body = v
	return s
}

func (s *ListAggregateCompliancePacksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
