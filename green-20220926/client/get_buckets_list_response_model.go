// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBucketsListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBucketsListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBucketsListResponse
	GetStatusCode() *int32
	SetBody(v *GetBucketsListResponseBody) *GetBucketsListResponse
	GetBody() *GetBucketsListResponseBody
}

type GetBucketsListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBucketsListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBucketsListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBucketsListResponse) GoString() string {
	return s.String()
}

func (s *GetBucketsListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBucketsListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBucketsListResponse) GetBody() *GetBucketsListResponseBody {
	return s.Body
}

func (s *GetBucketsListResponse) SetHeaders(v map[string]*string) *GetBucketsListResponse {
	s.Headers = v
	return s
}

func (s *GetBucketsListResponse) SetStatusCode(v int32) *GetBucketsListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketsListResponse) SetBody(v *GetBucketsListResponseBody) *GetBucketsListResponse {
	s.Body = v
	return s
}

func (s *GetBucketsListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
