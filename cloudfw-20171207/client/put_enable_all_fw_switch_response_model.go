// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutEnableAllFwSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutEnableAllFwSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutEnableAllFwSwitchResponse
	GetStatusCode() *int32
	SetBody(v *PutEnableAllFwSwitchResponseBody) *PutEnableAllFwSwitchResponse
	GetBody() *PutEnableAllFwSwitchResponseBody
}

type PutEnableAllFwSwitchResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutEnableAllFwSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutEnableAllFwSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s PutEnableAllFwSwitchResponse) GoString() string {
	return s.String()
}

func (s *PutEnableAllFwSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutEnableAllFwSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutEnableAllFwSwitchResponse) GetBody() *PutEnableAllFwSwitchResponseBody {
	return s.Body
}

func (s *PutEnableAllFwSwitchResponse) SetHeaders(v map[string]*string) *PutEnableAllFwSwitchResponse {
	s.Headers = v
	return s
}

func (s *PutEnableAllFwSwitchResponse) SetStatusCode(v int32) *PutEnableAllFwSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *PutEnableAllFwSwitchResponse) SetBody(v *PutEnableAllFwSwitchResponseBody) *PutEnableAllFwSwitchResponse {
	s.Body = v
	return s
}

func (s *PutEnableAllFwSwitchResponse) Validate() error {
	return dara.Validate(s)
}
