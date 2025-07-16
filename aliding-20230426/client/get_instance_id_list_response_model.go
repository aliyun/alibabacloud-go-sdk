// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceIdListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceIdListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceIdListResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceIdListResponseBody) *GetInstanceIdListResponse
	GetBody() *GetInstanceIdListResponseBody
}

type GetInstanceIdListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceIdListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceIdListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceIdListResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceIdListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceIdListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceIdListResponse) GetBody() *GetInstanceIdListResponseBody {
	return s.Body
}

func (s *GetInstanceIdListResponse) SetHeaders(v map[string]*string) *GetInstanceIdListResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceIdListResponse) SetStatusCode(v int32) *GetInstanceIdListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceIdListResponse) SetBody(v *GetInstanceIdListResponseBody) *GetInstanceIdListResponse {
	s.Body = v
	return s
}

func (s *GetInstanceIdListResponse) Validate() error {
	return dara.Validate(s)
}
