// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnTagDataAssetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnTagDataAssetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnTagDataAssetsResponse
	GetStatusCode() *int32
	SetBody(v *UnTagDataAssetsResponseBody) *UnTagDataAssetsResponse
	GetBody() *UnTagDataAssetsResponseBody
}

type UnTagDataAssetsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnTagDataAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnTagDataAssetsResponse) String() string {
	return dara.Prettify(s)
}

func (s UnTagDataAssetsResponse) GoString() string {
	return s.String()
}

func (s *UnTagDataAssetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnTagDataAssetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnTagDataAssetsResponse) GetBody() *UnTagDataAssetsResponseBody {
	return s.Body
}

func (s *UnTagDataAssetsResponse) SetHeaders(v map[string]*string) *UnTagDataAssetsResponse {
	s.Headers = v
	return s
}

func (s *UnTagDataAssetsResponse) SetStatusCode(v int32) *UnTagDataAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *UnTagDataAssetsResponse) SetBody(v *UnTagDataAssetsResponseBody) *UnTagDataAssetsResponse {
	s.Body = v
	return s
}

func (s *UnTagDataAssetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
