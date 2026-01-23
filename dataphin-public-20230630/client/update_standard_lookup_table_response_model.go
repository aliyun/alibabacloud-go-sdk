// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardLookupTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateStandardLookupTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateStandardLookupTableResponse
	GetStatusCode() *int32
	SetBody(v *UpdateStandardLookupTableResponseBody) *UpdateStandardLookupTableResponse
	GetBody() *UpdateStandardLookupTableResponseBody
}

type UpdateStandardLookupTableResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStandardLookupTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStandardLookupTableResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardLookupTableResponse) GoString() string {
	return s.String()
}

func (s *UpdateStandardLookupTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateStandardLookupTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateStandardLookupTableResponse) GetBody() *UpdateStandardLookupTableResponseBody {
	return s.Body
}

func (s *UpdateStandardLookupTableResponse) SetHeaders(v map[string]*string) *UpdateStandardLookupTableResponse {
	s.Headers = v
	return s
}

func (s *UpdateStandardLookupTableResponse) SetStatusCode(v int32) *UpdateStandardLookupTableResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStandardLookupTableResponse) SetBody(v *UpdateStandardLookupTableResponseBody) *UpdateStandardLookupTableResponse {
	s.Body = v
	return s
}

func (s *UpdateStandardLookupTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
