// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySlsLogStoreListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySlsLogStoreListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySlsLogStoreListResponse
	GetStatusCode() *int32
	SetBody(v *QuerySlsLogStoreListResponseBody) *QuerySlsLogStoreListResponse
	GetBody() *QuerySlsLogStoreListResponseBody
}

type QuerySlsLogStoreListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySlsLogStoreListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySlsLogStoreListResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySlsLogStoreListResponse) GoString() string {
	return s.String()
}

func (s *QuerySlsLogStoreListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySlsLogStoreListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySlsLogStoreListResponse) GetBody() *QuerySlsLogStoreListResponseBody {
	return s.Body
}

func (s *QuerySlsLogStoreListResponse) SetHeaders(v map[string]*string) *QuerySlsLogStoreListResponse {
	s.Headers = v
	return s
}

func (s *QuerySlsLogStoreListResponse) SetStatusCode(v int32) *QuerySlsLogStoreListResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySlsLogStoreListResponse) SetBody(v *QuerySlsLogStoreListResponseBody) *QuerySlsLogStoreListResponse {
	s.Body = v
	return s
}

func (s *QuerySlsLogStoreListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
