// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostCenterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCostCenterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCostCenterResponse
	GetStatusCode() *int32
	SetBody(v *QueryCostCenterResponseBody) *QueryCostCenterResponse
	GetBody() *QueryCostCenterResponseBody
}

type QueryCostCenterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCostCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCostCenterResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterResponse) GoString() string {
	return s.String()
}

func (s *QueryCostCenterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCostCenterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCostCenterResponse) GetBody() *QueryCostCenterResponseBody {
	return s.Body
}

func (s *QueryCostCenterResponse) SetHeaders(v map[string]*string) *QueryCostCenterResponse {
	s.Headers = v
	return s
}

func (s *QueryCostCenterResponse) SetStatusCode(v int32) *QueryCostCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCostCenterResponse) SetBody(v *QueryCostCenterResponseBody) *QueryCostCenterResponse {
	s.Body = v
	return s
}

func (s *QueryCostCenterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
