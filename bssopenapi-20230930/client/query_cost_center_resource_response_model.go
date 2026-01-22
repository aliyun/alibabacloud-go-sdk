// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostCenterResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCostCenterResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCostCenterResourceResponse
	GetStatusCode() *int32
	SetBody(v *QueryCostCenterResourceResponseBody) *QueryCostCenterResourceResponse
	GetBody() *QueryCostCenterResourceResponseBody
}

type QueryCostCenterResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCostCenterResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCostCenterResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterResourceResponse) GoString() string {
	return s.String()
}

func (s *QueryCostCenterResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCostCenterResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCostCenterResourceResponse) GetBody() *QueryCostCenterResourceResponseBody {
	return s.Body
}

func (s *QueryCostCenterResourceResponse) SetHeaders(v map[string]*string) *QueryCostCenterResourceResponse {
	s.Headers = v
	return s
}

func (s *QueryCostCenterResourceResponse) SetStatusCode(v int32) *QueryCostCenterResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCostCenterResourceResponse) SetBody(v *QueryCostCenterResourceResponseBody) *QueryCostCenterResourceResponse {
	s.Body = v
	return s
}

func (s *QueryCostCenterResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
