// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPtsSceneRunningDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPtsSceneRunningDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPtsSceneRunningDataResponse
	GetStatusCode() *int32
	SetBody(v *GetPtsSceneRunningDataResponseBody) *GetPtsSceneRunningDataResponse
	GetBody() *GetPtsSceneRunningDataResponseBody
}

type GetPtsSceneRunningDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPtsSceneRunningDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPtsSceneRunningDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneRunningDataResponse) GoString() string {
	return s.String()
}

func (s *GetPtsSceneRunningDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPtsSceneRunningDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPtsSceneRunningDataResponse) GetBody() *GetPtsSceneRunningDataResponseBody {
	return s.Body
}

func (s *GetPtsSceneRunningDataResponse) SetHeaders(v map[string]*string) *GetPtsSceneRunningDataResponse {
	s.Headers = v
	return s
}

func (s *GetPtsSceneRunningDataResponse) SetStatusCode(v int32) *GetPtsSceneRunningDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPtsSceneRunningDataResponse) SetBody(v *GetPtsSceneRunningDataResponseBody) *GetPtsSceneRunningDataResponse {
	s.Body = v
	return s
}

func (s *GetPtsSceneRunningDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
