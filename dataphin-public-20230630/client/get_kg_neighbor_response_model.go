// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKgNeighborResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKgNeighborResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKgNeighborResponse
	GetStatusCode() *int32
	SetBody(v *GetKgNeighborResponseBody) *GetKgNeighborResponse
	GetBody() *GetKgNeighborResponseBody
}

type GetKgNeighborResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKgNeighborResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKgNeighborResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKgNeighborResponse) GoString() string {
	return s.String()
}

func (s *GetKgNeighborResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKgNeighborResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKgNeighborResponse) GetBody() *GetKgNeighborResponseBody {
	return s.Body
}

func (s *GetKgNeighborResponse) SetHeaders(v map[string]*string) *GetKgNeighborResponse {
	s.Headers = v
	return s
}

func (s *GetKgNeighborResponse) SetStatusCode(v int32) *GetKgNeighborResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKgNeighborResponse) SetBody(v *GetKgNeighborResponseBody) *GetKgNeighborResponse {
	s.Body = v
	return s
}

func (s *GetKgNeighborResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
