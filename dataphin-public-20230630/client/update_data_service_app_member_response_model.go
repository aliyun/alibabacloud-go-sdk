// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataServiceAppMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataServiceAppMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataServiceAppMemberResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataServiceAppMemberResponseBody) *UpdateDataServiceAppMemberResponse
	GetBody() *UpdateDataServiceAppMemberResponseBody
}

type UpdateDataServiceAppMemberResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataServiceAppMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataServiceAppMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataServiceAppMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataServiceAppMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataServiceAppMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataServiceAppMemberResponse) GetBody() *UpdateDataServiceAppMemberResponseBody {
	return s.Body
}

func (s *UpdateDataServiceAppMemberResponse) SetHeaders(v map[string]*string) *UpdateDataServiceAppMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataServiceAppMemberResponse) SetStatusCode(v int32) *UpdateDataServiceAppMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataServiceAppMemberResponse) SetBody(v *UpdateDataServiceAppMemberResponseBody) *UpdateDataServiceAppMemberResponse {
	s.Body = v
	return s
}

func (s *UpdateDataServiceAppMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
