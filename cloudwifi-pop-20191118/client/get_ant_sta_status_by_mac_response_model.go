// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAntStaStatusByMacResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAntStaStatusByMacResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAntStaStatusByMacResponse
	GetStatusCode() *int32
	SetBody(v *GetAntStaStatusByMacResponseBody) *GetAntStaStatusByMacResponse
	GetBody() *GetAntStaStatusByMacResponseBody
}

type GetAntStaStatusByMacResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAntStaStatusByMacResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAntStaStatusByMacResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAntStaStatusByMacResponse) GoString() string {
	return s.String()
}

func (s *GetAntStaStatusByMacResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAntStaStatusByMacResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAntStaStatusByMacResponse) GetBody() *GetAntStaStatusByMacResponseBody {
	return s.Body
}

func (s *GetAntStaStatusByMacResponse) SetHeaders(v map[string]*string) *GetAntStaStatusByMacResponse {
	s.Headers = v
	return s
}

func (s *GetAntStaStatusByMacResponse) SetStatusCode(v int32) *GetAntStaStatusByMacResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAntStaStatusByMacResponse) SetBody(v *GetAntStaStatusByMacResponseBody) *GetAntStaStatusByMacResponse {
	s.Body = v
	return s
}

func (s *GetAntStaStatusByMacResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
