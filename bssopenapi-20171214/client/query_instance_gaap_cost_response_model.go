// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceGaapCostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryInstanceGaapCostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryInstanceGaapCostResponse
	GetStatusCode() *int32
	SetBody(v *QueryInstanceGaapCostResponseBody) *QueryInstanceGaapCostResponse
	GetBody() *QueryInstanceGaapCostResponseBody
}

type QueryInstanceGaapCostResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInstanceGaapCostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInstanceGaapCostResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceGaapCostResponse) GoString() string {
	return s.String()
}

func (s *QueryInstanceGaapCostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryInstanceGaapCostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryInstanceGaapCostResponse) GetBody() *QueryInstanceGaapCostResponseBody {
	return s.Body
}

func (s *QueryInstanceGaapCostResponse) SetHeaders(v map[string]*string) *QueryInstanceGaapCostResponse {
	s.Headers = v
	return s
}

func (s *QueryInstanceGaapCostResponse) SetStatusCode(v int32) *QueryInstanceGaapCostResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInstanceGaapCostResponse) SetBody(v *QueryInstanceGaapCostResponseBody) *QueryInstanceGaapCostResponse {
	s.Body = v
	return s
}

func (s *QueryInstanceGaapCostResponse) Validate() error {
	return dara.Validate(s)
}
