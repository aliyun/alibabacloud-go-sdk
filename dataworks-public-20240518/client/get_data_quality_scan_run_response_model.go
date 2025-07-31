// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityScanRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataQualityScanRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataQualityScanRunResponse
	GetStatusCode() *int32
	SetBody(v *GetDataQualityScanRunResponseBody) *GetDataQualityScanRunResponse
	GetBody() *GetDataQualityScanRunResponseBody
}

type GetDataQualityScanRunResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataQualityScanRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataQualityScanRunResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanRunResponse) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataQualityScanRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataQualityScanRunResponse) GetBody() *GetDataQualityScanRunResponseBody {
	return s.Body
}

func (s *GetDataQualityScanRunResponse) SetHeaders(v map[string]*string) *GetDataQualityScanRunResponse {
	s.Headers = v
	return s
}

func (s *GetDataQualityScanRunResponse) SetStatusCode(v int32) *GetDataQualityScanRunResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataQualityScanRunResponse) SetBody(v *GetDataQualityScanRunResponseBody) *GetDataQualityScanRunResponse {
	s.Body = v
	return s
}

func (s *GetDataQualityScanRunResponse) Validate() error {
	return dara.Validate(s)
}
