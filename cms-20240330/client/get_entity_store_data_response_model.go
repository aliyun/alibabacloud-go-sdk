// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEntityStoreDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEntityStoreDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEntityStoreDataResponse
	GetStatusCode() *int32
	SetBody(v *GetEntityStoreDataResponseBody) *GetEntityStoreDataResponse
	GetBody() *GetEntityStoreDataResponseBody
}

type GetEntityStoreDataResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEntityStoreDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEntityStoreDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEntityStoreDataResponse) GoString() string {
	return s.String()
}

func (s *GetEntityStoreDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEntityStoreDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEntityStoreDataResponse) GetBody() *GetEntityStoreDataResponseBody {
	return s.Body
}

func (s *GetEntityStoreDataResponse) SetHeaders(v map[string]*string) *GetEntityStoreDataResponse {
	s.Headers = v
	return s
}

func (s *GetEntityStoreDataResponse) SetStatusCode(v int32) *GetEntityStoreDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEntityStoreDataResponse) SetBody(v *GetEntityStoreDataResponseBody) *GetEntityStoreDataResponse {
	s.Body = v
	return s
}

func (s *GetEntityStoreDataResponse) Validate() error {
	return dara.Validate(s)
}
