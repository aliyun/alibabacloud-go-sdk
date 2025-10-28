// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertOrUpdateRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertOrUpdateRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertOrUpdateRegionResponse
	GetStatusCode() *int32
	SetBody(v *InsertOrUpdateRegionResponseBody) *InsertOrUpdateRegionResponse
	GetBody() *InsertOrUpdateRegionResponseBody
}

type InsertOrUpdateRegionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertOrUpdateRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertOrUpdateRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertOrUpdateRegionResponse) GoString() string {
	return s.String()
}

func (s *InsertOrUpdateRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertOrUpdateRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertOrUpdateRegionResponse) GetBody() *InsertOrUpdateRegionResponseBody {
	return s.Body
}

func (s *InsertOrUpdateRegionResponse) SetHeaders(v map[string]*string) *InsertOrUpdateRegionResponse {
	s.Headers = v
	return s
}

func (s *InsertOrUpdateRegionResponse) SetStatusCode(v int32) *InsertOrUpdateRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertOrUpdateRegionResponse) SetBody(v *InsertOrUpdateRegionResponseBody) *InsertOrUpdateRegionResponse {
	s.Body = v
	return s
}

func (s *InsertOrUpdateRegionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
