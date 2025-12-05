// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJMeterSceneRunningDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJMeterSceneRunningDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJMeterSceneRunningDataResponse
	GetStatusCode() *int32
	SetBody(v *GetJMeterSceneRunningDataResponseBody) *GetJMeterSceneRunningDataResponse
	GetBody() *GetJMeterSceneRunningDataResponseBody
}

type GetJMeterSceneRunningDataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJMeterSceneRunningDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJMeterSceneRunningDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJMeterSceneRunningDataResponse) GoString() string {
	return s.String()
}

func (s *GetJMeterSceneRunningDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJMeterSceneRunningDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJMeterSceneRunningDataResponse) GetBody() *GetJMeterSceneRunningDataResponseBody {
	return s.Body
}

func (s *GetJMeterSceneRunningDataResponse) SetHeaders(v map[string]*string) *GetJMeterSceneRunningDataResponse {
	s.Headers = v
	return s
}

func (s *GetJMeterSceneRunningDataResponse) SetStatusCode(v int32) *GetJMeterSceneRunningDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJMeterSceneRunningDataResponse) SetBody(v *GetJMeterSceneRunningDataResponseBody) *GetJMeterSceneRunningDataResponse {
	s.Body = v
	return s
}

func (s *GetJMeterSceneRunningDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
