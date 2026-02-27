// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceAppMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataServiceAppMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataServiceAppMembersResponse
	GetStatusCode() *int32
	SetBody(v *GetDataServiceAppMembersResponseBody) *GetDataServiceAppMembersResponse
	GetBody() *GetDataServiceAppMembersResponseBody
}

type GetDataServiceAppMembersResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServiceAppMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServiceAppMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAppMembersResponse) GoString() string {
	return s.String()
}

func (s *GetDataServiceAppMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataServiceAppMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataServiceAppMembersResponse) GetBody() *GetDataServiceAppMembersResponseBody {
	return s.Body
}

func (s *GetDataServiceAppMembersResponse) SetHeaders(v map[string]*string) *GetDataServiceAppMembersResponse {
	s.Headers = v
	return s
}

func (s *GetDataServiceAppMembersResponse) SetStatusCode(v int32) *GetDataServiceAppMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServiceAppMembersResponse) SetBody(v *GetDataServiceAppMembersResponseBody) *GetDataServiceAppMembersResponse {
	s.Body = v
	return s
}

func (s *GetDataServiceAppMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
