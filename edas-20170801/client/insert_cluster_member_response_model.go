// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertClusterMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertClusterMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertClusterMemberResponse
	GetStatusCode() *int32
	SetBody(v *InsertClusterMemberResponseBody) *InsertClusterMemberResponse
	GetBody() *InsertClusterMemberResponseBody
}

type InsertClusterMemberResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertClusterMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertClusterMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertClusterMemberResponse) GoString() string {
	return s.String()
}

func (s *InsertClusterMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertClusterMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertClusterMemberResponse) GetBody() *InsertClusterMemberResponseBody {
	return s.Body
}

func (s *InsertClusterMemberResponse) SetHeaders(v map[string]*string) *InsertClusterMemberResponse {
	s.Headers = v
	return s
}

func (s *InsertClusterMemberResponse) SetStatusCode(v int32) *InsertClusterMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertClusterMemberResponse) SetBody(v *InsertClusterMemberResponseBody) *InsertClusterMemberResponse {
	s.Body = v
	return s
}

func (s *InsertClusterMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
