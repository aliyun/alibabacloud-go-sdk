// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStaDetailedStatusByMacResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStaDetailedStatusByMacResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStaDetailedStatusByMacResponse
	GetStatusCode() *int32
	SetBody(v *GetStaDetailedStatusByMacResponseBody) *GetStaDetailedStatusByMacResponse
	GetBody() *GetStaDetailedStatusByMacResponseBody
}

type GetStaDetailedStatusByMacResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStaDetailedStatusByMacResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStaDetailedStatusByMacResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStaDetailedStatusByMacResponse) GoString() string {
	return s.String()
}

func (s *GetStaDetailedStatusByMacResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStaDetailedStatusByMacResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStaDetailedStatusByMacResponse) GetBody() *GetStaDetailedStatusByMacResponseBody {
	return s.Body
}

func (s *GetStaDetailedStatusByMacResponse) SetHeaders(v map[string]*string) *GetStaDetailedStatusByMacResponse {
	s.Headers = v
	return s
}

func (s *GetStaDetailedStatusByMacResponse) SetStatusCode(v int32) *GetStaDetailedStatusByMacResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStaDetailedStatusByMacResponse) SetBody(v *GetStaDetailedStatusByMacResponseBody) *GetStaDetailedStatusByMacResponse {
	s.Body = v
	return s
}

func (s *GetStaDetailedStatusByMacResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
