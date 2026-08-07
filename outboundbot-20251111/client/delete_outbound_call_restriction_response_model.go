// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOutboundCallRestrictionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOutboundCallRestrictionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOutboundCallRestrictionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOutboundCallRestrictionResponseBody) *DeleteOutboundCallRestrictionResponse
	GetBody() *DeleteOutboundCallRestrictionResponseBody
}

type DeleteOutboundCallRestrictionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOutboundCallRestrictionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOutboundCallRestrictionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOutboundCallRestrictionResponse) GoString() string {
	return s.String()
}

func (s *DeleteOutboundCallRestrictionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOutboundCallRestrictionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOutboundCallRestrictionResponse) GetBody() *DeleteOutboundCallRestrictionResponseBody {
	return s.Body
}

func (s *DeleteOutboundCallRestrictionResponse) SetHeaders(v map[string]*string) *DeleteOutboundCallRestrictionResponse {
	s.Headers = v
	return s
}

func (s *DeleteOutboundCallRestrictionResponse) SetStatusCode(v int32) *DeleteOutboundCallRestrictionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOutboundCallRestrictionResponse) SetBody(v *DeleteOutboundCallRestrictionResponseBody) *DeleteOutboundCallRestrictionResponse {
	s.Body = v
	return s
}

func (s *DeleteOutboundCallRestrictionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
