// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAssetTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataAssetTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataAssetTagResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataAssetTagResponseBody) *DeleteDataAssetTagResponse
	GetBody() *DeleteDataAssetTagResponseBody
}

type DeleteDataAssetTagResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataAssetTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataAssetTagResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAssetTagResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataAssetTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataAssetTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataAssetTagResponse) GetBody() *DeleteDataAssetTagResponseBody {
	return s.Body
}

func (s *DeleteDataAssetTagResponse) SetHeaders(v map[string]*string) *DeleteDataAssetTagResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataAssetTagResponse) SetStatusCode(v int32) *DeleteDataAssetTagResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataAssetTagResponse) SetBody(v *DeleteDataAssetTagResponseBody) *DeleteDataAssetTagResponse {
	s.Body = v
	return s
}

func (s *DeleteDataAssetTagResponse) Validate() error {
	return dara.Validate(s)
}
