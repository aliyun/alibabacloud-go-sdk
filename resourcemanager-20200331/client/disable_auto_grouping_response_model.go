// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAutoGroupingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableAutoGroupingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableAutoGroupingResponse
	GetStatusCode() *int32
	SetBody(v *DisableAutoGroupingResponseBody) *DisableAutoGroupingResponse
	GetBody() *DisableAutoGroupingResponseBody
}

type DisableAutoGroupingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableAutoGroupingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableAutoGroupingResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableAutoGroupingResponse) GoString() string {
	return s.String()
}

func (s *DisableAutoGroupingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableAutoGroupingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableAutoGroupingResponse) GetBody() *DisableAutoGroupingResponseBody {
	return s.Body
}

func (s *DisableAutoGroupingResponse) SetHeaders(v map[string]*string) *DisableAutoGroupingResponse {
	s.Headers = v
	return s
}

func (s *DisableAutoGroupingResponse) SetStatusCode(v int32) *DisableAutoGroupingResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAutoGroupingResponse) SetBody(v *DisableAutoGroupingResponseBody) *DisableAutoGroupingResponse {
	s.Body = v
	return s
}

func (s *DisableAutoGroupingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
