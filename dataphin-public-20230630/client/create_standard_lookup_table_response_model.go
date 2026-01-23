// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardLookupTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStandardLookupTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStandardLookupTableResponse
	GetStatusCode() *int32
	SetBody(v *CreateStandardLookupTableResponseBody) *CreateStandardLookupTableResponse
	GetBody() *CreateStandardLookupTableResponseBody
}

type CreateStandardLookupTableResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStandardLookupTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStandardLookupTableResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardLookupTableResponse) GoString() string {
	return s.String()
}

func (s *CreateStandardLookupTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStandardLookupTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStandardLookupTableResponse) GetBody() *CreateStandardLookupTableResponseBody {
	return s.Body
}

func (s *CreateStandardLookupTableResponse) SetHeaders(v map[string]*string) *CreateStandardLookupTableResponse {
	s.Headers = v
	return s
}

func (s *CreateStandardLookupTableResponse) SetStatusCode(v int32) *CreateStandardLookupTableResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStandardLookupTableResponse) SetBody(v *CreateStandardLookupTableResponseBody) *CreateStandardLookupTableResponse {
	s.Body = v
	return s
}

func (s *CreateStandardLookupTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
