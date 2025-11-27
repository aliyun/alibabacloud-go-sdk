// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySyncStatusByAliUidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySyncStatusByAliUidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySyncStatusByAliUidResponse
	GetStatusCode() *int32
	SetBody(v *QuerySyncStatusByAliUidResponseBody) *QuerySyncStatusByAliUidResponse
	GetBody() *QuerySyncStatusByAliUidResponseBody
}

type QuerySyncStatusByAliUidResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySyncStatusByAliUidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySyncStatusByAliUidResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySyncStatusByAliUidResponse) GoString() string {
	return s.String()
}

func (s *QuerySyncStatusByAliUidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySyncStatusByAliUidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySyncStatusByAliUidResponse) GetBody() *QuerySyncStatusByAliUidResponseBody {
	return s.Body
}

func (s *QuerySyncStatusByAliUidResponse) SetHeaders(v map[string]*string) *QuerySyncStatusByAliUidResponse {
	s.Headers = v
	return s
}

func (s *QuerySyncStatusByAliUidResponse) SetStatusCode(v int32) *QuerySyncStatusByAliUidResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponse) SetBody(v *QuerySyncStatusByAliUidResponseBody) *QuerySyncStatusByAliUidResponse {
	s.Body = v
	return s
}

func (s *QuerySyncStatusByAliUidResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
