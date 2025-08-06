// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncUserProdAccountAndBucketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncUserProdAccountAndBucketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncUserProdAccountAndBucketResponse
	GetStatusCode() *int32
	SetBody(v *SyncUserProdAccountAndBucketResponseBody) *SyncUserProdAccountAndBucketResponse
	GetBody() *SyncUserProdAccountAndBucketResponseBody
}

type SyncUserProdAccountAndBucketResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncUserProdAccountAndBucketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncUserProdAccountAndBucketResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncUserProdAccountAndBucketResponse) GoString() string {
	return s.String()
}

func (s *SyncUserProdAccountAndBucketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncUserProdAccountAndBucketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncUserProdAccountAndBucketResponse) GetBody() *SyncUserProdAccountAndBucketResponseBody {
	return s.Body
}

func (s *SyncUserProdAccountAndBucketResponse) SetHeaders(v map[string]*string) *SyncUserProdAccountAndBucketResponse {
	s.Headers = v
	return s
}

func (s *SyncUserProdAccountAndBucketResponse) SetStatusCode(v int32) *SyncUserProdAccountAndBucketResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncUserProdAccountAndBucketResponse) SetBody(v *SyncUserProdAccountAndBucketResponseBody) *SyncUserProdAccountAndBucketResponse {
	s.Body = v
	return s
}

func (s *SyncUserProdAccountAndBucketResponse) Validate() error {
	return dara.Validate(s)
}
