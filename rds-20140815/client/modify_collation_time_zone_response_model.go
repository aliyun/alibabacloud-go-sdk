// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCollationTimeZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCollationTimeZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCollationTimeZoneResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCollationTimeZoneResponseBody) *ModifyCollationTimeZoneResponse
	GetBody() *ModifyCollationTimeZoneResponseBody
}

type ModifyCollationTimeZoneResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCollationTimeZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCollationTimeZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCollationTimeZoneResponse) GoString() string {
	return s.String()
}

func (s *ModifyCollationTimeZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCollationTimeZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCollationTimeZoneResponse) GetBody() *ModifyCollationTimeZoneResponseBody {
	return s.Body
}

func (s *ModifyCollationTimeZoneResponse) SetHeaders(v map[string]*string) *ModifyCollationTimeZoneResponse {
	s.Headers = v
	return s
}

func (s *ModifyCollationTimeZoneResponse) SetStatusCode(v int32) *ModifyCollationTimeZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCollationTimeZoneResponse) SetBody(v *ModifyCollationTimeZoneResponseBody) *ModifyCollationTimeZoneResponse {
	s.Body = v
	return s
}

func (s *ModifyCollationTimeZoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
