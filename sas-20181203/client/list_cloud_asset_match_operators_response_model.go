// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAssetMatchOperatorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudAssetMatchOperatorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudAssetMatchOperatorsResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudAssetMatchOperatorsResponseBody) *ListCloudAssetMatchOperatorsResponse
	GetBody() *ListCloudAssetMatchOperatorsResponseBody
}

type ListCloudAssetMatchOperatorsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudAssetMatchOperatorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudAssetMatchOperatorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAssetMatchOperatorsResponse) GoString() string {
	return s.String()
}

func (s *ListCloudAssetMatchOperatorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudAssetMatchOperatorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudAssetMatchOperatorsResponse) GetBody() *ListCloudAssetMatchOperatorsResponseBody {
	return s.Body
}

func (s *ListCloudAssetMatchOperatorsResponse) SetHeaders(v map[string]*string) *ListCloudAssetMatchOperatorsResponse {
	s.Headers = v
	return s
}

func (s *ListCloudAssetMatchOperatorsResponse) SetStatusCode(v int32) *ListCloudAssetMatchOperatorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudAssetMatchOperatorsResponse) SetBody(v *ListCloudAssetMatchOperatorsResponseBody) *ListCloudAssetMatchOperatorsResponse {
	s.Body = v
	return s
}

func (s *ListCloudAssetMatchOperatorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
