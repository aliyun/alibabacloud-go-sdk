// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemberAccessRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMemberAccessRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMemberAccessRecordResponse
	GetStatusCode() *int32
	SetBody(v *ListMemberAccessRecordResponseBody) *ListMemberAccessRecordResponse
	GetBody() *ListMemberAccessRecordResponseBody
}

type ListMemberAccessRecordResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMemberAccessRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMemberAccessRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMemberAccessRecordResponse) GoString() string {
	return s.String()
}

func (s *ListMemberAccessRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMemberAccessRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMemberAccessRecordResponse) GetBody() *ListMemberAccessRecordResponseBody {
	return s.Body
}

func (s *ListMemberAccessRecordResponse) SetHeaders(v map[string]*string) *ListMemberAccessRecordResponse {
	s.Headers = v
	return s
}

func (s *ListMemberAccessRecordResponse) SetStatusCode(v int32) *ListMemberAccessRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMemberAccessRecordResponse) SetBody(v *ListMemberAccessRecordResponseBody) *ListMemberAccessRecordResponse {
	s.Body = v
	return s
}

func (s *ListMemberAccessRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
