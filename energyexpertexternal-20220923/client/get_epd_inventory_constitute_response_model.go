// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEpdInventoryConstituteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEpdInventoryConstituteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEpdInventoryConstituteResponse
	GetStatusCode() *int32
	SetBody(v *GetEpdInventoryConstituteResponseBody) *GetEpdInventoryConstituteResponse
	GetBody() *GetEpdInventoryConstituteResponseBody
}

type GetEpdInventoryConstituteResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEpdInventoryConstituteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEpdInventoryConstituteResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEpdInventoryConstituteResponse) GoString() string {
	return s.String()
}

func (s *GetEpdInventoryConstituteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEpdInventoryConstituteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEpdInventoryConstituteResponse) GetBody() *GetEpdInventoryConstituteResponseBody {
	return s.Body
}

func (s *GetEpdInventoryConstituteResponse) SetHeaders(v map[string]*string) *GetEpdInventoryConstituteResponse {
	s.Headers = v
	return s
}

func (s *GetEpdInventoryConstituteResponse) SetStatusCode(v int32) *GetEpdInventoryConstituteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEpdInventoryConstituteResponse) SetBody(v *GetEpdInventoryConstituteResponseBody) *GetEpdInventoryConstituteResponse {
	s.Body = v
	return s
}

func (s *GetEpdInventoryConstituteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
