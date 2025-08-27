// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCooperatorSyncPayStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CooperatorSyncPayStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CooperatorSyncPayStatusResponse
	GetStatusCode() *int32
	SetBody(v *CooperatorSyncPayStatusResponseBody) *CooperatorSyncPayStatusResponse
	GetBody() *CooperatorSyncPayStatusResponseBody
}

type CooperatorSyncPayStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CooperatorSyncPayStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CooperatorSyncPayStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s CooperatorSyncPayStatusResponse) GoString() string {
	return s.String()
}

func (s *CooperatorSyncPayStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CooperatorSyncPayStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CooperatorSyncPayStatusResponse) GetBody() *CooperatorSyncPayStatusResponseBody {
	return s.Body
}

func (s *CooperatorSyncPayStatusResponse) SetHeaders(v map[string]*string) *CooperatorSyncPayStatusResponse {
	s.Headers = v
	return s
}

func (s *CooperatorSyncPayStatusResponse) SetStatusCode(v int32) *CooperatorSyncPayStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CooperatorSyncPayStatusResponse) SetBody(v *CooperatorSyncPayStatusResponseBody) *CooperatorSyncPayStatusResponse {
	s.Body = v
	return s
}

func (s *CooperatorSyncPayStatusResponse) Validate() error {
	return dara.Validate(s)
}
