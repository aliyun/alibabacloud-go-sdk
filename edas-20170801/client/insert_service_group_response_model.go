// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertServiceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertServiceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertServiceGroupResponse
	GetStatusCode() *int32
	SetBody(v *InsertServiceGroupResponseBody) *InsertServiceGroupResponse
	GetBody() *InsertServiceGroupResponseBody
}

type InsertServiceGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertServiceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertServiceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *InsertServiceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertServiceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertServiceGroupResponse) GetBody() *InsertServiceGroupResponseBody {
	return s.Body
}

func (s *InsertServiceGroupResponse) SetHeaders(v map[string]*string) *InsertServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *InsertServiceGroupResponse) SetStatusCode(v int32) *InsertServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertServiceGroupResponse) SetBody(v *InsertServiceGroupResponseBody) *InsertServiceGroupResponse {
	s.Body = v
	return s
}

func (s *InsertServiceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
