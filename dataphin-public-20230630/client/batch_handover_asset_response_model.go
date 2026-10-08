// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchHandoverAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchHandoverAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchHandoverAssetResponse
	GetStatusCode() *int32
	SetBody(v *BatchHandoverAssetResponseBody) *BatchHandoverAssetResponse
	GetBody() *BatchHandoverAssetResponseBody
}

type BatchHandoverAssetResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchHandoverAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchHandoverAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchHandoverAssetResponse) GoString() string {
	return s.String()
}

func (s *BatchHandoverAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchHandoverAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchHandoverAssetResponse) GetBody() *BatchHandoverAssetResponseBody {
	return s.Body
}

func (s *BatchHandoverAssetResponse) SetHeaders(v map[string]*string) *BatchHandoverAssetResponse {
	s.Headers = v
	return s
}

func (s *BatchHandoverAssetResponse) SetStatusCode(v int32) *BatchHandoverAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchHandoverAssetResponse) SetBody(v *BatchHandoverAssetResponseBody) *BatchHandoverAssetResponse {
	s.Body = v
	return s
}

func (s *BatchHandoverAssetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
