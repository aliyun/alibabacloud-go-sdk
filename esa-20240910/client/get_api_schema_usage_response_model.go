// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiSchemaUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApiSchemaUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApiSchemaUsageResponse
	GetStatusCode() *int32
	SetBody(v *GetApiSchemaUsageResponseBody) *GetApiSchemaUsageResponse
	GetBody() *GetApiSchemaUsageResponseBody
}

type GetApiSchemaUsageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApiSchemaUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApiSchemaUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApiSchemaUsageResponse) GoString() string {
	return s.String()
}

func (s *GetApiSchemaUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApiSchemaUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApiSchemaUsageResponse) GetBody() *GetApiSchemaUsageResponseBody {
	return s.Body
}

func (s *GetApiSchemaUsageResponse) SetHeaders(v map[string]*string) *GetApiSchemaUsageResponse {
	s.Headers = v
	return s
}

func (s *GetApiSchemaUsageResponse) SetStatusCode(v int32) *GetApiSchemaUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApiSchemaUsageResponse) SetBody(v *GetApiSchemaUsageResponseBody) *GetApiSchemaUsageResponse {
	s.Body = v
	return s
}

func (s *GetApiSchemaUsageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
