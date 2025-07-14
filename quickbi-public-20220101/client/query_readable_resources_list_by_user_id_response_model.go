// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReadableResourcesListByUserIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryReadableResourcesListByUserIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryReadableResourcesListByUserIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryReadableResourcesListByUserIdResponseBody) *QueryReadableResourcesListByUserIdResponse
	GetBody() *QueryReadableResourcesListByUserIdResponseBody
}

type QueryReadableResourcesListByUserIdResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryReadableResourcesListByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryReadableResourcesListByUserIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryReadableResourcesListByUserIdResponse) GoString() string {
	return s.String()
}

func (s *QueryReadableResourcesListByUserIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryReadableResourcesListByUserIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryReadableResourcesListByUserIdResponse) GetBody() *QueryReadableResourcesListByUserIdResponseBody {
	return s.Body
}

func (s *QueryReadableResourcesListByUserIdResponse) SetHeaders(v map[string]*string) *QueryReadableResourcesListByUserIdResponse {
	s.Headers = v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponse) SetStatusCode(v int32) *QueryReadableResourcesListByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponse) SetBody(v *QueryReadableResourcesListByUserIdResponseBody) *QueryReadableResourcesListByUserIdResponse {
	s.Body = v
	return s
}

func (s *QueryReadableResourcesListByUserIdResponse) Validate() error {
	return dara.Validate(s)
}
