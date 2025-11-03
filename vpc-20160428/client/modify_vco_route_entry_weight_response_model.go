// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVcoRouteEntryWeightResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVcoRouteEntryWeightResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVcoRouteEntryWeightResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVcoRouteEntryWeightResponseBody) *ModifyVcoRouteEntryWeightResponse
	GetBody() *ModifyVcoRouteEntryWeightResponseBody
}

type ModifyVcoRouteEntryWeightResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVcoRouteEntryWeightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVcoRouteEntryWeightResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVcoRouteEntryWeightResponse) GoString() string {
	return s.String()
}

func (s *ModifyVcoRouteEntryWeightResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVcoRouteEntryWeightResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVcoRouteEntryWeightResponse) GetBody() *ModifyVcoRouteEntryWeightResponseBody {
	return s.Body
}

func (s *ModifyVcoRouteEntryWeightResponse) SetHeaders(v map[string]*string) *ModifyVcoRouteEntryWeightResponse {
	s.Headers = v
	return s
}

func (s *ModifyVcoRouteEntryWeightResponse) SetStatusCode(v int32) *ModifyVcoRouteEntryWeightResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVcoRouteEntryWeightResponse) SetBody(v *ModifyVcoRouteEntryWeightResponseBody) *ModifyVcoRouteEntryWeightResponse {
	s.Body = v
	return s
}

func (s *ModifyVcoRouteEntryWeightResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
