// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertCollectionDataAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpsertCollectionDataAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpsertCollectionDataAsyncResponse
	GetStatusCode() *int32
	SetBody(v *UpsertCollectionDataAsyncResponseBody) *UpsertCollectionDataAsyncResponse
	GetBody() *UpsertCollectionDataAsyncResponseBody
}

type UpsertCollectionDataAsyncResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpsertCollectionDataAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpsertCollectionDataAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s UpsertCollectionDataAsyncResponse) GoString() string {
	return s.String()
}

func (s *UpsertCollectionDataAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpsertCollectionDataAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpsertCollectionDataAsyncResponse) GetBody() *UpsertCollectionDataAsyncResponseBody {
	return s.Body
}

func (s *UpsertCollectionDataAsyncResponse) SetHeaders(v map[string]*string) *UpsertCollectionDataAsyncResponse {
	s.Headers = v
	return s
}

func (s *UpsertCollectionDataAsyncResponse) SetStatusCode(v int32) *UpsertCollectionDataAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *UpsertCollectionDataAsyncResponse) SetBody(v *UpsertCollectionDataAsyncResponseBody) *UpsertCollectionDataAsyncResponse {
	s.Body = v
	return s
}

func (s *UpsertCollectionDataAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
