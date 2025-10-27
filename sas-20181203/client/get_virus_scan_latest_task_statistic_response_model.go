// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVirusScanLatestTaskStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVirusScanLatestTaskStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVirusScanLatestTaskStatisticResponse
	GetStatusCode() *int32
	SetBody(v *GetVirusScanLatestTaskStatisticResponseBody) *GetVirusScanLatestTaskStatisticResponse
	GetBody() *GetVirusScanLatestTaskStatisticResponseBody
}

type GetVirusScanLatestTaskStatisticResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVirusScanLatestTaskStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVirusScanLatestTaskStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVirusScanLatestTaskStatisticResponse) GoString() string {
	return s.String()
}

func (s *GetVirusScanLatestTaskStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVirusScanLatestTaskStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVirusScanLatestTaskStatisticResponse) GetBody() *GetVirusScanLatestTaskStatisticResponseBody {
	return s.Body
}

func (s *GetVirusScanLatestTaskStatisticResponse) SetHeaders(v map[string]*string) *GetVirusScanLatestTaskStatisticResponse {
	s.Headers = v
	return s
}

func (s *GetVirusScanLatestTaskStatisticResponse) SetStatusCode(v int32) *GetVirusScanLatestTaskStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVirusScanLatestTaskStatisticResponse) SetBody(v *GetVirusScanLatestTaskStatisticResponseBody) *GetVirusScanLatestTaskStatisticResponse {
	s.Body = v
	return s
}

func (s *GetVirusScanLatestTaskStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
