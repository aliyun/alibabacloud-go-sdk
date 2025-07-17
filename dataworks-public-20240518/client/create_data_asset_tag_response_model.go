// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAssetTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataAssetTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataAssetTagResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataAssetTagResponseBody) *CreateDataAssetTagResponse
	GetBody() *CreateDataAssetTagResponseBody
}

type CreateDataAssetTagResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataAssetTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataAssetTagResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAssetTagResponse) GoString() string {
	return s.String()
}

func (s *CreateDataAssetTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataAssetTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataAssetTagResponse) GetBody() *CreateDataAssetTagResponseBody {
	return s.Body
}

func (s *CreateDataAssetTagResponse) SetHeaders(v map[string]*string) *CreateDataAssetTagResponse {
	s.Headers = v
	return s
}

func (s *CreateDataAssetTagResponse) SetStatusCode(v int32) *CreateDataAssetTagResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataAssetTagResponse) SetBody(v *CreateDataAssetTagResponseBody) *CreateDataAssetTagResponse {
	s.Body = v
	return s
}

func (s *CreateDataAssetTagResponse) Validate() error {
	return dara.Validate(s)
}
