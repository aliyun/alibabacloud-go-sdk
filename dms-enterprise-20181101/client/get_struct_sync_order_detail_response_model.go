// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStructSyncOrderDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStructSyncOrderDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStructSyncOrderDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetStructSyncOrderDetailResponseBody) *GetStructSyncOrderDetailResponse
	GetBody() *GetStructSyncOrderDetailResponseBody
}

type GetStructSyncOrderDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStructSyncOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStructSyncOrderDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *GetStructSyncOrderDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStructSyncOrderDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStructSyncOrderDetailResponse) GetBody() *GetStructSyncOrderDetailResponseBody {
	return s.Body
}

func (s *GetStructSyncOrderDetailResponse) SetHeaders(v map[string]*string) *GetStructSyncOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *GetStructSyncOrderDetailResponse) SetStatusCode(v int32) *GetStructSyncOrderDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStructSyncOrderDetailResponse) SetBody(v *GetStructSyncOrderDetailResponseBody) *GetStructSyncOrderDetailResponse {
	s.Body = v
	return s
}

func (s *GetStructSyncOrderDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
