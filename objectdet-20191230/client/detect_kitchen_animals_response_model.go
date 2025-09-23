// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectKitchenAnimalsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectKitchenAnimalsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectKitchenAnimalsResponse
	GetStatusCode() *int32
	SetBody(v *DetectKitchenAnimalsResponseBody) *DetectKitchenAnimalsResponse
	GetBody() *DetectKitchenAnimalsResponseBody
}

type DetectKitchenAnimalsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectKitchenAnimalsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectKitchenAnimalsResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectKitchenAnimalsResponse) GoString() string {
	return s.String()
}

func (s *DetectKitchenAnimalsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectKitchenAnimalsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectKitchenAnimalsResponse) GetBody() *DetectKitchenAnimalsResponseBody {
	return s.Body
}

func (s *DetectKitchenAnimalsResponse) SetHeaders(v map[string]*string) *DetectKitchenAnimalsResponse {
	s.Headers = v
	return s
}

func (s *DetectKitchenAnimalsResponse) SetStatusCode(v int32) *DetectKitchenAnimalsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectKitchenAnimalsResponse) SetBody(v *DetectKitchenAnimalsResponseBody) *DetectKitchenAnimalsResponse {
	s.Body = v
	return s
}

func (s *DetectKitchenAnimalsResponse) Validate() error {
	return dara.Validate(s)
}
