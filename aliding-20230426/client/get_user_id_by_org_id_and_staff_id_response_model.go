// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByOrgIdAndStaffIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserIdByOrgIdAndStaffIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserIdByOrgIdAndStaffIdResponse
	GetStatusCode() *int32
	SetBody(v *GetUserIdByOrgIdAndStaffIdResponseBody) *GetUserIdByOrgIdAndStaffIdResponse
	GetBody() *GetUserIdByOrgIdAndStaffIdResponseBody
}

type GetUserIdByOrgIdAndStaffIdResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserIdByOrgIdAndStaffIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserIdByOrgIdAndStaffIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByOrgIdAndStaffIdResponse) GoString() string {
	return s.String()
}

func (s *GetUserIdByOrgIdAndStaffIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserIdByOrgIdAndStaffIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserIdByOrgIdAndStaffIdResponse) GetBody() *GetUserIdByOrgIdAndStaffIdResponseBody {
	return s.Body
}

func (s *GetUserIdByOrgIdAndStaffIdResponse) SetHeaders(v map[string]*string) *GetUserIdByOrgIdAndStaffIdResponse {
	s.Headers = v
	return s
}

func (s *GetUserIdByOrgIdAndStaffIdResponse) SetStatusCode(v int32) *GetUserIdByOrgIdAndStaffIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserIdByOrgIdAndStaffIdResponse) SetBody(v *GetUserIdByOrgIdAndStaffIdResponseBody) *GetUserIdByOrgIdAndStaffIdResponse {
	s.Body = v
	return s
}

func (s *GetUserIdByOrgIdAndStaffIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
