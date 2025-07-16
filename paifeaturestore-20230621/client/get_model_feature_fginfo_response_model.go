// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelFeatureFGInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModelFeatureFGInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModelFeatureFGInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetModelFeatureFGInfoResponseBody) *GetModelFeatureFGInfoResponse
	GetBody() *GetModelFeatureFGInfoResponseBody
}

type GetModelFeatureFGInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelFeatureFGInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelFeatureFGInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModelFeatureFGInfoResponse) GoString() string {
	return s.String()
}

func (s *GetModelFeatureFGInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModelFeatureFGInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModelFeatureFGInfoResponse) GetBody() *GetModelFeatureFGInfoResponseBody {
	return s.Body
}

func (s *GetModelFeatureFGInfoResponse) SetHeaders(v map[string]*string) *GetModelFeatureFGInfoResponse {
	s.Headers = v
	return s
}

func (s *GetModelFeatureFGInfoResponse) SetStatusCode(v int32) *GetModelFeatureFGInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelFeatureFGInfoResponse) SetBody(v *GetModelFeatureFGInfoResponseBody) *GetModelFeatureFGInfoResponse {
	s.Body = v
	return s
}

func (s *GetModelFeatureFGInfoResponse) Validate() error {
	return dara.Validate(s)
}
