// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDatabasesToGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDatabasesToGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDatabasesToGroupResponse
	GetStatusCode() *int32
	SetBody(v *AddDatabasesToGroupResponseBody) *AddDatabasesToGroupResponse
	GetBody() *AddDatabasesToGroupResponseBody
}

type AddDatabasesToGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDatabasesToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDatabasesToGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDatabasesToGroupResponse) GoString() string {
	return s.String()
}

func (s *AddDatabasesToGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDatabasesToGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDatabasesToGroupResponse) GetBody() *AddDatabasesToGroupResponseBody {
	return s.Body
}

func (s *AddDatabasesToGroupResponse) SetHeaders(v map[string]*string) *AddDatabasesToGroupResponse {
	s.Headers = v
	return s
}

func (s *AddDatabasesToGroupResponse) SetStatusCode(v int32) *AddDatabasesToGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDatabasesToGroupResponse) SetBody(v *AddDatabasesToGroupResponseBody) *AddDatabasesToGroupResponse {
	s.Body = v
	return s
}

func (s *AddDatabasesToGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
