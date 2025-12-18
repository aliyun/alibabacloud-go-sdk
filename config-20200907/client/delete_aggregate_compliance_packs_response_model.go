// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAggregateCompliancePacksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAggregateCompliancePacksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAggregateCompliancePacksResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAggregateCompliancePacksResponseBody) *DeleteAggregateCompliancePacksResponse
	GetBody() *DeleteAggregateCompliancePacksResponseBody
}

type DeleteAggregateCompliancePacksResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAggregateCompliancePacksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAggregateCompliancePacksResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregateCompliancePacksResponse) GoString() string {
	return s.String()
}

func (s *DeleteAggregateCompliancePacksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAggregateCompliancePacksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAggregateCompliancePacksResponse) GetBody() *DeleteAggregateCompliancePacksResponseBody {
	return s.Body
}

func (s *DeleteAggregateCompliancePacksResponse) SetHeaders(v map[string]*string) *DeleteAggregateCompliancePacksResponse {
	s.Headers = v
	return s
}

func (s *DeleteAggregateCompliancePacksResponse) SetStatusCode(v int32) *DeleteAggregateCompliancePacksResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAggregateCompliancePacksResponse) SetBody(v *DeleteAggregateCompliancePacksResponseBody) *DeleteAggregateCompliancePacksResponse {
	s.Body = v
	return s
}

func (s *DeleteAggregateCompliancePacksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
