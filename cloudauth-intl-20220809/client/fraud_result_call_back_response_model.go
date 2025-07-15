// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFraudResultCallBackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FraudResultCallBackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FraudResultCallBackResponse
	GetStatusCode() *int32
	SetBody(v *FraudResultCallBackResponseBody) *FraudResultCallBackResponse
	GetBody() *FraudResultCallBackResponseBody
}

type FraudResultCallBackResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FraudResultCallBackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FraudResultCallBackResponse) String() string {
	return dara.Prettify(s)
}

func (s FraudResultCallBackResponse) GoString() string {
	return s.String()
}

func (s *FraudResultCallBackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FraudResultCallBackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FraudResultCallBackResponse) GetBody() *FraudResultCallBackResponseBody {
	return s.Body
}

func (s *FraudResultCallBackResponse) SetHeaders(v map[string]*string) *FraudResultCallBackResponse {
	s.Headers = v
	return s
}

func (s *FraudResultCallBackResponse) SetStatusCode(v int32) *FraudResultCallBackResponse {
	s.StatusCode = &v
	return s
}

func (s *FraudResultCallBackResponse) SetBody(v *FraudResultCallBackResponseBody) *FraudResultCallBackResponse {
	s.Body = v
	return s
}

func (s *FraudResultCallBackResponse) Validate() error {
	return dara.Validate(s)
}
