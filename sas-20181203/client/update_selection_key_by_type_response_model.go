// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSelectionKeyByTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSelectionKeyByTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSelectionKeyByTypeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSelectionKeyByTypeResponseBody) *UpdateSelectionKeyByTypeResponse
	GetBody() *UpdateSelectionKeyByTypeResponseBody
}

type UpdateSelectionKeyByTypeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSelectionKeyByTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSelectionKeyByTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSelectionKeyByTypeResponse) GoString() string {
	return s.String()
}

func (s *UpdateSelectionKeyByTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSelectionKeyByTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSelectionKeyByTypeResponse) GetBody() *UpdateSelectionKeyByTypeResponseBody {
	return s.Body
}

func (s *UpdateSelectionKeyByTypeResponse) SetHeaders(v map[string]*string) *UpdateSelectionKeyByTypeResponse {
	s.Headers = v
	return s
}

func (s *UpdateSelectionKeyByTypeResponse) SetStatusCode(v int32) *UpdateSelectionKeyByTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSelectionKeyByTypeResponse) SetBody(v *UpdateSelectionKeyByTypeResponseBody) *UpdateSelectionKeyByTypeResponse {
	s.Body = v
	return s
}

func (s *UpdateSelectionKeyByTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
