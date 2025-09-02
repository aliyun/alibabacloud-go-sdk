// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgPlatformQueryProjectsAndSchemaFromMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgPlatformQueryProjectsAndSchemaFromMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgPlatformQueryProjectsAndSchemaFromMetaResponse
	GetStatusCode() *int32
	SetBody(v *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody) *DsgPlatformQueryProjectsAndSchemaFromMetaResponse
	GetBody() *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody
}

type DsgPlatformQueryProjectsAndSchemaFromMetaResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgPlatformQueryProjectsAndSchemaFromMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgPlatformQueryProjectsAndSchemaFromMetaResponse) GoString() string {
	return s.String()
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponse) GetBody() *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody {
	return s.Body
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponse) SetHeaders(v map[string]*string) *DsgPlatformQueryProjectsAndSchemaFromMetaResponse {
	s.Headers = v
	return s
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponse) SetStatusCode(v int32) *DsgPlatformQueryProjectsAndSchemaFromMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponse) SetBody(v *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody) *DsgPlatformQueryProjectsAndSchemaFromMetaResponse {
	s.Body = v
	return s
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponse) Validate() error {
	return dara.Validate(s)
}
