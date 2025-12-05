// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDatabasesFromGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveDatabasesFromGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveDatabasesFromGroupResponse
	GetStatusCode() *int32
	SetBody(v *RemoveDatabasesFromGroupResponseBody) *RemoveDatabasesFromGroupResponse
	GetBody() *RemoveDatabasesFromGroupResponseBody
}

type RemoveDatabasesFromGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveDatabasesFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveDatabasesFromGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveDatabasesFromGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveDatabasesFromGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveDatabasesFromGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveDatabasesFromGroupResponse) GetBody() *RemoveDatabasesFromGroupResponseBody {
	return s.Body
}

func (s *RemoveDatabasesFromGroupResponse) SetHeaders(v map[string]*string) *RemoveDatabasesFromGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveDatabasesFromGroupResponse) SetStatusCode(v int32) *RemoveDatabasesFromGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDatabasesFromGroupResponse) SetBody(v *RemoveDatabasesFromGroupResponseBody) *RemoveDatabasesFromGroupResponse {
	s.Body = v
	return s
}

func (s *RemoveDatabasesFromGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
