// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMapFromHavanaBindIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MapFromHavanaBindIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MapFromHavanaBindIdResponse
	GetStatusCode() *int32
	SetBody(v *MapFromHavanaBindIdResponseBody) *MapFromHavanaBindIdResponse
	GetBody() *MapFromHavanaBindIdResponseBody
}

type MapFromHavanaBindIdResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MapFromHavanaBindIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MapFromHavanaBindIdResponse) String() string {
	return dara.Prettify(s)
}

func (s MapFromHavanaBindIdResponse) GoString() string {
	return s.String()
}

func (s *MapFromHavanaBindIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MapFromHavanaBindIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MapFromHavanaBindIdResponse) GetBody() *MapFromHavanaBindIdResponseBody {
	return s.Body
}

func (s *MapFromHavanaBindIdResponse) SetHeaders(v map[string]*string) *MapFromHavanaBindIdResponse {
	s.Headers = v
	return s
}

func (s *MapFromHavanaBindIdResponse) SetStatusCode(v int32) *MapFromHavanaBindIdResponse {
	s.StatusCode = &v
	return s
}

func (s *MapFromHavanaBindIdResponse) SetBody(v *MapFromHavanaBindIdResponseBody) *MapFromHavanaBindIdResponse {
	s.Body = v
	return s
}

func (s *MapFromHavanaBindIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
