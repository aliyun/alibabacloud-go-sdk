// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppNetworkResponse
	GetStatusCode() *int32
	SetBody(v *GetAppNetworkResponseBody) *GetAppNetworkResponse
	GetBody() *GetAppNetworkResponseBody
}

type GetAppNetworkResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppNetworkResponse) GoString() string {
	return s.String()
}

func (s *GetAppNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppNetworkResponse) GetBody() *GetAppNetworkResponseBody {
	return s.Body
}

func (s *GetAppNetworkResponse) SetHeaders(v map[string]*string) *GetAppNetworkResponse {
	s.Headers = v
	return s
}

func (s *GetAppNetworkResponse) SetStatusCode(v int32) *GetAppNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppNetworkResponse) SetBody(v *GetAppNetworkResponseBody) *GetAppNetworkResponse {
	s.Body = v
	return s
}

func (s *GetAppNetworkResponse) Validate() error {
	return dara.Validate(s)
}
