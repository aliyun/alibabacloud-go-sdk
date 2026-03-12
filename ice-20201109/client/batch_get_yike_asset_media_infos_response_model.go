// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetYikeAssetMediaInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchGetYikeAssetMediaInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchGetYikeAssetMediaInfosResponse
	GetStatusCode() *int32
	SetBody(v *BatchGetYikeAssetMediaInfosResponseBody) *BatchGetYikeAssetMediaInfosResponse
	GetBody() *BatchGetYikeAssetMediaInfosResponseBody
}

type BatchGetYikeAssetMediaInfosResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetYikeAssetMediaInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetYikeAssetMediaInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchGetYikeAssetMediaInfosResponse) GoString() string {
	return s.String()
}

func (s *BatchGetYikeAssetMediaInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchGetYikeAssetMediaInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchGetYikeAssetMediaInfosResponse) GetBody() *BatchGetYikeAssetMediaInfosResponseBody {
	return s.Body
}

func (s *BatchGetYikeAssetMediaInfosResponse) SetHeaders(v map[string]*string) *BatchGetYikeAssetMediaInfosResponse {
	s.Headers = v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponse) SetStatusCode(v int32) *BatchGetYikeAssetMediaInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponse) SetBody(v *BatchGetYikeAssetMediaInfosResponseBody) *BatchGetYikeAssetMediaInfosResponse {
	s.Body = v
	return s
}

func (s *BatchGetYikeAssetMediaInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
