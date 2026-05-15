// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertAiOutboundPhoneNumsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertAiOutboundPhoneNumsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertAiOutboundPhoneNumsResponse
	GetStatusCode() *int32
	SetBody(v *InsertAiOutboundPhoneNumsResponseBody) *InsertAiOutboundPhoneNumsResponse
	GetBody() *InsertAiOutboundPhoneNumsResponseBody
}

type InsertAiOutboundPhoneNumsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertAiOutboundPhoneNumsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertAiOutboundPhoneNumsResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertAiOutboundPhoneNumsResponse) GoString() string {
	return s.String()
}

func (s *InsertAiOutboundPhoneNumsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertAiOutboundPhoneNumsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertAiOutboundPhoneNumsResponse) GetBody() *InsertAiOutboundPhoneNumsResponseBody {
	return s.Body
}

func (s *InsertAiOutboundPhoneNumsResponse) SetHeaders(v map[string]*string) *InsertAiOutboundPhoneNumsResponse {
	s.Headers = v
	return s
}

func (s *InsertAiOutboundPhoneNumsResponse) SetStatusCode(v int32) *InsertAiOutboundPhoneNumsResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertAiOutboundPhoneNumsResponse) SetBody(v *InsertAiOutboundPhoneNumsResponseBody) *InsertAiOutboundPhoneNumsResponse {
	s.Body = v
	return s
}

func (s *InsertAiOutboundPhoneNumsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
