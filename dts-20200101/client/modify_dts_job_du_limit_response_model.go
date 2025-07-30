// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDtsJobDuLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDtsJobDuLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDtsJobDuLimitResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDtsJobDuLimitResponseBody) *ModifyDtsJobDuLimitResponse
	GetBody() *ModifyDtsJobDuLimitResponseBody
}

type ModifyDtsJobDuLimitResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDtsJobDuLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDtsJobDuLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobDuLimitResponse) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobDuLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDtsJobDuLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDtsJobDuLimitResponse) GetBody() *ModifyDtsJobDuLimitResponseBody {
	return s.Body
}

func (s *ModifyDtsJobDuLimitResponse) SetHeaders(v map[string]*string) *ModifyDtsJobDuLimitResponse {
	s.Headers = v
	return s
}

func (s *ModifyDtsJobDuLimitResponse) SetStatusCode(v int32) *ModifyDtsJobDuLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDtsJobDuLimitResponse) SetBody(v *ModifyDtsJobDuLimitResponseBody) *ModifyDtsJobDuLimitResponse {
	s.Body = v
	return s
}

func (s *ModifyDtsJobDuLimitResponse) Validate() error {
	return dara.Validate(s)
}
