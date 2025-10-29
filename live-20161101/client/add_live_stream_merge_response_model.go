// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveStreamMergeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLiveStreamMergeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLiveStreamMergeResponse
	GetStatusCode() *int32
	SetBody(v *AddLiveStreamMergeResponseBody) *AddLiveStreamMergeResponse
	GetBody() *AddLiveStreamMergeResponseBody
}

type AddLiveStreamMergeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveStreamMergeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveStreamMergeResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLiveStreamMergeResponse) GoString() string {
	return s.String()
}

func (s *AddLiveStreamMergeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLiveStreamMergeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLiveStreamMergeResponse) GetBody() *AddLiveStreamMergeResponseBody {
	return s.Body
}

func (s *AddLiveStreamMergeResponse) SetHeaders(v map[string]*string) *AddLiveStreamMergeResponse {
	s.Headers = v
	return s
}

func (s *AddLiveStreamMergeResponse) SetStatusCode(v int32) *AddLiveStreamMergeResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveStreamMergeResponse) SetBody(v *AddLiveStreamMergeResponseBody) *AddLiveStreamMergeResponse {
	s.Body = v
	return s
}

func (s *AddLiveStreamMergeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
