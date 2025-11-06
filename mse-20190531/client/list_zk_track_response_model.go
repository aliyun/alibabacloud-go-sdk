// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListZkTrackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListZkTrackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListZkTrackResponse
	GetStatusCode() *int32
	SetBody(v *ListZkTrackResponseBody) *ListZkTrackResponse
	GetBody() *ListZkTrackResponseBody
}

type ListZkTrackResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListZkTrackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListZkTrackResponse) String() string {
	return dara.Prettify(s)
}

func (s ListZkTrackResponse) GoString() string {
	return s.String()
}

func (s *ListZkTrackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListZkTrackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListZkTrackResponse) GetBody() *ListZkTrackResponseBody {
	return s.Body
}

func (s *ListZkTrackResponse) SetHeaders(v map[string]*string) *ListZkTrackResponse {
	s.Headers = v
	return s
}

func (s *ListZkTrackResponse) SetStatusCode(v int32) *ListZkTrackResponse {
	s.StatusCode = &v
	return s
}

func (s *ListZkTrackResponse) SetBody(v *ListZkTrackResponseBody) *ListZkTrackResponse {
	s.Body = v
	return s
}

func (s *ListZkTrackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
