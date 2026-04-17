// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCreateLindormV2InstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckCreateLindormV2InstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckCreateLindormV2InstanceResponse
	GetStatusCode() *int32
	SetBody(v *CheckCreateLindormV2InstanceResponseBody) *CheckCreateLindormV2InstanceResponse
	GetBody() *CheckCreateLindormV2InstanceResponseBody
}

type CheckCreateLindormV2InstanceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckCreateLindormV2InstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckCreateLindormV2InstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckCreateLindormV2InstanceResponse) GoString() string {
	return s.String()
}

func (s *CheckCreateLindormV2InstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckCreateLindormV2InstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckCreateLindormV2InstanceResponse) GetBody() *CheckCreateLindormV2InstanceResponseBody {
	return s.Body
}

func (s *CheckCreateLindormV2InstanceResponse) SetHeaders(v map[string]*string) *CheckCreateLindormV2InstanceResponse {
	s.Headers = v
	return s
}

func (s *CheckCreateLindormV2InstanceResponse) SetStatusCode(v int32) *CheckCreateLindormV2InstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckCreateLindormV2InstanceResponse) SetBody(v *CheckCreateLindormV2InstanceResponseBody) *CheckCreateLindormV2InstanceResponse {
	s.Body = v
	return s
}

func (s *CheckCreateLindormV2InstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
