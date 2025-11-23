// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStructSyncExecSqlDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStructSyncExecSqlDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStructSyncExecSqlDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetStructSyncExecSqlDetailResponseBody) *GetStructSyncExecSqlDetailResponse
	GetBody() *GetStructSyncExecSqlDetailResponseBody
}

type GetStructSyncExecSqlDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStructSyncExecSqlDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStructSyncExecSqlDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncExecSqlDetailResponse) GoString() string {
	return s.String()
}

func (s *GetStructSyncExecSqlDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStructSyncExecSqlDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStructSyncExecSqlDetailResponse) GetBody() *GetStructSyncExecSqlDetailResponseBody {
	return s.Body
}

func (s *GetStructSyncExecSqlDetailResponse) SetHeaders(v map[string]*string) *GetStructSyncExecSqlDetailResponse {
	s.Headers = v
	return s
}

func (s *GetStructSyncExecSqlDetailResponse) SetStatusCode(v int32) *GetStructSyncExecSqlDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStructSyncExecSqlDetailResponse) SetBody(v *GetStructSyncExecSqlDetailResponseBody) *GetStructSyncExecSqlDetailResponse {
	s.Body = v
	return s
}

func (s *GetStructSyncExecSqlDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
