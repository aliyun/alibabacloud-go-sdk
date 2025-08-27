// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChannelCorpCreateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChannelCorpCreateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChannelCorpCreateResponse
	GetStatusCode() *int32
	SetBody(v *ChannelCorpCreateResponseBody) *ChannelCorpCreateResponse
	GetBody() *ChannelCorpCreateResponseBody
}

type ChannelCorpCreateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChannelCorpCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChannelCorpCreateResponse) String() string {
	return dara.Prettify(s)
}

func (s ChannelCorpCreateResponse) GoString() string {
	return s.String()
}

func (s *ChannelCorpCreateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChannelCorpCreateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChannelCorpCreateResponse) GetBody() *ChannelCorpCreateResponseBody {
	return s.Body
}

func (s *ChannelCorpCreateResponse) SetHeaders(v map[string]*string) *ChannelCorpCreateResponse {
	s.Headers = v
	return s
}

func (s *ChannelCorpCreateResponse) SetStatusCode(v int32) *ChannelCorpCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *ChannelCorpCreateResponse) SetBody(v *ChannelCorpCreateResponseBody) *ChannelCorpCreateResponse {
	s.Body = v
	return s
}

func (s *ChannelCorpCreateResponse) Validate() error {
	return dara.Validate(s)
}
