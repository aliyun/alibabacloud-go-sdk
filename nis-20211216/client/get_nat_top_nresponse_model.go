// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNatTopNResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNatTopNResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNatTopNResponse
	GetStatusCode() *int32
	SetBody(v *GetNatTopNResponseBody) *GetNatTopNResponse
	GetBody() *GetNatTopNResponseBody
}

type GetNatTopNResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNatTopNResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNatTopNResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNatTopNResponse) GoString() string {
	return s.String()
}

func (s *GetNatTopNResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNatTopNResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNatTopNResponse) GetBody() *GetNatTopNResponseBody {
	return s.Body
}

func (s *GetNatTopNResponse) SetHeaders(v map[string]*string) *GetNatTopNResponse {
	s.Headers = v
	return s
}

func (s *GetNatTopNResponse) SetStatusCode(v int32) *GetNatTopNResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNatTopNResponse) SetBody(v *GetNatTopNResponseBody) *GetNatTopNResponse {
	s.Body = v
	return s
}

func (s *GetNatTopNResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
