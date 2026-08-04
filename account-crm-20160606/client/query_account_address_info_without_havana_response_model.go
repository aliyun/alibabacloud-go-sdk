// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountAddressInfoWithoutHavanaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAccountAddressInfoWithoutHavanaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAccountAddressInfoWithoutHavanaResponse
	GetStatusCode() *int32
	SetBody(v *QueryAccountAddressInfoWithoutHavanaResponseBody) *QueryAccountAddressInfoWithoutHavanaResponse
	GetBody() *QueryAccountAddressInfoWithoutHavanaResponseBody
}

type QueryAccountAddressInfoWithoutHavanaResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAccountAddressInfoWithoutHavanaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAccountAddressInfoWithoutHavanaResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountAddressInfoWithoutHavanaResponse) GoString() string {
	return s.String()
}

func (s *QueryAccountAddressInfoWithoutHavanaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAccountAddressInfoWithoutHavanaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAccountAddressInfoWithoutHavanaResponse) GetBody() *QueryAccountAddressInfoWithoutHavanaResponseBody {
	return s.Body
}

func (s *QueryAccountAddressInfoWithoutHavanaResponse) SetHeaders(v map[string]*string) *QueryAccountAddressInfoWithoutHavanaResponse {
	s.Headers = v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponse) SetStatusCode(v int32) *QueryAccountAddressInfoWithoutHavanaResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponse) SetBody(v *QueryAccountAddressInfoWithoutHavanaResponseBody) *QueryAccountAddressInfoWithoutHavanaResponse {
	s.Body = v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
