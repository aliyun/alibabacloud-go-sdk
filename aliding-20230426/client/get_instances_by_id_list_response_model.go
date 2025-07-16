// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstancesByIdListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstancesByIdListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstancesByIdListResponse
	GetStatusCode() *int32
	SetBody(v *GetInstancesByIdListResponseBody) *GetInstancesByIdListResponse
	GetBody() *GetInstancesByIdListResponseBody
}

type GetInstancesByIdListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstancesByIdListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstancesByIdListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesByIdListResponse) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstancesByIdListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstancesByIdListResponse) GetBody() *GetInstancesByIdListResponseBody {
	return s.Body
}

func (s *GetInstancesByIdListResponse) SetHeaders(v map[string]*string) *GetInstancesByIdListResponse {
	s.Headers = v
	return s
}

func (s *GetInstancesByIdListResponse) SetStatusCode(v int32) *GetInstancesByIdListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstancesByIdListResponse) SetBody(v *GetInstancesByIdListResponseBody) *GetInstancesByIdListResponse {
	s.Body = v
	return s
}

func (s *GetInstancesByIdListResponse) Validate() error {
	return dara.Validate(s)
}
