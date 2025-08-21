// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpDefenseThresholdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIpDefenseThresholdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIpDefenseThresholdResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIpDefenseThresholdResponseBody) *ModifyIpDefenseThresholdResponse
	GetBody() *ModifyIpDefenseThresholdResponseBody
}

type ModifyIpDefenseThresholdResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIpDefenseThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIpDefenseThresholdResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpDefenseThresholdResponse) GoString() string {
	return s.String()
}

func (s *ModifyIpDefenseThresholdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIpDefenseThresholdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIpDefenseThresholdResponse) GetBody() *ModifyIpDefenseThresholdResponseBody {
	return s.Body
}

func (s *ModifyIpDefenseThresholdResponse) SetHeaders(v map[string]*string) *ModifyIpDefenseThresholdResponse {
	s.Headers = v
	return s
}

func (s *ModifyIpDefenseThresholdResponse) SetStatusCode(v int32) *ModifyIpDefenseThresholdResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIpDefenseThresholdResponse) SetBody(v *ModifyIpDefenseThresholdResponseBody) *ModifyIpDefenseThresholdResponse {
	s.Body = v
	return s
}

func (s *ModifyIpDefenseThresholdResponse) Validate() error {
	return dara.Validate(s)
}
