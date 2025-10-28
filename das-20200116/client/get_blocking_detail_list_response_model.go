// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBlockingDetailListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBlockingDetailListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBlockingDetailListResponse
	GetStatusCode() *int32
	SetBody(v *GetBlockingDetailListResponseBody) *GetBlockingDetailListResponse
	GetBody() *GetBlockingDetailListResponseBody
}

type GetBlockingDetailListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBlockingDetailListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBlockingDetailListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBlockingDetailListResponse) GoString() string {
	return s.String()
}

func (s *GetBlockingDetailListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBlockingDetailListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBlockingDetailListResponse) GetBody() *GetBlockingDetailListResponseBody {
	return s.Body
}

func (s *GetBlockingDetailListResponse) SetHeaders(v map[string]*string) *GetBlockingDetailListResponse {
	s.Headers = v
	return s
}

func (s *GetBlockingDetailListResponse) SetStatusCode(v int32) *GetBlockingDetailListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBlockingDetailListResponse) SetBody(v *GetBlockingDetailListResponseBody) *GetBlockingDetailListResponse {
	s.Body = v
	return s
}

func (s *GetBlockingDetailListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
