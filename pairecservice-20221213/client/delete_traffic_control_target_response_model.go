// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficControlTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTrafficControlTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTrafficControlTargetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTrafficControlTargetResponseBody) *DeleteTrafficControlTargetResponse
	GetBody() *DeleteTrafficControlTargetResponseBody
}

type DeleteTrafficControlTargetResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTrafficControlTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTrafficControlTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficControlTargetResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrafficControlTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTrafficControlTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTrafficControlTargetResponse) GetBody() *DeleteTrafficControlTargetResponseBody {
	return s.Body
}

func (s *DeleteTrafficControlTargetResponse) SetHeaders(v map[string]*string) *DeleteTrafficControlTargetResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrafficControlTargetResponse) SetStatusCode(v int32) *DeleteTrafficControlTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrafficControlTargetResponse) SetBody(v *DeleteTrafficControlTargetResponseBody) *DeleteTrafficControlTargetResponse {
	s.Body = v
	return s
}

func (s *DeleteTrafficControlTargetResponse) Validate() error {
	return dara.Validate(s)
}
