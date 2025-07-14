// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmartqAuthTransferResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SmartqAuthTransferResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SmartqAuthTransferResponse
	GetStatusCode() *int32
	SetBody(v *SmartqAuthTransferResponseBody) *SmartqAuthTransferResponse
	GetBody() *SmartqAuthTransferResponseBody
}

type SmartqAuthTransferResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SmartqAuthTransferResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SmartqAuthTransferResponse) String() string {
	return dara.Prettify(s)
}

func (s SmartqAuthTransferResponse) GoString() string {
	return s.String()
}

func (s *SmartqAuthTransferResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SmartqAuthTransferResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SmartqAuthTransferResponse) GetBody() *SmartqAuthTransferResponseBody {
	return s.Body
}

func (s *SmartqAuthTransferResponse) SetHeaders(v map[string]*string) *SmartqAuthTransferResponse {
	s.Headers = v
	return s
}

func (s *SmartqAuthTransferResponse) SetStatusCode(v int32) *SmartqAuthTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *SmartqAuthTransferResponse) SetBody(v *SmartqAuthTransferResponseBody) *SmartqAuthTransferResponse {
	s.Body = v
	return s
}

func (s *SmartqAuthTransferResponse) Validate() error {
	return dara.Validate(s)
}
