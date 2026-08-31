// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetTopicsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAssetTopicsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAssetTopicsResponse
	GetStatusCode() *int32
	SetBody(v *ListAssetTopicsResponseBody) *ListAssetTopicsResponse
	GetBody() *ListAssetTopicsResponseBody
}

type ListAssetTopicsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAssetTopicsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAssetTopicsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAssetTopicsResponse) GoString() string {
	return s.String()
}

func (s *ListAssetTopicsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAssetTopicsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAssetTopicsResponse) GetBody() *ListAssetTopicsResponseBody {
	return s.Body
}

func (s *ListAssetTopicsResponse) SetHeaders(v map[string]*string) *ListAssetTopicsResponse {
	s.Headers = v
	return s
}

func (s *ListAssetTopicsResponse) SetStatusCode(v int32) *ListAssetTopicsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAssetTopicsResponse) SetBody(v *ListAssetTopicsResponseBody) *ListAssetTopicsResponse {
	s.Body = v
	return s
}

func (s *ListAssetTopicsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
