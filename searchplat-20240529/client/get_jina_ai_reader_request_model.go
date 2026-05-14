// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJinaAiReaderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUrl(v string) *GetJinaAiReaderRequest
	GetUrl() *string
}

type GetJinaAiReaderRequest struct {
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetJinaAiReaderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJinaAiReaderRequest) GoString() string {
	return s.String()
}

func (s *GetJinaAiReaderRequest) GetUrl() *string {
	return s.Url
}

func (s *GetJinaAiReaderRequest) SetUrl(v string) *GetJinaAiReaderRequest {
	s.Url = &v
	return s
}

func (s *GetJinaAiReaderRequest) Validate() error {
	return dara.Validate(s)
}
