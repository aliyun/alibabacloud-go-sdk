// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachAssetGroupToInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachAssetGroupToInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachAssetGroupToInstanceResponse
	GetStatusCode() *int32
	SetBody(v *AttachAssetGroupToInstanceResponseBody) *AttachAssetGroupToInstanceResponse
	GetBody() *AttachAssetGroupToInstanceResponseBody
}

type AttachAssetGroupToInstanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachAssetGroupToInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachAssetGroupToInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachAssetGroupToInstanceResponse) GoString() string {
	return s.String()
}

func (s *AttachAssetGroupToInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachAssetGroupToInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachAssetGroupToInstanceResponse) GetBody() *AttachAssetGroupToInstanceResponseBody {
	return s.Body
}

func (s *AttachAssetGroupToInstanceResponse) SetHeaders(v map[string]*string) *AttachAssetGroupToInstanceResponse {
	s.Headers = v
	return s
}

func (s *AttachAssetGroupToInstanceResponse) SetStatusCode(v int32) *AttachAssetGroupToInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachAssetGroupToInstanceResponse) SetBody(v *AttachAssetGroupToInstanceResponseBody) *AttachAssetGroupToInstanceResponse {
	s.Body = v
	return s
}

func (s *AttachAssetGroupToInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
