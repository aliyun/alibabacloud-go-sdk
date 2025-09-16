// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAdvanceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAdvanceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAdvanceConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAdvanceConfigResponseBody) *ModifyAdvanceConfigResponse
	GetBody() *ModifyAdvanceConfigResponseBody
}

type ModifyAdvanceConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAdvanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAdvanceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAdvanceConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyAdvanceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAdvanceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAdvanceConfigResponse) GetBody() *ModifyAdvanceConfigResponseBody {
	return s.Body
}

func (s *ModifyAdvanceConfigResponse) SetHeaders(v map[string]*string) *ModifyAdvanceConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyAdvanceConfigResponse) SetStatusCode(v int32) *ModifyAdvanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAdvanceConfigResponse) SetBody(v *ModifyAdvanceConfigResponseBody) *ModifyAdvanceConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyAdvanceConfigResponse) Validate() error {
	return dara.Validate(s)
}
