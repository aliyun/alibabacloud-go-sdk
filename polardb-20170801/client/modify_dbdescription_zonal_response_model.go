// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBDescriptionZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBDescriptionZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBDescriptionZonalResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBDescriptionZonalResponseBody) *ModifyDBDescriptionZonalResponse
	GetBody() *ModifyDBDescriptionZonalResponseBody
}

type ModifyDBDescriptionZonalResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBDescriptionZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBDescriptionZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBDescriptionZonalResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBDescriptionZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBDescriptionZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBDescriptionZonalResponse) GetBody() *ModifyDBDescriptionZonalResponseBody {
	return s.Body
}

func (s *ModifyDBDescriptionZonalResponse) SetHeaders(v map[string]*string) *ModifyDBDescriptionZonalResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBDescriptionZonalResponse) SetStatusCode(v int32) *ModifyDBDescriptionZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBDescriptionZonalResponse) SetBody(v *ModifyDBDescriptionZonalResponseBody) *ModifyDBDescriptionZonalResponse {
	s.Body = v
	return s
}

func (s *ModifyDBDescriptionZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
