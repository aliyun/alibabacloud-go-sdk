// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDtsJobNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDtsJobNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDtsJobNameResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDtsJobNameResponseBody) *ModifyDtsJobNameResponse
	GetBody() *ModifyDtsJobNameResponseBody
}

type ModifyDtsJobNameResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDtsJobNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDtsJobNameResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDtsJobNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDtsJobNameResponse) GetBody() *ModifyDtsJobNameResponseBody {
	return s.Body
}

func (s *ModifyDtsJobNameResponse) SetHeaders(v map[string]*string) *ModifyDtsJobNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyDtsJobNameResponse) SetStatusCode(v int32) *ModifyDtsJobNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDtsJobNameResponse) SetBody(v *ModifyDtsJobNameResponseBody) *ModifyDtsJobNameResponse {
	s.Body = v
	return s
}

func (s *ModifyDtsJobNameResponse) Validate() error {
	return dara.Validate(s)
}
