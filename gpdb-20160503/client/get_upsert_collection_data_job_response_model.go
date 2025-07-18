// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUpsertCollectionDataJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUpsertCollectionDataJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUpsertCollectionDataJobResponse
	GetStatusCode() *int32
	SetBody(v *GetUpsertCollectionDataJobResponseBody) *GetUpsertCollectionDataJobResponse
	GetBody() *GetUpsertCollectionDataJobResponseBody
}

type GetUpsertCollectionDataJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUpsertCollectionDataJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUpsertCollectionDataJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUpsertCollectionDataJobResponse) GoString() string {
	return s.String()
}

func (s *GetUpsertCollectionDataJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUpsertCollectionDataJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUpsertCollectionDataJobResponse) GetBody() *GetUpsertCollectionDataJobResponseBody {
	return s.Body
}

func (s *GetUpsertCollectionDataJobResponse) SetHeaders(v map[string]*string) *GetUpsertCollectionDataJobResponse {
	s.Headers = v
	return s
}

func (s *GetUpsertCollectionDataJobResponse) SetStatusCode(v int32) *GetUpsertCollectionDataJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUpsertCollectionDataJobResponse) SetBody(v *GetUpsertCollectionDataJobResponseBody) *GetUpsertCollectionDataJobResponse {
	s.Body = v
	return s
}

func (s *GetUpsertCollectionDataJobResponse) Validate() error {
	return dara.Validate(s)
}
