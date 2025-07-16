// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFCTriggerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFCTriggerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFCTriggerResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFCTriggerResponseBody) *DescribeFCTriggerResponse
	GetBody() *DescribeFCTriggerResponseBody
}

type DescribeFCTriggerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFCTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFCTriggerResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFCTriggerResponse) GoString() string {
	return s.String()
}

func (s *DescribeFCTriggerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFCTriggerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFCTriggerResponse) GetBody() *DescribeFCTriggerResponseBody {
	return s.Body
}

func (s *DescribeFCTriggerResponse) SetHeaders(v map[string]*string) *DescribeFCTriggerResponse {
	s.Headers = v
	return s
}

func (s *DescribeFCTriggerResponse) SetStatusCode(v int32) *DescribeFCTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFCTriggerResponse) SetBody(v *DescribeFCTriggerResponseBody) *DescribeFCTriggerResponse {
	s.Body = v
	return s
}

func (s *DescribeFCTriggerResponse) Validate() error {
	return dara.Validate(s)
}
