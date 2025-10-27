// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudAssetCriteriaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCloudAssetCriteriaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCloudAssetCriteriaResponse
	GetStatusCode() *int32
	SetBody(v *GetCloudAssetCriteriaResponseBody) *GetCloudAssetCriteriaResponse
	GetBody() *GetCloudAssetCriteriaResponseBody
}

type GetCloudAssetCriteriaResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCloudAssetCriteriaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCloudAssetCriteriaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAssetCriteriaResponse) GoString() string {
	return s.String()
}

func (s *GetCloudAssetCriteriaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCloudAssetCriteriaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCloudAssetCriteriaResponse) GetBody() *GetCloudAssetCriteriaResponseBody {
	return s.Body
}

func (s *GetCloudAssetCriteriaResponse) SetHeaders(v map[string]*string) *GetCloudAssetCriteriaResponse {
	s.Headers = v
	return s
}

func (s *GetCloudAssetCriteriaResponse) SetStatusCode(v int32) *GetCloudAssetCriteriaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCloudAssetCriteriaResponse) SetBody(v *GetCloudAssetCriteriaResponseBody) *GetCloudAssetCriteriaResponse {
	s.Body = v
	return s
}

func (s *GetCloudAssetCriteriaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
