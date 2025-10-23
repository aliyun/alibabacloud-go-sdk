// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGwpInventoryConstituteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGwpInventoryConstituteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGwpInventoryConstituteResponse
	GetStatusCode() *int32
	SetBody(v *GetGwpInventoryConstituteResponseBody) *GetGwpInventoryConstituteResponse
	GetBody() *GetGwpInventoryConstituteResponseBody
}

type GetGwpInventoryConstituteResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGwpInventoryConstituteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGwpInventoryConstituteResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGwpInventoryConstituteResponse) GoString() string {
	return s.String()
}

func (s *GetGwpInventoryConstituteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGwpInventoryConstituteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGwpInventoryConstituteResponse) GetBody() *GetGwpInventoryConstituteResponseBody {
	return s.Body
}

func (s *GetGwpInventoryConstituteResponse) SetHeaders(v map[string]*string) *GetGwpInventoryConstituteResponse {
	s.Headers = v
	return s
}

func (s *GetGwpInventoryConstituteResponse) SetStatusCode(v int32) *GetGwpInventoryConstituteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGwpInventoryConstituteResponse) SetBody(v *GetGwpInventoryConstituteResponseBody) *GetGwpInventoryConstituteResponse {
	s.Body = v
	return s
}

func (s *GetGwpInventoryConstituteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
