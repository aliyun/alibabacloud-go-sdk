// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUrgentDemandItemListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUrgentDemandItemListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUrgentDemandItemListResponse
	GetStatusCode() *int32
	SetBody(v *GetUrgentDemandItemListResponseBody) *GetUrgentDemandItemListResponse
	GetBody() *GetUrgentDemandItemListResponseBody
}

type GetUrgentDemandItemListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUrgentDemandItemListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUrgentDemandItemListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUrgentDemandItemListResponse) GoString() string {
	return s.String()
}

func (s *GetUrgentDemandItemListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUrgentDemandItemListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUrgentDemandItemListResponse) GetBody() *GetUrgentDemandItemListResponseBody {
	return s.Body
}

func (s *GetUrgentDemandItemListResponse) SetHeaders(v map[string]*string) *GetUrgentDemandItemListResponse {
	s.Headers = v
	return s
}

func (s *GetUrgentDemandItemListResponse) SetStatusCode(v int32) *GetUrgentDemandItemListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUrgentDemandItemListResponse) SetBody(v *GetUrgentDemandItemListResponseBody) *GetUrgentDemandItemListResponse {
	s.Body = v
	return s
}

func (s *GetUrgentDemandItemListResponse) Validate() error {
	return dara.Validate(s)
}
