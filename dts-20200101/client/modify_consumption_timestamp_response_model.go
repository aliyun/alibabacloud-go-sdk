// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConsumptionTimestampResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyConsumptionTimestampResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyConsumptionTimestampResponse
	GetStatusCode() *int32
	SetBody(v *ModifyConsumptionTimestampResponseBody) *ModifyConsumptionTimestampResponse
	GetBody() *ModifyConsumptionTimestampResponseBody
}

type ModifyConsumptionTimestampResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyConsumptionTimestampResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyConsumptionTimestampResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyConsumptionTimestampResponse) GoString() string {
	return s.String()
}

func (s *ModifyConsumptionTimestampResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyConsumptionTimestampResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyConsumptionTimestampResponse) GetBody() *ModifyConsumptionTimestampResponseBody {
	return s.Body
}

func (s *ModifyConsumptionTimestampResponse) SetHeaders(v map[string]*string) *ModifyConsumptionTimestampResponse {
	s.Headers = v
	return s
}

func (s *ModifyConsumptionTimestampResponse) SetStatusCode(v int32) *ModifyConsumptionTimestampResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyConsumptionTimestampResponse) SetBody(v *ModifyConsumptionTimestampResponseBody) *ModifyConsumptionTimestampResponse {
	s.Body = v
	return s
}

func (s *ModifyConsumptionTimestampResponse) Validate() error {
	return dara.Validate(s)
}
