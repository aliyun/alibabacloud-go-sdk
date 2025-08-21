// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncCatalogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncCatalogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncCatalogsResponse
	GetStatusCode() *int32
	SetBody(v *SyncCatalogsResponseBody) *SyncCatalogsResponse
	GetBody() *SyncCatalogsResponseBody
}

type SyncCatalogsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncCatalogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncCatalogsResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncCatalogsResponse) GoString() string {
	return s.String()
}

func (s *SyncCatalogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncCatalogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncCatalogsResponse) GetBody() *SyncCatalogsResponseBody {
	return s.Body
}

func (s *SyncCatalogsResponse) SetHeaders(v map[string]*string) *SyncCatalogsResponse {
	s.Headers = v
	return s
}

func (s *SyncCatalogsResponse) SetStatusCode(v int32) *SyncCatalogsResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncCatalogsResponse) SetBody(v *SyncCatalogsResponseBody) *SyncCatalogsResponse {
	s.Body = v
	return s
}

func (s *SyncCatalogsResponse) Validate() error {
	return dara.Validate(s)
}
