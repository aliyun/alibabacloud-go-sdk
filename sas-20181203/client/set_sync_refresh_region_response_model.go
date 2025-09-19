// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSyncRefreshRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetSyncRefreshRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetSyncRefreshRegionResponse
	GetStatusCode() *int32
	SetBody(v *SetSyncRefreshRegionResponseBody) *SetSyncRefreshRegionResponse
	GetBody() *SetSyncRefreshRegionResponseBody
}

type SetSyncRefreshRegionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSyncRefreshRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSyncRefreshRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s SetSyncRefreshRegionResponse) GoString() string {
	return s.String()
}

func (s *SetSyncRefreshRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetSyncRefreshRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetSyncRefreshRegionResponse) GetBody() *SetSyncRefreshRegionResponseBody {
	return s.Body
}

func (s *SetSyncRefreshRegionResponse) SetHeaders(v map[string]*string) *SetSyncRefreshRegionResponse {
	s.Headers = v
	return s
}

func (s *SetSyncRefreshRegionResponse) SetStatusCode(v int32) *SetSyncRefreshRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSyncRefreshRegionResponse) SetBody(v *SetSyncRefreshRegionResponseBody) *SetSyncRefreshRegionResponse {
	s.Body = v
	return s
}

func (s *SetSyncRefreshRegionResponse) Validate() error {
	return dara.Validate(s)
}
