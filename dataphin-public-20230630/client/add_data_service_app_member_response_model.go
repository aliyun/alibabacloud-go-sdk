// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataServiceAppMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDataServiceAppMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDataServiceAppMemberResponse
	GetStatusCode() *int32
	SetBody(v *AddDataServiceAppMemberResponseBody) *AddDataServiceAppMemberResponse
	GetBody() *AddDataServiceAppMemberResponseBody
}

type AddDataServiceAppMemberResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDataServiceAppMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDataServiceAppMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDataServiceAppMemberResponse) GoString() string {
	return s.String()
}

func (s *AddDataServiceAppMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDataServiceAppMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDataServiceAppMemberResponse) GetBody() *AddDataServiceAppMemberResponseBody {
	return s.Body
}

func (s *AddDataServiceAppMemberResponse) SetHeaders(v map[string]*string) *AddDataServiceAppMemberResponse {
	s.Headers = v
	return s
}

func (s *AddDataServiceAppMemberResponse) SetStatusCode(v int32) *AddDataServiceAppMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDataServiceAppMemberResponse) SetBody(v *AddDataServiceAppMemberResponseBody) *AddDataServiceAppMemberResponse {
	s.Body = v
	return s
}

func (s *AddDataServiceAppMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
