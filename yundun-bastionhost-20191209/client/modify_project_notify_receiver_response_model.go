// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyProjectNotifyReceiverResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyProjectNotifyReceiverResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyProjectNotifyReceiverResponse
	GetStatusCode() *int32
	SetBody(v *ModifyProjectNotifyReceiverResponseBody) *ModifyProjectNotifyReceiverResponse
	GetBody() *ModifyProjectNotifyReceiverResponseBody
}

type ModifyProjectNotifyReceiverResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyProjectNotifyReceiverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyProjectNotifyReceiverResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyProjectNotifyReceiverResponse) GoString() string {
	return s.String()
}

func (s *ModifyProjectNotifyReceiverResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyProjectNotifyReceiverResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyProjectNotifyReceiverResponse) GetBody() *ModifyProjectNotifyReceiverResponseBody {
	return s.Body
}

func (s *ModifyProjectNotifyReceiverResponse) SetHeaders(v map[string]*string) *ModifyProjectNotifyReceiverResponse {
	s.Headers = v
	return s
}

func (s *ModifyProjectNotifyReceiverResponse) SetStatusCode(v int32) *ModifyProjectNotifyReceiverResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyProjectNotifyReceiverResponse) SetBody(v *ModifyProjectNotifyReceiverResponseBody) *ModifyProjectNotifyReceiverResponse {
	s.Body = v
	return s
}

func (s *ModifyProjectNotifyReceiverResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
