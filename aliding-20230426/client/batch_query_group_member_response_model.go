// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryGroupMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchQueryGroupMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchQueryGroupMemberResponse
	GetStatusCode() *int32
	SetBody(v *BatchQueryGroupMemberResponseBody) *BatchQueryGroupMemberResponse
	GetBody() *BatchQueryGroupMemberResponseBody
}

type BatchQueryGroupMemberResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQueryGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchQueryGroupMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryGroupMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchQueryGroupMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchQueryGroupMemberResponse) GetBody() *BatchQueryGroupMemberResponseBody {
	return s.Body
}

func (s *BatchQueryGroupMemberResponse) SetHeaders(v map[string]*string) *BatchQueryGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryGroupMemberResponse) SetStatusCode(v int32) *BatchQueryGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchQueryGroupMemberResponse) SetBody(v *BatchQueryGroupMemberResponseBody) *BatchQueryGroupMemberResponse {
	s.Body = v
	return s
}

func (s *BatchQueryGroupMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
