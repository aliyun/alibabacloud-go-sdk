// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseThresholdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDefenseThresholdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDefenseThresholdResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDefenseThresholdResponseBody) *ModifyDefenseThresholdResponse
	GetBody() *ModifyDefenseThresholdResponseBody
}

type ModifyDefenseThresholdResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefenseThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDefenseThresholdResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseThresholdResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseThresholdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDefenseThresholdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDefenseThresholdResponse) GetBody() *ModifyDefenseThresholdResponseBody {
	return s.Body
}

func (s *ModifyDefenseThresholdResponse) SetHeaders(v map[string]*string) *ModifyDefenseThresholdResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseThresholdResponse) SetStatusCode(v int32) *ModifyDefenseThresholdResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseThresholdResponse) SetBody(v *ModifyDefenseThresholdResponseBody) *ModifyDefenseThresholdResponse {
	s.Body = v
	return s
}

func (s *ModifyDefenseThresholdResponse) Validate() error {
	return dara.Validate(s)
}
