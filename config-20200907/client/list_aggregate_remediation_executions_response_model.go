// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateRemediationExecutionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAggregateRemediationExecutionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAggregateRemediationExecutionsResponse
	GetStatusCode() *int32
	SetBody(v *ListAggregateRemediationExecutionsResponseBody) *ListAggregateRemediationExecutionsResponse
	GetBody() *ListAggregateRemediationExecutionsResponseBody
}

type ListAggregateRemediationExecutionsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAggregateRemediationExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAggregateRemediationExecutionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateRemediationExecutionsResponse) GoString() string {
	return s.String()
}

func (s *ListAggregateRemediationExecutionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAggregateRemediationExecutionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAggregateRemediationExecutionsResponse) GetBody() *ListAggregateRemediationExecutionsResponseBody {
	return s.Body
}

func (s *ListAggregateRemediationExecutionsResponse) SetHeaders(v map[string]*string) *ListAggregateRemediationExecutionsResponse {
	s.Headers = v
	return s
}

func (s *ListAggregateRemediationExecutionsResponse) SetStatusCode(v int32) *ListAggregateRemediationExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggregateRemediationExecutionsResponse) SetBody(v *ListAggregateRemediationExecutionsResponseBody) *ListAggregateRemediationExecutionsResponse {
	s.Body = v
	return s
}

func (s *ListAggregateRemediationExecutionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
