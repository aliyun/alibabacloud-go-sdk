// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCollectionDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCollectionDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCollectionDataResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCollectionDataResponseBody) *DeleteCollectionDataResponse
	GetBody() *DeleteCollectionDataResponseBody
}

type DeleteCollectionDataResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCollectionDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCollectionDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCollectionDataResponse) GoString() string {
	return s.String()
}

func (s *DeleteCollectionDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCollectionDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCollectionDataResponse) GetBody() *DeleteCollectionDataResponseBody {
	return s.Body
}

func (s *DeleteCollectionDataResponse) SetHeaders(v map[string]*string) *DeleteCollectionDataResponse {
	s.Headers = v
	return s
}

func (s *DeleteCollectionDataResponse) SetStatusCode(v int32) *DeleteCollectionDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCollectionDataResponse) SetBody(v *DeleteCollectionDataResponseBody) *DeleteCollectionDataResponse {
	s.Body = v
	return s
}

func (s *DeleteCollectionDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
