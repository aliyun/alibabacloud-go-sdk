// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityScanRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataQualityScanRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataQualityScanRunResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataQualityScanRunResponseBody) *CreateDataQualityScanRunResponse
	GetBody() *CreateDataQualityScanRunResponseBody
}

type CreateDataQualityScanRunResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataQualityScanRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataQualityScanRunResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityScanRunResponse) GoString() string {
	return s.String()
}

func (s *CreateDataQualityScanRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataQualityScanRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataQualityScanRunResponse) GetBody() *CreateDataQualityScanRunResponseBody {
	return s.Body
}

func (s *CreateDataQualityScanRunResponse) SetHeaders(v map[string]*string) *CreateDataQualityScanRunResponse {
	s.Headers = v
	return s
}

func (s *CreateDataQualityScanRunResponse) SetStatusCode(v int32) *CreateDataQualityScanRunResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataQualityScanRunResponse) SetBody(v *CreateDataQualityScanRunResponseBody) *CreateDataQualityScanRunResponse {
	s.Body = v
	return s
}

func (s *CreateDataQualityScanRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
