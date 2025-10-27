// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOpenLogShipperResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyOpenLogShipperResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyOpenLogShipperResponse
	GetStatusCode() *int32
	SetBody(v *ModifyOpenLogShipperResponseBody) *ModifyOpenLogShipperResponse
	GetBody() *ModifyOpenLogShipperResponseBody
}

type ModifyOpenLogShipperResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyOpenLogShipperResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOpenLogShipperResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyOpenLogShipperResponse) GoString() string {
	return s.String()
}

func (s *ModifyOpenLogShipperResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyOpenLogShipperResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyOpenLogShipperResponse) GetBody() *ModifyOpenLogShipperResponseBody {
	return s.Body
}

func (s *ModifyOpenLogShipperResponse) SetHeaders(v map[string]*string) *ModifyOpenLogShipperResponse {
	s.Headers = v
	return s
}

func (s *ModifyOpenLogShipperResponse) SetStatusCode(v int32) *ModifyOpenLogShipperResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyOpenLogShipperResponse) SetBody(v *ModifyOpenLogShipperResponseBody) *ModifyOpenLogShipperResponse {
	s.Body = v
	return s
}

func (s *ModifyOpenLogShipperResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
