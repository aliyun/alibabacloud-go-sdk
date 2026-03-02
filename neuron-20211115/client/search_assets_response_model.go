// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAssetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchAssetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchAssetsResponse
	GetStatusCode() *int32
	SetBody(v []*DTO) *SearchAssetsResponse
	GetBody() []*DTO
}

type SearchAssetsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*DTO             `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s SearchAssetsResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchAssetsResponse) GoString() string {
	return s.String()
}

func (s *SearchAssetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchAssetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchAssetsResponse) GetBody() []*DTO {
	return s.Body
}

func (s *SearchAssetsResponse) SetHeaders(v map[string]*string) *SearchAssetsResponse {
	s.Headers = v
	return s
}

func (s *SearchAssetsResponse) SetStatusCode(v int32) *SearchAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAssetsResponse) SetBody(v []*DTO) *SearchAssetsResponse {
	s.Body = v
	return s
}

func (s *SearchAssetsResponse) Validate() error {
	if s.Body != nil {
		for _, item := range s.Body {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
