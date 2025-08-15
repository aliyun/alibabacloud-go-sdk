// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshUserSyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshUserSyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshUserSyncResponse
	GetStatusCode() *int32
}

type RefreshUserSyncResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s RefreshUserSyncResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshUserSyncResponse) GoString() string {
	return s.String()
}

func (s *RefreshUserSyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshUserSyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshUserSyncResponse) SetHeaders(v map[string]*string) *RefreshUserSyncResponse {
	s.Headers = v
	return s
}

func (s *RefreshUserSyncResponse) SetStatusCode(v int32) *RefreshUserSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshUserSyncResponse) Validate() error {
	return dara.Validate(s)
}
