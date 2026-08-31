// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAssetsGovernObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataAssetsGovernObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataAssetsGovernObjectResponse
	GetStatusCode() *int32
	SetBody(v *GetDataAssetsGovernObjectResponseBody) *GetDataAssetsGovernObjectResponse
	GetBody() *GetDataAssetsGovernObjectResponseBody
}

type GetDataAssetsGovernObjectResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataAssetsGovernObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataAssetsGovernObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataAssetsGovernObjectResponse) GoString() string {
	return s.String()
}

func (s *GetDataAssetsGovernObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataAssetsGovernObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataAssetsGovernObjectResponse) GetBody() *GetDataAssetsGovernObjectResponseBody {
	return s.Body
}

func (s *GetDataAssetsGovernObjectResponse) SetHeaders(v map[string]*string) *GetDataAssetsGovernObjectResponse {
	s.Headers = v
	return s
}

func (s *GetDataAssetsGovernObjectResponse) SetStatusCode(v int32) *GetDataAssetsGovernObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataAssetsGovernObjectResponse) SetBody(v *GetDataAssetsGovernObjectResponseBody) *GetDataAssetsGovernObjectResponse {
	s.Body = v
	return s
}

func (s *GetDataAssetsGovernObjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
