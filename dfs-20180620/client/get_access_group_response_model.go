// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccessGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccessGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetAccessGroupResponseBody) *GetAccessGroupResponse
	GetBody() *GetAccessGroupResponseBody
}

type GetAccessGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccessGroupResponse) GoString() string {
	return s.String()
}

func (s *GetAccessGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccessGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccessGroupResponse) GetBody() *GetAccessGroupResponseBody {
	return s.Body
}

func (s *GetAccessGroupResponse) SetHeaders(v map[string]*string) *GetAccessGroupResponse {
	s.Headers = v
	return s
}

func (s *GetAccessGroupResponse) SetStatusCode(v int32) *GetAccessGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessGroupResponse) SetBody(v *GetAccessGroupResponseBody) *GetAccessGroupResponse {
	s.Body = v
	return s
}

func (s *GetAccessGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
