// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEnsRouteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyEnsRouteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyEnsRouteEntryResponse
	GetStatusCode() *int32
	SetBody(v *ModifyEnsRouteEntryResponseBody) *ModifyEnsRouteEntryResponse
	GetBody() *ModifyEnsRouteEntryResponseBody
}

type ModifyEnsRouteEntryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyEnsRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyEnsRouteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyEnsRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *ModifyEnsRouteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyEnsRouteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyEnsRouteEntryResponse) GetBody() *ModifyEnsRouteEntryResponseBody {
	return s.Body
}

func (s *ModifyEnsRouteEntryResponse) SetHeaders(v map[string]*string) *ModifyEnsRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *ModifyEnsRouteEntryResponse) SetStatusCode(v int32) *ModifyEnsRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEnsRouteEntryResponse) SetBody(v *ModifyEnsRouteEntryResponseBody) *ModifyEnsRouteEntryResponse {
	s.Body = v
	return s
}

func (s *ModifyEnsRouteEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
