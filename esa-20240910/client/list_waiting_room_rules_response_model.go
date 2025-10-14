// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWaitingRoomRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWaitingRoomRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWaitingRoomRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListWaitingRoomRulesResponseBody) *ListWaitingRoomRulesResponse
	GetBody() *ListWaitingRoomRulesResponseBody
}

type ListWaitingRoomRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWaitingRoomRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWaitingRoomRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWaitingRoomRulesResponse) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWaitingRoomRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWaitingRoomRulesResponse) GetBody() *ListWaitingRoomRulesResponseBody {
	return s.Body
}

func (s *ListWaitingRoomRulesResponse) SetHeaders(v map[string]*string) *ListWaitingRoomRulesResponse {
	s.Headers = v
	return s
}

func (s *ListWaitingRoomRulesResponse) SetStatusCode(v int32) *ListWaitingRoomRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWaitingRoomRulesResponse) SetBody(v *ListWaitingRoomRulesResponseBody) *ListWaitingRoomRulesResponse {
	s.Body = v
	return s
}

func (s *ListWaitingRoomRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
