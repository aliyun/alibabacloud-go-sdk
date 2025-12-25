// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRoomPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddRoomPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddRoomPlanResponse
	GetStatusCode() *int32
	SetBody(v *AddRoomPlanResponseBody) *AddRoomPlanResponse
	GetBody() *AddRoomPlanResponseBody
}

type AddRoomPlanResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRoomPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRoomPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s AddRoomPlanResponse) GoString() string {
	return s.String()
}

func (s *AddRoomPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddRoomPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddRoomPlanResponse) GetBody() *AddRoomPlanResponseBody {
	return s.Body
}

func (s *AddRoomPlanResponse) SetHeaders(v map[string]*string) *AddRoomPlanResponse {
	s.Headers = v
	return s
}

func (s *AddRoomPlanResponse) SetStatusCode(v int32) *AddRoomPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRoomPlanResponse) SetBody(v *AddRoomPlanResponseBody) *AddRoomPlanResponse {
	s.Body = v
	return s
}

func (s *AddRoomPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
