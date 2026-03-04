// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterYikeAssetMediaInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterYikeAssetMediaInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterYikeAssetMediaInfoResponse
	GetStatusCode() *int32
	SetBody(v *RegisterYikeAssetMediaInfoResponseBody) *RegisterYikeAssetMediaInfoResponse
	GetBody() *RegisterYikeAssetMediaInfoResponseBody
}

type RegisterYikeAssetMediaInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterYikeAssetMediaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterYikeAssetMediaInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterYikeAssetMediaInfoResponse) GoString() string {
	return s.String()
}

func (s *RegisterYikeAssetMediaInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterYikeAssetMediaInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterYikeAssetMediaInfoResponse) GetBody() *RegisterYikeAssetMediaInfoResponseBody {
	return s.Body
}

func (s *RegisterYikeAssetMediaInfoResponse) SetHeaders(v map[string]*string) *RegisterYikeAssetMediaInfoResponse {
	s.Headers = v
	return s
}

func (s *RegisterYikeAssetMediaInfoResponse) SetStatusCode(v int32) *RegisterYikeAssetMediaInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterYikeAssetMediaInfoResponse) SetBody(v *RegisterYikeAssetMediaInfoResponseBody) *RegisterYikeAssetMediaInfoResponse {
	s.Body = v
	return s
}

func (s *RegisterYikeAssetMediaInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
