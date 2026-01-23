// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardLookupTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStandardLookupTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStandardLookupTableResponse
	GetStatusCode() *int32
	SetBody(v *GetStandardLookupTableResponseBody) *GetStandardLookupTableResponse
	GetBody() *GetStandardLookupTableResponseBody
}

type GetStandardLookupTableResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStandardLookupTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStandardLookupTableResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStandardLookupTableResponse) GoString() string {
	return s.String()
}

func (s *GetStandardLookupTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStandardLookupTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStandardLookupTableResponse) GetBody() *GetStandardLookupTableResponseBody {
	return s.Body
}

func (s *GetStandardLookupTableResponse) SetHeaders(v map[string]*string) *GetStandardLookupTableResponse {
	s.Headers = v
	return s
}

func (s *GetStandardLookupTableResponse) SetStatusCode(v int32) *GetStandardLookupTableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStandardLookupTableResponse) SetBody(v *GetStandardLookupTableResponseBody) *GetStandardLookupTableResponse {
	s.Body = v
	return s
}

func (s *GetStandardLookupTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
