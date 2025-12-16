// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAppGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveAppGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveAppGroupResponse
	GetStatusCode() *int32
	SetBody(v *RemoveAppGroupResponseBody) *RemoveAppGroupResponse
	GetBody() *RemoveAppGroupResponseBody
}

type RemoveAppGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveAppGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveAppGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveAppGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveAppGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveAppGroupResponse) GetBody() *RemoveAppGroupResponseBody {
	return s.Body
}

func (s *RemoveAppGroupResponse) SetHeaders(v map[string]*string) *RemoveAppGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveAppGroupResponse) SetStatusCode(v int32) *RemoveAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveAppGroupResponse) SetBody(v *RemoveAppGroupResponseBody) *RemoveAppGroupResponse {
	s.Body = v
	return s
}

func (s *RemoveAppGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
