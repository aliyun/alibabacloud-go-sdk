// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeRouterWeightResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBNodeRouterWeightResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBNodeRouterWeightResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBNodeRouterWeightResponseBody) *ModifyDBNodeRouterWeightResponse
	GetBody() *ModifyDBNodeRouterWeightResponseBody
}

type ModifyDBNodeRouterWeightResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBNodeRouterWeightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBNodeRouterWeightResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeRouterWeightResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeRouterWeightResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBNodeRouterWeightResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBNodeRouterWeightResponse) GetBody() *ModifyDBNodeRouterWeightResponseBody {
	return s.Body
}

func (s *ModifyDBNodeRouterWeightResponse) SetHeaders(v map[string]*string) *ModifyDBNodeRouterWeightResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBNodeRouterWeightResponse) SetStatusCode(v int32) *ModifyDBNodeRouterWeightResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBNodeRouterWeightResponse) SetBody(v *ModifyDBNodeRouterWeightResponseBody) *ModifyDBNodeRouterWeightResponse {
	s.Body = v
	return s
}

func (s *ModifyDBNodeRouterWeightResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
