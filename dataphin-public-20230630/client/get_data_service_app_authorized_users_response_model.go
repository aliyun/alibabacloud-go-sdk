// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceAppAuthorizedUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataServiceAppAuthorizedUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataServiceAppAuthorizedUsersResponse
	GetStatusCode() *int32
	SetBody(v *GetDataServiceAppAuthorizedUsersResponseBody) *GetDataServiceAppAuthorizedUsersResponse
	GetBody() *GetDataServiceAppAuthorizedUsersResponseBody
}

type GetDataServiceAppAuthorizedUsersResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServiceAppAuthorizedUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServiceAppAuthorizedUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppAuthorizedUsersResponse) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppAuthorizedUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataServiceAppAuthorizedUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataServiceAppAuthorizedUsersResponse) GetBody() *GetDataServiceAppAuthorizedUsersResponseBody {
	return s.Body
}

func (s *GetDataServiceAppAuthorizedUsersResponse) SetHeaders(v map[string]*string) *GetDataServiceAppAuthorizedUsersResponse {
	s.Headers = v
	return s
}

func (s *GetDataServiceAppAuthorizedUsersResponse) SetStatusCode(v int32) *GetDataServiceAppAuthorizedUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServiceAppAuthorizedUsersResponse) SetBody(v *GetDataServiceAppAuthorizedUsersResponseBody) *GetDataServiceAppAuthorizedUsersResponse {
	s.Body = v
	return s
}

func (s *GetDataServiceAppAuthorizedUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
