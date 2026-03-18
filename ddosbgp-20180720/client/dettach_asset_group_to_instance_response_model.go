// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDettachAssetGroupToInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DettachAssetGroupToInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DettachAssetGroupToInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DettachAssetGroupToInstanceResponseBody) *DettachAssetGroupToInstanceResponse
	GetBody() *DettachAssetGroupToInstanceResponseBody
}

type DettachAssetGroupToInstanceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DettachAssetGroupToInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DettachAssetGroupToInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DettachAssetGroupToInstanceResponse) GoString() string {
	return s.String()
}

func (s *DettachAssetGroupToInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DettachAssetGroupToInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DettachAssetGroupToInstanceResponse) GetBody() *DettachAssetGroupToInstanceResponseBody {
	return s.Body
}

func (s *DettachAssetGroupToInstanceResponse) SetHeaders(v map[string]*string) *DettachAssetGroupToInstanceResponse {
	s.Headers = v
	return s
}

func (s *DettachAssetGroupToInstanceResponse) SetStatusCode(v int32) *DettachAssetGroupToInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DettachAssetGroupToInstanceResponse) SetBody(v *DettachAssetGroupToInstanceResponseBody) *DettachAssetGroupToInstanceResponse {
	s.Body = v
	return s
}

func (s *DettachAssetGroupToInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
