// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDictDataItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDictDataItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDictDataItemResponse
	GetStatusCode() *int32
	SetBody(v *QueryDictDataItemResponseBody) *QueryDictDataItemResponse
	GetBody() *QueryDictDataItemResponseBody
}

type QueryDictDataItemResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDictDataItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDictDataItemResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDictDataItemResponse) GoString() string {
	return s.String()
}

func (s *QueryDictDataItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDictDataItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDictDataItemResponse) GetBody() *QueryDictDataItemResponseBody {
	return s.Body
}

func (s *QueryDictDataItemResponse) SetHeaders(v map[string]*string) *QueryDictDataItemResponse {
	s.Headers = v
	return s
}

func (s *QueryDictDataItemResponse) SetStatusCode(v int32) *QueryDictDataItemResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDictDataItemResponse) SetBody(v *QueryDictDataItemResponseBody) *QueryDictDataItemResponse {
	s.Body = v
	return s
}

func (s *QueryDictDataItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
