// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutContactGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutContactGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutContactGroupResponse
	GetStatusCode() *int32
	SetBody(v *PutContactGroupResponseBody) *PutContactGroupResponse
	GetBody() *PutContactGroupResponseBody
}

type PutContactGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutContactGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s PutContactGroupResponse) GoString() string {
	return s.String()
}

func (s *PutContactGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutContactGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutContactGroupResponse) GetBody() *PutContactGroupResponseBody {
	return s.Body
}

func (s *PutContactGroupResponse) SetHeaders(v map[string]*string) *PutContactGroupResponse {
	s.Headers = v
	return s
}

func (s *PutContactGroupResponse) SetStatusCode(v int32) *PutContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *PutContactGroupResponse) SetBody(v *PutContactGroupResponseBody) *PutContactGroupResponse {
	s.Body = v
	return s
}

func (s *PutContactGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
