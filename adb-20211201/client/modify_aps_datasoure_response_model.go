// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApsDatasoureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApsDatasoureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApsDatasoureResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApsDatasoureResponseBody) *ModifyApsDatasoureResponse
	GetBody() *ModifyApsDatasoureResponseBody
}

type ModifyApsDatasoureResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApsDatasoureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApsDatasoureResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApsDatasoureResponse) GoString() string {
	return s.String()
}

func (s *ModifyApsDatasoureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApsDatasoureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApsDatasoureResponse) GetBody() *ModifyApsDatasoureResponseBody {
	return s.Body
}

func (s *ModifyApsDatasoureResponse) SetHeaders(v map[string]*string) *ModifyApsDatasoureResponse {
	s.Headers = v
	return s
}

func (s *ModifyApsDatasoureResponse) SetStatusCode(v int32) *ModifyApsDatasoureResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApsDatasoureResponse) SetBody(v *ModifyApsDatasoureResponseBody) *ModifyApsDatasoureResponse {
	s.Body = v
	return s
}

func (s *ModifyApsDatasoureResponse) Validate() error {
	return dara.Validate(s)
}
