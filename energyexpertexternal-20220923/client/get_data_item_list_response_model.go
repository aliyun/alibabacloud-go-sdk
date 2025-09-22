// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataItemListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataItemListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataItemListResponse
	GetStatusCode() *int32
	SetBody(v *GetDataItemListResponseBody) *GetDataItemListResponse
	GetBody() *GetDataItemListResponseBody
}

type GetDataItemListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataItemListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataItemListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataItemListResponse) GoString() string {
	return s.String()
}

func (s *GetDataItemListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataItemListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataItemListResponse) GetBody() *GetDataItemListResponseBody {
	return s.Body
}

func (s *GetDataItemListResponse) SetHeaders(v map[string]*string) *GetDataItemListResponse {
	s.Headers = v
	return s
}

func (s *GetDataItemListResponse) SetStatusCode(v int32) *GetDataItemListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataItemListResponse) SetBody(v *GetDataItemListResponseBody) *GetDataItemListResponse {
	s.Body = v
	return s
}

func (s *GetDataItemListResponse) Validate() error {
	return dara.Validate(s)
}
