// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertCollectionDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpsertCollectionDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpsertCollectionDataResponse
	GetStatusCode() *int32
	SetBody(v *UpsertCollectionDataResponseBody) *UpsertCollectionDataResponse
	GetBody() *UpsertCollectionDataResponseBody
}

type UpsertCollectionDataResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpsertCollectionDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpsertCollectionDataResponse) String() string {
	return dara.Prettify(s)
}

func (s UpsertCollectionDataResponse) GoString() string {
	return s.String()
}

func (s *UpsertCollectionDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpsertCollectionDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpsertCollectionDataResponse) GetBody() *UpsertCollectionDataResponseBody {
	return s.Body
}

func (s *UpsertCollectionDataResponse) SetHeaders(v map[string]*string) *UpsertCollectionDataResponse {
	s.Headers = v
	return s
}

func (s *UpsertCollectionDataResponse) SetStatusCode(v int32) *UpsertCollectionDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UpsertCollectionDataResponse) SetBody(v *UpsertCollectionDataResponseBody) *UpsertCollectionDataResponse {
	s.Body = v
	return s
}

func (s *UpsertCollectionDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
