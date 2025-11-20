// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketOperateRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTicketOperateRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTicketOperateRecordResponse
	GetStatusCode() *int32
	SetBody(v *ListTicketOperateRecordResponseBody) *ListTicketOperateRecordResponse
	GetBody() *ListTicketOperateRecordResponseBody
}

type ListTicketOperateRecordResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTicketOperateRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTicketOperateRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTicketOperateRecordResponse) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTicketOperateRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTicketOperateRecordResponse) GetBody() *ListTicketOperateRecordResponseBody {
	return s.Body
}

func (s *ListTicketOperateRecordResponse) SetHeaders(v map[string]*string) *ListTicketOperateRecordResponse {
	s.Headers = v
	return s
}

func (s *ListTicketOperateRecordResponse) SetStatusCode(v int32) *ListTicketOperateRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTicketOperateRecordResponse) SetBody(v *ListTicketOperateRecordResponseBody) *ListTicketOperateRecordResponse {
	s.Body = v
	return s
}

func (s *ListTicketOperateRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
