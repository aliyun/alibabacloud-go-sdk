// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPersonalStorageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPersonalStorageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPersonalStorageListResponse
	GetStatusCode() *int32
	SetBody(v *GetPersonalStorageListResponseBody) *GetPersonalStorageListResponse
	GetBody() *GetPersonalStorageListResponseBody
}

type GetPersonalStorageListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPersonalStorageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPersonalStorageListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPersonalStorageListResponse) GoString() string {
	return s.String()
}

func (s *GetPersonalStorageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPersonalStorageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPersonalStorageListResponse) GetBody() *GetPersonalStorageListResponseBody {
	return s.Body
}

func (s *GetPersonalStorageListResponse) SetHeaders(v map[string]*string) *GetPersonalStorageListResponse {
	s.Headers = v
	return s
}

func (s *GetPersonalStorageListResponse) SetStatusCode(v int32) *GetPersonalStorageListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPersonalStorageListResponse) SetBody(v *GetPersonalStorageListResponseBody) *GetPersonalStorageListResponse {
	s.Body = v
	return s
}

func (s *GetPersonalStorageListResponse) Validate() error {
	return dara.Validate(s)
}
