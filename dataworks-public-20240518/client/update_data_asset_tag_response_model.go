// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataAssetTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataAssetTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataAssetTagResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataAssetTagResponseBody) *UpdateDataAssetTagResponse
	GetBody() *UpdateDataAssetTagResponseBody
}

type UpdateDataAssetTagResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataAssetTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataAssetTagResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAssetTagResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataAssetTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataAssetTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataAssetTagResponse) GetBody() *UpdateDataAssetTagResponseBody {
	return s.Body
}

func (s *UpdateDataAssetTagResponse) SetHeaders(v map[string]*string) *UpdateDataAssetTagResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataAssetTagResponse) SetStatusCode(v int32) *UpdateDataAssetTagResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataAssetTagResponse) SetBody(v *UpdateDataAssetTagResponseBody) *UpdateDataAssetTagResponse {
	s.Body = v
	return s
}

func (s *UpdateDataAssetTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
