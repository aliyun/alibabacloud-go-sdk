// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceProjectAddableUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataServiceProjectAddableUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataServiceProjectAddableUsersResponse
	GetStatusCode() *int32
	SetBody(v *GetDataServiceProjectAddableUsersResponseBody) *GetDataServiceProjectAddableUsersResponse
	GetBody() *GetDataServiceProjectAddableUsersResponseBody
}

type GetDataServiceProjectAddableUsersResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServiceProjectAddableUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServiceProjectAddableUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceProjectAddableUsersResponse) GoString() string {
	return s.String()
}

func (s *GetDataServiceProjectAddableUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataServiceProjectAddableUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataServiceProjectAddableUsersResponse) GetBody() *GetDataServiceProjectAddableUsersResponseBody {
	return s.Body
}

func (s *GetDataServiceProjectAddableUsersResponse) SetHeaders(v map[string]*string) *GetDataServiceProjectAddableUsersResponse {
	s.Headers = v
	return s
}

func (s *GetDataServiceProjectAddableUsersResponse) SetStatusCode(v int32) *GetDataServiceProjectAddableUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServiceProjectAddableUsersResponse) SetBody(v *GetDataServiceProjectAddableUsersResponseBody) *GetDataServiceProjectAddableUsersResponse {
	s.Body = v
	return s
}

func (s *GetDataServiceProjectAddableUsersResponse) Validate() error {
	return dara.Validate(s)
}
