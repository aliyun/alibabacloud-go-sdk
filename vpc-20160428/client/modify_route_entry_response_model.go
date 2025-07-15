// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRouteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRouteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRouteEntryResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRouteEntryResponseBody) *ModifyRouteEntryResponse
	GetBody() *ModifyRouteEntryResponseBody
}

type ModifyRouteEntryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRouteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *ModifyRouteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRouteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRouteEntryResponse) GetBody() *ModifyRouteEntryResponseBody {
	return s.Body
}

func (s *ModifyRouteEntryResponse) SetHeaders(v map[string]*string) *ModifyRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *ModifyRouteEntryResponse) SetStatusCode(v int32) *ModifyRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRouteEntryResponse) SetBody(v *ModifyRouteEntryResponseBody) *ModifyRouteEntryResponse {
	s.Body = v
	return s
}

func (s *ModifyRouteEntryResponse) Validate() error {
	return dara.Validate(s)
}
