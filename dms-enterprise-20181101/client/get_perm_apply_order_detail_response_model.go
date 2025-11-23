// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPermApplyOrderDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPermApplyOrderDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPermApplyOrderDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetPermApplyOrderDetailResponseBody) *GetPermApplyOrderDetailResponse
	GetBody() *GetPermApplyOrderDetailResponseBody
}

type GetPermApplyOrderDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPermApplyOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPermApplyOrderDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPermApplyOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *GetPermApplyOrderDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPermApplyOrderDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPermApplyOrderDetailResponse) GetBody() *GetPermApplyOrderDetailResponseBody {
	return s.Body
}

func (s *GetPermApplyOrderDetailResponse) SetHeaders(v map[string]*string) *GetPermApplyOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *GetPermApplyOrderDetailResponse) SetStatusCode(v int32) *GetPermApplyOrderDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPermApplyOrderDetailResponse) SetBody(v *GetPermApplyOrderDetailResponseBody) *GetPermApplyOrderDetailResponse {
	s.Body = v
	return s
}

func (s *GetPermApplyOrderDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
