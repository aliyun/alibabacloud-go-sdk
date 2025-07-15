// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllInstanceIdListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAllInstanceIdListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAllInstanceIdListResponse
	GetStatusCode() *int32
	SetBody(v *GetAllInstanceIdListResponseBody) *GetAllInstanceIdListResponse
	GetBody() *GetAllInstanceIdListResponseBody
}

type GetAllInstanceIdListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAllInstanceIdListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAllInstanceIdListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAllInstanceIdListResponse) GoString() string {
	return s.String()
}

func (s *GetAllInstanceIdListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAllInstanceIdListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAllInstanceIdListResponse) GetBody() *GetAllInstanceIdListResponseBody {
	return s.Body
}

func (s *GetAllInstanceIdListResponse) SetHeaders(v map[string]*string) *GetAllInstanceIdListResponse {
	s.Headers = v
	return s
}

func (s *GetAllInstanceIdListResponse) SetStatusCode(v int32) *GetAllInstanceIdListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAllInstanceIdListResponse) SetBody(v *GetAllInstanceIdListResponseBody) *GetAllInstanceIdListResponse {
	s.Body = v
	return s
}

func (s *GetAllInstanceIdListResponse) Validate() error {
	return dara.Validate(s)
}
