// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamMergeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveStreamMergeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveStreamMergeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveStreamMergeResponseBody) *DeleteLiveStreamMergeResponse
	GetBody() *DeleteLiveStreamMergeResponseBody
}

type DeleteLiveStreamMergeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveStreamMergeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveStreamMergeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamMergeResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamMergeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveStreamMergeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveStreamMergeResponse) GetBody() *DeleteLiveStreamMergeResponseBody {
	return s.Body
}

func (s *DeleteLiveStreamMergeResponse) SetHeaders(v map[string]*string) *DeleteLiveStreamMergeResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveStreamMergeResponse) SetStatusCode(v int32) *DeleteLiveStreamMergeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveStreamMergeResponse) SetBody(v *DeleteLiveStreamMergeResponseBody) *DeleteLiveStreamMergeResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveStreamMergeResponse) Validate() error {
	return dara.Validate(s)
}
