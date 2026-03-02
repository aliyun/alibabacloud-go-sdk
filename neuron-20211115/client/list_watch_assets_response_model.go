// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWatchAssetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWatchAssetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWatchAssetsResponse
	GetStatusCode() *int32
	SetBody(v []*DTO) *ListWatchAssetsResponse
	GetBody() []*DTO
}

type ListWatchAssetsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*DTO             `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ListWatchAssetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWatchAssetsResponse) GoString() string {
	return s.String()
}

func (s *ListWatchAssetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWatchAssetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWatchAssetsResponse) GetBody() []*DTO {
	return s.Body
}

func (s *ListWatchAssetsResponse) SetHeaders(v map[string]*string) *ListWatchAssetsResponse {
	s.Headers = v
	return s
}

func (s *ListWatchAssetsResponse) SetStatusCode(v int32) *ListWatchAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWatchAssetsResponse) SetBody(v []*DTO) *ListWatchAssetsResponse {
	s.Body = v
	return s
}

func (s *ListWatchAssetsResponse) Validate() error {
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
