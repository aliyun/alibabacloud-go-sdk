// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDisableAllFwSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutDisableAllFwSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutDisableAllFwSwitchResponse
	GetStatusCode() *int32
	SetBody(v *PutDisableAllFwSwitchResponseBody) *PutDisableAllFwSwitchResponse
	GetBody() *PutDisableAllFwSwitchResponseBody
}

type PutDisableAllFwSwitchResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutDisableAllFwSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutDisableAllFwSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s PutDisableAllFwSwitchResponse) GoString() string {
	return s.String()
}

func (s *PutDisableAllFwSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutDisableAllFwSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutDisableAllFwSwitchResponse) GetBody() *PutDisableAllFwSwitchResponseBody {
	return s.Body
}

func (s *PutDisableAllFwSwitchResponse) SetHeaders(v map[string]*string) *PutDisableAllFwSwitchResponse {
	s.Headers = v
	return s
}

func (s *PutDisableAllFwSwitchResponse) SetStatusCode(v int32) *PutDisableAllFwSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *PutDisableAllFwSwitchResponse) SetBody(v *PutDisableAllFwSwitchResponseBody) *PutDisableAllFwSwitchResponse {
	s.Body = v
	return s
}

func (s *PutDisableAllFwSwitchResponse) Validate() error {
	return dara.Validate(s)
}
