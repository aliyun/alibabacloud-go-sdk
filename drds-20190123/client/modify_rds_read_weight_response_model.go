// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRdsReadWeightResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRdsReadWeightResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRdsReadWeightResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRdsReadWeightResponseBody) *ModifyRdsReadWeightResponse
	GetBody() *ModifyRdsReadWeightResponseBody
}

type ModifyRdsReadWeightResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRdsReadWeightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRdsReadWeightResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRdsReadWeightResponse) GoString() string {
	return s.String()
}

func (s *ModifyRdsReadWeightResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRdsReadWeightResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRdsReadWeightResponse) GetBody() *ModifyRdsReadWeightResponseBody {
	return s.Body
}

func (s *ModifyRdsReadWeightResponse) SetHeaders(v map[string]*string) *ModifyRdsReadWeightResponse {
	s.Headers = v
	return s
}

func (s *ModifyRdsReadWeightResponse) SetStatusCode(v int32) *ModifyRdsReadWeightResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRdsReadWeightResponse) SetBody(v *ModifyRdsReadWeightResponseBody) *ModifyRdsReadWeightResponse {
	s.Body = v
	return s
}

func (s *ModifyRdsReadWeightResponse) Validate() error {
	return dara.Validate(s)
}
