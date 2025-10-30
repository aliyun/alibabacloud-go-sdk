// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelUpsertCollectionDataJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelUpsertCollectionDataJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelUpsertCollectionDataJobResponse
	GetStatusCode() *int32
	SetBody(v *CancelUpsertCollectionDataJobResponseBody) *CancelUpsertCollectionDataJobResponse
	GetBody() *CancelUpsertCollectionDataJobResponseBody
}

type CancelUpsertCollectionDataJobResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelUpsertCollectionDataJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelUpsertCollectionDataJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelUpsertCollectionDataJobResponse) GoString() string {
	return s.String()
}

func (s *CancelUpsertCollectionDataJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelUpsertCollectionDataJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelUpsertCollectionDataJobResponse) GetBody() *CancelUpsertCollectionDataJobResponseBody {
	return s.Body
}

func (s *CancelUpsertCollectionDataJobResponse) SetHeaders(v map[string]*string) *CancelUpsertCollectionDataJobResponse {
	s.Headers = v
	return s
}

func (s *CancelUpsertCollectionDataJobResponse) SetStatusCode(v int32) *CancelUpsertCollectionDataJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelUpsertCollectionDataJobResponse) SetBody(v *CancelUpsertCollectionDataJobResponseBody) *CancelUpsertCollectionDataJobResponse {
	s.Body = v
	return s
}

func (s *CancelUpsertCollectionDataJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
