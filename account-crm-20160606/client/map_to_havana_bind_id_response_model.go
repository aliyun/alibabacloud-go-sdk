// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMapToHavanaBindIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MapToHavanaBindIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MapToHavanaBindIdResponse
	GetStatusCode() *int32
	SetBody(v *MapToHavanaBindIdResponseBody) *MapToHavanaBindIdResponse
	GetBody() *MapToHavanaBindIdResponseBody
}

type MapToHavanaBindIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MapToHavanaBindIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MapToHavanaBindIdResponse) String() string {
	return dara.Prettify(s)
}

func (s MapToHavanaBindIdResponse) GoString() string {
	return s.String()
}

func (s *MapToHavanaBindIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MapToHavanaBindIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MapToHavanaBindIdResponse) GetBody() *MapToHavanaBindIdResponseBody {
	return s.Body
}

func (s *MapToHavanaBindIdResponse) SetHeaders(v map[string]*string) *MapToHavanaBindIdResponse {
	s.Headers = v
	return s
}

func (s *MapToHavanaBindIdResponse) SetStatusCode(v int32) *MapToHavanaBindIdResponse {
	s.StatusCode = &v
	return s
}

func (s *MapToHavanaBindIdResponse) SetBody(v *MapToHavanaBindIdResponseBody) *MapToHavanaBindIdResponse {
	s.Body = v
	return s
}

func (s *MapToHavanaBindIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
