// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrafficControlTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTrafficControlTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTrafficControlTargetResponse
	GetStatusCode() *int32
	SetBody(v *GetTrafficControlTargetResponseBody) *GetTrafficControlTargetResponse
	GetBody() *GetTrafficControlTargetResponseBody
}

type GetTrafficControlTargetResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrafficControlTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrafficControlTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTrafficControlTargetResponse) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTrafficControlTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTrafficControlTargetResponse) GetBody() *GetTrafficControlTargetResponseBody {
	return s.Body
}

func (s *GetTrafficControlTargetResponse) SetHeaders(v map[string]*string) *GetTrafficControlTargetResponse {
	s.Headers = v
	return s
}

func (s *GetTrafficControlTargetResponse) SetStatusCode(v int32) *GetTrafficControlTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrafficControlTargetResponse) SetBody(v *GetTrafficControlTargetResponseBody) *GetTrafficControlTargetResponse {
	s.Body = v
	return s
}

func (s *GetTrafficControlTargetResponse) Validate() error {
	return dara.Validate(s)
}
