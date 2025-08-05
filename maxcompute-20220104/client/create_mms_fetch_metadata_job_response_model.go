// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMmsFetchMetadataJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMmsFetchMetadataJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMmsFetchMetadataJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateMmsFetchMetadataJobResponseBody) *CreateMmsFetchMetadataJobResponse
	GetBody() *CreateMmsFetchMetadataJobResponseBody
}

type CreateMmsFetchMetadataJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMmsFetchMetadataJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMmsFetchMetadataJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMmsFetchMetadataJobResponse) GoString() string {
	return s.String()
}

func (s *CreateMmsFetchMetadataJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMmsFetchMetadataJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMmsFetchMetadataJobResponse) GetBody() *CreateMmsFetchMetadataJobResponseBody {
	return s.Body
}

func (s *CreateMmsFetchMetadataJobResponse) SetHeaders(v map[string]*string) *CreateMmsFetchMetadataJobResponse {
	s.Headers = v
	return s
}

func (s *CreateMmsFetchMetadataJobResponse) SetStatusCode(v int32) *CreateMmsFetchMetadataJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMmsFetchMetadataJobResponse) SetBody(v *CreateMmsFetchMetadataJobResponseBody) *CreateMmsFetchMetadataJobResponse {
	s.Body = v
	return s
}

func (s *CreateMmsFetchMetadataJobResponse) Validate() error {
	return dara.Validate(s)
}
