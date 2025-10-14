// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataBatchIngestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataBatchIngestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataBatchIngestionResponse
	GetStatusCode() *int32
	SetBody(v *GetDataBatchIngestionResponseBody) *GetDataBatchIngestionResponse
	GetBody() *GetDataBatchIngestionResponseBody
}

type GetDataBatchIngestionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataBatchIngestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataBatchIngestionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataBatchIngestionResponse) GoString() string {
	return s.String()
}

func (s *GetDataBatchIngestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataBatchIngestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataBatchIngestionResponse) GetBody() *GetDataBatchIngestionResponseBody {
	return s.Body
}

func (s *GetDataBatchIngestionResponse) SetHeaders(v map[string]*string) *GetDataBatchIngestionResponse {
	s.Headers = v
	return s
}

func (s *GetDataBatchIngestionResponse) SetStatusCode(v int32) *GetDataBatchIngestionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataBatchIngestionResponse) SetBody(v *GetDataBatchIngestionResponseBody) *GetDataBatchIngestionResponse {
	s.Body = v
	return s
}

func (s *GetDataBatchIngestionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
