// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLdpsResourceCostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLdpsResourceCostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLdpsResourceCostResponse
	GetStatusCode() *int32
	SetBody(v *GetLdpsResourceCostResponseBody) *GetLdpsResourceCostResponse
	GetBody() *GetLdpsResourceCostResponseBody
}

type GetLdpsResourceCostResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLdpsResourceCostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLdpsResourceCostResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLdpsResourceCostResponse) GoString() string {
	return s.String()
}

func (s *GetLdpsResourceCostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLdpsResourceCostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLdpsResourceCostResponse) GetBody() *GetLdpsResourceCostResponseBody {
	return s.Body
}

func (s *GetLdpsResourceCostResponse) SetHeaders(v map[string]*string) *GetLdpsResourceCostResponse {
	s.Headers = v
	return s
}

func (s *GetLdpsResourceCostResponse) SetStatusCode(v int32) *GetLdpsResourceCostResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLdpsResourceCostResponse) SetBody(v *GetLdpsResourceCostResponseBody) *GetLdpsResourceCostResponse {
	s.Body = v
	return s
}

func (s *GetLdpsResourceCostResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
