// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApsWorkloadNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApsWorkloadNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApsWorkloadNameResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApsWorkloadNameResponseBody) *ModifyApsWorkloadNameResponse
	GetBody() *ModifyApsWorkloadNameResponseBody
}

type ModifyApsWorkloadNameResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApsWorkloadNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApsWorkloadNameResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApsWorkloadNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyApsWorkloadNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApsWorkloadNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApsWorkloadNameResponse) GetBody() *ModifyApsWorkloadNameResponseBody {
	return s.Body
}

func (s *ModifyApsWorkloadNameResponse) SetHeaders(v map[string]*string) *ModifyApsWorkloadNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyApsWorkloadNameResponse) SetStatusCode(v int32) *ModifyApsWorkloadNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApsWorkloadNameResponse) SetBody(v *ModifyApsWorkloadNameResponseBody) *ModifyApsWorkloadNameResponse {
	s.Body = v
	return s
}

func (s *ModifyApsWorkloadNameResponse) Validate() error {
	return dara.Validate(s)
}
