// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperationCancelIgnoreSuspEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperationCancelIgnoreSuspEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperationCancelIgnoreSuspEventResponse
	GetStatusCode() *int32
	SetBody(v *OperationCancelIgnoreSuspEventResponseBody) *OperationCancelIgnoreSuspEventResponse
	GetBody() *OperationCancelIgnoreSuspEventResponseBody
}

type OperationCancelIgnoreSuspEventResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperationCancelIgnoreSuspEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperationCancelIgnoreSuspEventResponse) String() string {
	return dara.Prettify(s)
}

func (s OperationCancelIgnoreSuspEventResponse) GoString() string {
	return s.String()
}

func (s *OperationCancelIgnoreSuspEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperationCancelIgnoreSuspEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperationCancelIgnoreSuspEventResponse) GetBody() *OperationCancelIgnoreSuspEventResponseBody {
	return s.Body
}

func (s *OperationCancelIgnoreSuspEventResponse) SetHeaders(v map[string]*string) *OperationCancelIgnoreSuspEventResponse {
	s.Headers = v
	return s
}

func (s *OperationCancelIgnoreSuspEventResponse) SetStatusCode(v int32) *OperationCancelIgnoreSuspEventResponse {
	s.StatusCode = &v
	return s
}

func (s *OperationCancelIgnoreSuspEventResponse) SetBody(v *OperationCancelIgnoreSuspEventResponseBody) *OperationCancelIgnoreSuspEventResponse {
	s.Body = v
	return s
}

func (s *OperationCancelIgnoreSuspEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
