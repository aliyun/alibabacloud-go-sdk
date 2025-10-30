// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCollectionDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCollectionDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCollectionDataResponse
	GetStatusCode() *int32
	SetBody(v *QueryCollectionDataResponseBody) *QueryCollectionDataResponse
	GetBody() *QueryCollectionDataResponseBody
}

type QueryCollectionDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCollectionDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCollectionDataResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCollectionDataResponse) GoString() string {
	return s.String()
}

func (s *QueryCollectionDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCollectionDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCollectionDataResponse) GetBody() *QueryCollectionDataResponseBody {
	return s.Body
}

func (s *QueryCollectionDataResponse) SetHeaders(v map[string]*string) *QueryCollectionDataResponse {
	s.Headers = v
	return s
}

func (s *QueryCollectionDataResponse) SetStatusCode(v int32) *QueryCollectionDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCollectionDataResponse) SetBody(v *QueryCollectionDataResponseBody) *QueryCollectionDataResponse {
	s.Body = v
	return s
}

func (s *QueryCollectionDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
