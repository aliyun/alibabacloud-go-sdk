// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRealTimeDialogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RealTimeDialogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RealTimeDialogResponse
	GetStatusCode() *int32
	SetId(v string) *RealTimeDialogResponse
	GetId() *string
	SetEvent(v string) *RealTimeDialogResponse
	GetEvent() *string
	SetBody(v *RealTimeDialogResponseBody) *RealTimeDialogResponse
	GetBody() *RealTimeDialogResponseBody
}

type RealTimeDialogResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                     `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                     `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RealTimeDialogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RealTimeDialogResponse) String() string {
	return dara.Prettify(s)
}

func (s RealTimeDialogResponse) GoString() string {
	return s.String()
}

func (s *RealTimeDialogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RealTimeDialogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RealTimeDialogResponse) GetId() *string {
	return s.Id
}

func (s *RealTimeDialogResponse) GetEvent() *string {
	return s.Event
}

func (s *RealTimeDialogResponse) GetBody() *RealTimeDialogResponseBody {
	return s.Body
}

func (s *RealTimeDialogResponse) SetHeaders(v map[string]*string) *RealTimeDialogResponse {
	s.Headers = v
	return s
}

func (s *RealTimeDialogResponse) SetStatusCode(v int32) *RealTimeDialogResponse {
	s.StatusCode = &v
	return s
}

func (s *RealTimeDialogResponse) SetId(v string) *RealTimeDialogResponse {
	s.Id = &v
	return s
}

func (s *RealTimeDialogResponse) SetEvent(v string) *RealTimeDialogResponse {
	s.Event = &v
	return s
}

func (s *RealTimeDialogResponse) SetBody(v *RealTimeDialogResponseBody) *RealTimeDialogResponse {
	s.Body = v
	return s
}

func (s *RealTimeDialogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
