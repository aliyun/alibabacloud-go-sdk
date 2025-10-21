// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWordGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWordGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWordGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetWordGroupResponseBody) *GetWordGroupResponse
	GetBody() *GetWordGroupResponseBody
}

type GetWordGroupResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWordGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWordGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWordGroupResponse) GoString() string {
	return s.String()
}

func (s *GetWordGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWordGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWordGroupResponse) GetBody() *GetWordGroupResponseBody {
	return s.Body
}

func (s *GetWordGroupResponse) SetHeaders(v map[string]*string) *GetWordGroupResponse {
	s.Headers = v
	return s
}

func (s *GetWordGroupResponse) SetStatusCode(v int32) *GetWordGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWordGroupResponse) SetBody(v *GetWordGroupResponseBody) *GetWordGroupResponse {
	s.Body = v
	return s
}

func (s *GetWordGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
