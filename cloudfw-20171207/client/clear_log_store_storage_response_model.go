// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearLogStoreStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClearLogStoreStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClearLogStoreStorageResponse
	GetStatusCode() *int32
	SetBody(v *ClearLogStoreStorageResponseBody) *ClearLogStoreStorageResponse
	GetBody() *ClearLogStoreStorageResponseBody
}

type ClearLogStoreStorageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearLogStoreStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearLogStoreStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s ClearLogStoreStorageResponse) GoString() string {
	return s.String()
}

func (s *ClearLogStoreStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClearLogStoreStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClearLogStoreStorageResponse) GetBody() *ClearLogStoreStorageResponseBody {
	return s.Body
}

func (s *ClearLogStoreStorageResponse) SetHeaders(v map[string]*string) *ClearLogStoreStorageResponse {
	s.Headers = v
	return s
}

func (s *ClearLogStoreStorageResponse) SetStatusCode(v int32) *ClearLogStoreStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearLogStoreStorageResponse) SetBody(v *ClearLogStoreStorageResponseBody) *ClearLogStoreStorageResponse {
	s.Body = v
	return s
}

func (s *ClearLogStoreStorageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
