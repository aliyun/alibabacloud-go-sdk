// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardLookupTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStandardLookupTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStandardLookupTableResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStandardLookupTableResponseBody) *DeleteStandardLookupTableResponse
	GetBody() *DeleteStandardLookupTableResponseBody
}

type DeleteStandardLookupTableResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStandardLookupTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStandardLookupTableResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardLookupTableResponse) GoString() string {
	return s.String()
}

func (s *DeleteStandardLookupTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStandardLookupTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStandardLookupTableResponse) GetBody() *DeleteStandardLookupTableResponseBody {
	return s.Body
}

func (s *DeleteStandardLookupTableResponse) SetHeaders(v map[string]*string) *DeleteStandardLookupTableResponse {
	s.Headers = v
	return s
}

func (s *DeleteStandardLookupTableResponse) SetStatusCode(v int32) *DeleteStandardLookupTableResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStandardLookupTableResponse) SetBody(v *DeleteStandardLookupTableResponseBody) *DeleteStandardLookupTableResponse {
	s.Body = v
	return s
}

func (s *DeleteStandardLookupTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
