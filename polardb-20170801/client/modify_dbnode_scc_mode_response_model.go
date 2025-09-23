// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeSccModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBNodeSccModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBNodeSccModeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBNodeSccModeResponseBody) *ModifyDBNodeSccModeResponse
	GetBody() *ModifyDBNodeSccModeResponseBody
}

type ModifyDBNodeSccModeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBNodeSccModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBNodeSccModeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeSccModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeSccModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBNodeSccModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBNodeSccModeResponse) GetBody() *ModifyDBNodeSccModeResponseBody {
	return s.Body
}

func (s *ModifyDBNodeSccModeResponse) SetHeaders(v map[string]*string) *ModifyDBNodeSccModeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBNodeSccModeResponse) SetStatusCode(v int32) *ModifyDBNodeSccModeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBNodeSccModeResponse) SetBody(v *ModifyDBNodeSccModeResponseBody) *ModifyDBNodeSccModeResponse {
	s.Body = v
	return s
}

func (s *ModifyDBNodeSccModeResponse) Validate() error {
	return dara.Validate(s)
}
