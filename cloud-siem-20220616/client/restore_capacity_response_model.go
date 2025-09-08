// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreCapacityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestoreCapacityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestoreCapacityResponse
	GetStatusCode() *int32
	SetBody(v *RestoreCapacityResponseBody) *RestoreCapacityResponse
	GetBody() *RestoreCapacityResponseBody
}

type RestoreCapacityResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreCapacityResponse) String() string {
	return dara.Prettify(s)
}

func (s RestoreCapacityResponse) GoString() string {
	return s.String()
}

func (s *RestoreCapacityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestoreCapacityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestoreCapacityResponse) GetBody() *RestoreCapacityResponseBody {
	return s.Body
}

func (s *RestoreCapacityResponse) SetHeaders(v map[string]*string) *RestoreCapacityResponse {
	s.Headers = v
	return s
}

func (s *RestoreCapacityResponse) SetStatusCode(v int32) *RestoreCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreCapacityResponse) SetBody(v *RestoreCapacityResponseBody) *RestoreCapacityResponse {
	s.Body = v
	return s
}

func (s *RestoreCapacityResponse) Validate() error {
	return dara.Validate(s)
}
