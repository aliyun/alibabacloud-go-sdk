// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMmsFetchMetadataJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMmsFetchMetadataJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMmsFetchMetadataJobResponse
	GetStatusCode() *int32
	SetBody(v *GetMmsFetchMetadataJobResponseBody) *GetMmsFetchMetadataJobResponse
	GetBody() *GetMmsFetchMetadataJobResponseBody
}

type GetMmsFetchMetadataJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMmsFetchMetadataJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMmsFetchMetadataJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMmsFetchMetadataJobResponse) GoString() string {
	return s.String()
}

func (s *GetMmsFetchMetadataJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMmsFetchMetadataJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMmsFetchMetadataJobResponse) GetBody() *GetMmsFetchMetadataJobResponseBody {
	return s.Body
}

func (s *GetMmsFetchMetadataJobResponse) SetHeaders(v map[string]*string) *GetMmsFetchMetadataJobResponse {
	s.Headers = v
	return s
}

func (s *GetMmsFetchMetadataJobResponse) SetStatusCode(v int32) *GetMmsFetchMetadataJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMmsFetchMetadataJobResponse) SetBody(v *GetMmsFetchMetadataJobResponseBody) *GetMmsFetchMetadataJobResponse {
	s.Body = v
	return s
}

func (s *GetMmsFetchMetadataJobResponse) Validate() error {
	return dara.Validate(s)
}
