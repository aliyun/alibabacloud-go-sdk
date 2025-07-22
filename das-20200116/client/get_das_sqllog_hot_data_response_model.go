// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDasSQLLogHotDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDasSQLLogHotDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDasSQLLogHotDataResponse
	GetStatusCode() *int32
	SetBody(v *GetDasSQLLogHotDataResponseBody) *GetDasSQLLogHotDataResponse
	GetBody() *GetDasSQLLogHotDataResponseBody
}

type GetDasSQLLogHotDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDasSQLLogHotDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDasSQLLogHotDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDasSQLLogHotDataResponse) GoString() string {
	return s.String()
}

func (s *GetDasSQLLogHotDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDasSQLLogHotDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDasSQLLogHotDataResponse) GetBody() *GetDasSQLLogHotDataResponseBody {
	return s.Body
}

func (s *GetDasSQLLogHotDataResponse) SetHeaders(v map[string]*string) *GetDasSQLLogHotDataResponse {
	s.Headers = v
	return s
}

func (s *GetDasSQLLogHotDataResponse) SetStatusCode(v int32) *GetDasSQLLogHotDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDasSQLLogHotDataResponse) SetBody(v *GetDasSQLLogHotDataResponseBody) *GetDasSQLLogHotDataResponse {
	s.Body = v
	return s
}

func (s *GetDasSQLLogHotDataResponse) Validate() error {
	return dara.Validate(s)
}
