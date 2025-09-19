// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMachineCanRebootResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMachineCanRebootResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMachineCanRebootResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMachineCanRebootResponseBody) *DescribeMachineCanRebootResponse
	GetBody() *DescribeMachineCanRebootResponseBody
}

type DescribeMachineCanRebootResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMachineCanRebootResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMachineCanRebootResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMachineCanRebootResponse) GoString() string {
	return s.String()
}

func (s *DescribeMachineCanRebootResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMachineCanRebootResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMachineCanRebootResponse) GetBody() *DescribeMachineCanRebootResponseBody {
	return s.Body
}

func (s *DescribeMachineCanRebootResponse) SetHeaders(v map[string]*string) *DescribeMachineCanRebootResponse {
	s.Headers = v
	return s
}

func (s *DescribeMachineCanRebootResponse) SetStatusCode(v int32) *DescribeMachineCanRebootResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMachineCanRebootResponse) SetBody(v *DescribeMachineCanRebootResponseBody) *DescribeMachineCanRebootResponse {
	s.Body = v
	return s
}

func (s *DescribeMachineCanRebootResponse) Validate() error {
	return dara.Validate(s)
}
