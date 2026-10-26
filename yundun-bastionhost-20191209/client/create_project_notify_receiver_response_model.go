// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectNotifyReceiverResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProjectNotifyReceiverResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProjectNotifyReceiverResponse
	GetStatusCode() *int32
	SetBody(v *CreateProjectNotifyReceiverResponseBody) *CreateProjectNotifyReceiverResponse
	GetBody() *CreateProjectNotifyReceiverResponseBody
}

type CreateProjectNotifyReceiverResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProjectNotifyReceiverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProjectNotifyReceiverResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectNotifyReceiverResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectNotifyReceiverResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProjectNotifyReceiverResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProjectNotifyReceiverResponse) GetBody() *CreateProjectNotifyReceiverResponseBody {
	return s.Body
}

func (s *CreateProjectNotifyReceiverResponse) SetHeaders(v map[string]*string) *CreateProjectNotifyReceiverResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectNotifyReceiverResponse) SetStatusCode(v int32) *CreateProjectNotifyReceiverResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectNotifyReceiverResponse) SetBody(v *CreateProjectNotifyReceiverResponseBody) *CreateProjectNotifyReceiverResponse {
	s.Body = v
	return s
}

func (s *CreateProjectNotifyReceiverResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
