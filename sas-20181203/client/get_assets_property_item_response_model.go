// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetsPropertyItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAssetsPropertyItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAssetsPropertyItemResponse
	GetStatusCode() *int32
	SetBody(v *GetAssetsPropertyItemResponseBody) *GetAssetsPropertyItemResponse
	GetBody() *GetAssetsPropertyItemResponseBody
}

type GetAssetsPropertyItemResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAssetsPropertyItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAssetsPropertyItemResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAssetsPropertyItemResponse) GoString() string {
	return s.String()
}

func (s *GetAssetsPropertyItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAssetsPropertyItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAssetsPropertyItemResponse) GetBody() *GetAssetsPropertyItemResponseBody {
	return s.Body
}

func (s *GetAssetsPropertyItemResponse) SetHeaders(v map[string]*string) *GetAssetsPropertyItemResponse {
	s.Headers = v
	return s
}

func (s *GetAssetsPropertyItemResponse) SetStatusCode(v int32) *GetAssetsPropertyItemResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAssetsPropertyItemResponse) SetBody(v *GetAssetsPropertyItemResponseBody) *GetAssetsPropertyItemResponse {
	s.Body = v
	return s
}

func (s *GetAssetsPropertyItemResponse) Validate() error {
	return dara.Validate(s)
}
