// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetFileMetasStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDatasetFileMetasStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDatasetFileMetasStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetDatasetFileMetasStatisticsResponseBody) *GetDatasetFileMetasStatisticsResponse
	GetBody() *GetDatasetFileMetasStatisticsResponseBody
}

type GetDatasetFileMetasStatisticsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatasetFileMetasStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatasetFileMetasStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetFileMetasStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetDatasetFileMetasStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDatasetFileMetasStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDatasetFileMetasStatisticsResponse) GetBody() *GetDatasetFileMetasStatisticsResponseBody {
	return s.Body
}

func (s *GetDatasetFileMetasStatisticsResponse) SetHeaders(v map[string]*string) *GetDatasetFileMetasStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetDatasetFileMetasStatisticsResponse) SetStatusCode(v int32) *GetDatasetFileMetasStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatasetFileMetasStatisticsResponse) SetBody(v *GetDatasetFileMetasStatisticsResponseBody) *GetDatasetFileMetasStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetDatasetFileMetasStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
