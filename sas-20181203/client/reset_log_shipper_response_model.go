// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetLogShipperResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetLogShipperResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetLogShipperResponse
	GetStatusCode() *int32
	SetBody(v *ResetLogShipperResponseBody) *ResetLogShipperResponse
	GetBody() *ResetLogShipperResponseBody
}

type ResetLogShipperResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetLogShipperResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetLogShipperResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetLogShipperResponse) GoString() string {
	return s.String()
}

func (s *ResetLogShipperResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetLogShipperResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetLogShipperResponse) GetBody() *ResetLogShipperResponseBody {
	return s.Body
}

func (s *ResetLogShipperResponse) SetHeaders(v map[string]*string) *ResetLogShipperResponse {
	s.Headers = v
	return s
}

func (s *ResetLogShipperResponse) SetStatusCode(v int32) *ResetLogShipperResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetLogShipperResponse) SetBody(v *ResetLogShipperResponseBody) *ResetLogShipperResponse {
	s.Body = v
	return s
}

func (s *ResetLogShipperResponse) Validate() error {
	return dara.Validate(s)
}
