// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityScanRunLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataQualityScanRunLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataQualityScanRunLogResponse
	GetStatusCode() *int32
	SetBody(v *GetDataQualityScanRunLogResponseBody) *GetDataQualityScanRunLogResponse
	GetBody() *GetDataQualityScanRunLogResponseBody
}

type GetDataQualityScanRunLogResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataQualityScanRunLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataQualityScanRunLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanRunLogResponse) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanRunLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataQualityScanRunLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataQualityScanRunLogResponse) GetBody() *GetDataQualityScanRunLogResponseBody {
	return s.Body
}

func (s *GetDataQualityScanRunLogResponse) SetHeaders(v map[string]*string) *GetDataQualityScanRunLogResponse {
	s.Headers = v
	return s
}

func (s *GetDataQualityScanRunLogResponse) SetStatusCode(v int32) *GetDataQualityScanRunLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataQualityScanRunLogResponse) SetBody(v *GetDataQualityScanRunLogResponseBody) *GetDataQualityScanRunLogResponse {
	s.Body = v
	return s
}

func (s *GetDataQualityScanRunLogResponse) Validate() error {
	return dara.Validate(s)
}
