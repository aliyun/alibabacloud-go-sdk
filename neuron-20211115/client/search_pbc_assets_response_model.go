// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchPbcAssetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchPbcAssetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchPbcAssetsResponse
	GetStatusCode() *int32
	SetBody(v *SearchPbcAssetsResponseBody) *SearchPbcAssetsResponse
	GetBody() *SearchPbcAssetsResponseBody
}

type SearchPbcAssetsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchPbcAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchPbcAssetsResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchPbcAssetsResponse) GoString() string {
	return s.String()
}

func (s *SearchPbcAssetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchPbcAssetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchPbcAssetsResponse) GetBody() *SearchPbcAssetsResponseBody {
	return s.Body
}

func (s *SearchPbcAssetsResponse) SetHeaders(v map[string]*string) *SearchPbcAssetsResponse {
	s.Headers = v
	return s
}

func (s *SearchPbcAssetsResponse) SetStatusCode(v int32) *SearchPbcAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchPbcAssetsResponse) SetBody(v *SearchPbcAssetsResponseBody) *SearchPbcAssetsResponse {
	s.Body = v
	return s
}

func (s *SearchPbcAssetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
