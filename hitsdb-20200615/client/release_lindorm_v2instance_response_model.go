// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseLindormV2InstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseLindormV2InstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseLindormV2InstanceResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseLindormV2InstanceResponseBody) *ReleaseLindormV2InstanceResponse
	GetBody() *ReleaseLindormV2InstanceResponseBody
}

type ReleaseLindormV2InstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseLindormV2InstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseLindormV2InstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseLindormV2InstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseLindormV2InstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseLindormV2InstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseLindormV2InstanceResponse) GetBody() *ReleaseLindormV2InstanceResponseBody {
	return s.Body
}

func (s *ReleaseLindormV2InstanceResponse) SetHeaders(v map[string]*string) *ReleaseLindormV2InstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseLindormV2InstanceResponse) SetStatusCode(v int32) *ReleaseLindormV2InstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseLindormV2InstanceResponse) SetBody(v *ReleaseLindormV2InstanceResponseBody) *ReleaseLindormV2InstanceResponse {
	s.Body = v
	return s
}

func (s *ReleaseLindormV2InstanceResponse) Validate() error {
	return dara.Validate(s)
}
