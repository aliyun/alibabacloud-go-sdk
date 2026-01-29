// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCostUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCostUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCostUnitResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCostUnitResponseBody) *ModifyCostUnitResponse
	GetBody() *ModifyCostUnitResponseBody
}

type ModifyCostUnitResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCostUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCostUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostUnitResponse) GoString() string {
	return s.String()
}

func (s *ModifyCostUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCostUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCostUnitResponse) GetBody() *ModifyCostUnitResponseBody {
	return s.Body
}

func (s *ModifyCostUnitResponse) SetHeaders(v map[string]*string) *ModifyCostUnitResponse {
	s.Headers = v
	return s
}

func (s *ModifyCostUnitResponse) SetStatusCode(v int32) *ModifyCostUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCostUnitResponse) SetBody(v *ModifyCostUnitResponseBody) *ModifyCostUnitResponse {
	s.Body = v
	return s
}

func (s *ModifyCostUnitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
