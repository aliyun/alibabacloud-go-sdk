// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataServiceProjectMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDataServiceProjectMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDataServiceProjectMemberResponse
	GetStatusCode() *int32
	SetBody(v *AddDataServiceProjectMemberResponseBody) *AddDataServiceProjectMemberResponse
	GetBody() *AddDataServiceProjectMemberResponseBody
}

type AddDataServiceProjectMemberResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDataServiceProjectMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDataServiceProjectMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDataServiceProjectMemberResponse) GoString() string {
	return s.String()
}

func (s *AddDataServiceProjectMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDataServiceProjectMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDataServiceProjectMemberResponse) GetBody() *AddDataServiceProjectMemberResponseBody {
	return s.Body
}

func (s *AddDataServiceProjectMemberResponse) SetHeaders(v map[string]*string) *AddDataServiceProjectMemberResponse {
	s.Headers = v
	return s
}

func (s *AddDataServiceProjectMemberResponse) SetStatusCode(v int32) *AddDataServiceProjectMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDataServiceProjectMemberResponse) SetBody(v *AddDataServiceProjectMemberResponseBody) *AddDataServiceProjectMemberResponse {
	s.Body = v
	return s
}

func (s *AddDataServiceProjectMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
