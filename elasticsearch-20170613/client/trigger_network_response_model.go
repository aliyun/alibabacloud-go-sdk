// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TriggerNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TriggerNetworkResponse
	GetStatusCode() *int32
	SetBody(v *TriggerNetworkResponseBody) *TriggerNetworkResponse
	GetBody() *TriggerNetworkResponseBody
}

type TriggerNetworkResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s TriggerNetworkResponse) GoString() string {
	return s.String()
}

func (s *TriggerNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TriggerNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TriggerNetworkResponse) GetBody() *TriggerNetworkResponseBody {
	return s.Body
}

func (s *TriggerNetworkResponse) SetHeaders(v map[string]*string) *TriggerNetworkResponse {
	s.Headers = v
	return s
}

func (s *TriggerNetworkResponse) SetStatusCode(v int32) *TriggerNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerNetworkResponse) SetBody(v *TriggerNetworkResponseBody) *TriggerNetworkResponse {
	s.Body = v
	return s
}

func (s *TriggerNetworkResponse) Validate() error {
	return dara.Validate(s)
}
