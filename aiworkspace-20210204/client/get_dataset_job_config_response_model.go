// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetJobConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDatasetJobConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDatasetJobConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetDatasetJobConfigResponseBody) *GetDatasetJobConfigResponse
	GetBody() *GetDatasetJobConfigResponseBody
}

type GetDatasetJobConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatasetJobConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatasetJobConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetJobConfigResponse) GoString() string {
	return s.String()
}

func (s *GetDatasetJobConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDatasetJobConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDatasetJobConfigResponse) GetBody() *GetDatasetJobConfigResponseBody {
	return s.Body
}

func (s *GetDatasetJobConfigResponse) SetHeaders(v map[string]*string) *GetDatasetJobConfigResponse {
	s.Headers = v
	return s
}

func (s *GetDatasetJobConfigResponse) SetStatusCode(v int32) *GetDatasetJobConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatasetJobConfigResponse) SetBody(v *GetDatasetJobConfigResponseBody) *GetDatasetJobConfigResponse {
	s.Body = v
	return s
}

func (s *GetDatasetJobConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
