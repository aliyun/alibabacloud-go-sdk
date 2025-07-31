// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApiGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataServiceApiGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataServiceApiGroupsResponse
	GetStatusCode() *int32
	SetBody(v *GetDataServiceApiGroupsResponseBody) *GetDataServiceApiGroupsResponse
	GetBody() *GetDataServiceApiGroupsResponseBody
}

type GetDataServiceApiGroupsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServiceApiGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServiceApiGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiGroupsResponse) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataServiceApiGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataServiceApiGroupsResponse) GetBody() *GetDataServiceApiGroupsResponseBody {
	return s.Body
}

func (s *GetDataServiceApiGroupsResponse) SetHeaders(v map[string]*string) *GetDataServiceApiGroupsResponse {
	s.Headers = v
	return s
}

func (s *GetDataServiceApiGroupsResponse) SetStatusCode(v int32) *GetDataServiceApiGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServiceApiGroupsResponse) SetBody(v *GetDataServiceApiGroupsResponseBody) *GetDataServiceApiGroupsResponse {
	s.Body = v
	return s
}

func (s *GetDataServiceApiGroupsResponse) Validate() error {
	return dara.Validate(s)
}
