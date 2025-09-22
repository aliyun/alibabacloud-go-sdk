// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInventoryListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInventoryListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInventoryListResponse
	GetStatusCode() *int32
	SetBody(v *GetInventoryListResponseBody) *GetInventoryListResponse
	GetBody() *GetInventoryListResponseBody
}

type GetInventoryListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInventoryListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInventoryListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInventoryListResponse) GoString() string {
	return s.String()
}

func (s *GetInventoryListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInventoryListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInventoryListResponse) GetBody() *GetInventoryListResponseBody {
	return s.Body
}

func (s *GetInventoryListResponse) SetHeaders(v map[string]*string) *GetInventoryListResponse {
	s.Headers = v
	return s
}

func (s *GetInventoryListResponse) SetStatusCode(v int32) *GetInventoryListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInventoryListResponse) SetBody(v *GetInventoryListResponseBody) *GetInventoryListResponse {
	s.Body = v
	return s
}

func (s *GetInventoryListResponse) Validate() error {
	return dara.Validate(s)
}
