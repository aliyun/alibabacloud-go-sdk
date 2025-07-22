// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCostCenterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCostCenterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCostCenterResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCostCenterResponseBody) *ModifyCostCenterResponse
	GetBody() *ModifyCostCenterResponseBody
}

type ModifyCostCenterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCostCenterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCostCenterResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostCenterResponse) GoString() string {
	return s.String()
}

func (s *ModifyCostCenterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCostCenterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCostCenterResponse) GetBody() *ModifyCostCenterResponseBody {
	return s.Body
}

func (s *ModifyCostCenterResponse) SetHeaders(v map[string]*string) *ModifyCostCenterResponse {
	s.Headers = v
	return s
}

func (s *ModifyCostCenterResponse) SetStatusCode(v int32) *ModifyCostCenterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCostCenterResponse) SetBody(v *ModifyCostCenterResponseBody) *ModifyCostCenterResponse {
	s.Body = v
	return s
}

func (s *ModifyCostCenterResponse) Validate() error {
	return dara.Validate(s)
}
