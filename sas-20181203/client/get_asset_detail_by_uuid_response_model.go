// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetDetailByUuidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAssetDetailByUuidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAssetDetailByUuidResponse
	GetStatusCode() *int32
	SetBody(v *GetAssetDetailByUuidResponseBody) *GetAssetDetailByUuidResponse
	GetBody() *GetAssetDetailByUuidResponseBody
}

type GetAssetDetailByUuidResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAssetDetailByUuidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAssetDetailByUuidResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAssetDetailByUuidResponse) GoString() string {
	return s.String()
}

func (s *GetAssetDetailByUuidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAssetDetailByUuidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAssetDetailByUuidResponse) GetBody() *GetAssetDetailByUuidResponseBody {
	return s.Body
}

func (s *GetAssetDetailByUuidResponse) SetHeaders(v map[string]*string) *GetAssetDetailByUuidResponse {
	s.Headers = v
	return s
}

func (s *GetAssetDetailByUuidResponse) SetStatusCode(v int32) *GetAssetDetailByUuidResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAssetDetailByUuidResponse) SetBody(v *GetAssetDetailByUuidResponseBody) *GetAssetDetailByUuidResponse {
	s.Body = v
	return s
}

func (s *GetAssetDetailByUuidResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
