// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataArchiveOrderDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataArchiveOrderDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataArchiveOrderDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetDataArchiveOrderDetailResponseBody) *GetDataArchiveOrderDetailResponse
	GetBody() *GetDataArchiveOrderDetailResponseBody
}

type GetDataArchiveOrderDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataArchiveOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataArchiveOrderDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataArchiveOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDataArchiveOrderDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataArchiveOrderDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataArchiveOrderDetailResponse) GetBody() *GetDataArchiveOrderDetailResponseBody {
	return s.Body
}

func (s *GetDataArchiveOrderDetailResponse) SetHeaders(v map[string]*string) *GetDataArchiveOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDataArchiveOrderDetailResponse) SetStatusCode(v int32) *GetDataArchiveOrderDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataArchiveOrderDetailResponse) SetBody(v *GetDataArchiveOrderDetailResponseBody) *GetDataArchiveOrderDetailResponse {
	s.Body = v
	return s
}

func (s *GetDataArchiveOrderDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
