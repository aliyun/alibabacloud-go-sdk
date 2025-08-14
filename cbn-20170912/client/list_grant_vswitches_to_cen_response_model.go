// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGrantVSwitchesToCenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGrantVSwitchesToCenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGrantVSwitchesToCenResponse
	GetStatusCode() *int32
	SetBody(v *ListGrantVSwitchesToCenResponseBody) *ListGrantVSwitchesToCenResponse
	GetBody() *ListGrantVSwitchesToCenResponseBody
}

type ListGrantVSwitchesToCenResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGrantVSwitchesToCenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGrantVSwitchesToCenResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGrantVSwitchesToCenResponse) GoString() string {
	return s.String()
}

func (s *ListGrantVSwitchesToCenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGrantVSwitchesToCenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGrantVSwitchesToCenResponse) GetBody() *ListGrantVSwitchesToCenResponseBody {
	return s.Body
}

func (s *ListGrantVSwitchesToCenResponse) SetHeaders(v map[string]*string) *ListGrantVSwitchesToCenResponse {
	s.Headers = v
	return s
}

func (s *ListGrantVSwitchesToCenResponse) SetStatusCode(v int32) *ListGrantVSwitchesToCenResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGrantVSwitchesToCenResponse) SetBody(v *ListGrantVSwitchesToCenResponseBody) *ListGrantVSwitchesToCenResponse {
	s.Body = v
	return s
}

func (s *ListGrantVSwitchesToCenResponse) Validate() error {
	return dara.Validate(s)
}
