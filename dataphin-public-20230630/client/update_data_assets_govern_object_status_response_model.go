// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataAssetsGovernObjectStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDataAssetsGovernObjectStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDataAssetsGovernObjectStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDataAssetsGovernObjectStatusResponseBody) *UpdateDataAssetsGovernObjectStatusResponse
	GetBody() *UpdateDataAssetsGovernObjectStatusResponseBody
}

type UpdateDataAssetsGovernObjectStatusResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataAssetsGovernObjectStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataAssetsGovernObjectStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAssetsGovernObjectStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataAssetsGovernObjectStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDataAssetsGovernObjectStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDataAssetsGovernObjectStatusResponse) GetBody() *UpdateDataAssetsGovernObjectStatusResponseBody {
	return s.Body
}

func (s *UpdateDataAssetsGovernObjectStatusResponse) SetHeaders(v map[string]*string) *UpdateDataAssetsGovernObjectStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataAssetsGovernObjectStatusResponse) SetStatusCode(v int32) *UpdateDataAssetsGovernObjectStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataAssetsGovernObjectStatusResponse) SetBody(v *UpdateDataAssetsGovernObjectStatusResponseBody) *UpdateDataAssetsGovernObjectStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateDataAssetsGovernObjectStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
