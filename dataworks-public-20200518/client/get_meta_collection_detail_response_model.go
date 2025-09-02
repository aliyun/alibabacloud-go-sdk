// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaCollectionDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaCollectionDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaCollectionDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaCollectionDetailResponseBody) *GetMetaCollectionDetailResponse
	GetBody() *GetMetaCollectionDetailResponseBody
}

type GetMetaCollectionDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaCollectionDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaCollectionDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaCollectionDetailResponse) GoString() string {
	return s.String()
}

func (s *GetMetaCollectionDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaCollectionDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaCollectionDetailResponse) GetBody() *GetMetaCollectionDetailResponseBody {
	return s.Body
}

func (s *GetMetaCollectionDetailResponse) SetHeaders(v map[string]*string) *GetMetaCollectionDetailResponse {
	s.Headers = v
	return s
}

func (s *GetMetaCollectionDetailResponse) SetStatusCode(v int32) *GetMetaCollectionDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaCollectionDetailResponse) SetBody(v *GetMetaCollectionDetailResponseBody) *GetMetaCollectionDetailResponse {
	s.Body = v
	return s
}

func (s *GetMetaCollectionDetailResponse) Validate() error {
	return dara.Validate(s)
}
