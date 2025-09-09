// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolarDbReadWeightResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPolarDbReadWeightResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPolarDbReadWeightResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPolarDbReadWeightResponseBody) *ModifyPolarDbReadWeightResponse
	GetBody() *ModifyPolarDbReadWeightResponseBody
}

type ModifyPolarDbReadWeightResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPolarDbReadWeightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPolarDbReadWeightResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolarDbReadWeightResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolarDbReadWeightResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPolarDbReadWeightResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPolarDbReadWeightResponse) GetBody() *ModifyPolarDbReadWeightResponseBody {
	return s.Body
}

func (s *ModifyPolarDbReadWeightResponse) SetHeaders(v map[string]*string) *ModifyPolarDbReadWeightResponse {
	s.Headers = v
	return s
}

func (s *ModifyPolarDbReadWeightResponse) SetStatusCode(v int32) *ModifyPolarDbReadWeightResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPolarDbReadWeightResponse) SetBody(v *ModifyPolarDbReadWeightResponseBody) *ModifyPolarDbReadWeightResponse {
	s.Body = v
	return s
}

func (s *ModifyPolarDbReadWeightResponse) Validate() error {
	return dara.Validate(s)
}
