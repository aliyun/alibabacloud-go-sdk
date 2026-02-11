// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRetcodeShareUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetRetcodeShareUrlResponseBody
	GetRequestId() *string
	SetUrl(v string) *GetRetcodeShareUrlResponseBody
	GetUrl() *string
}

type GetRetcodeShareUrlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetRetcodeShareUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRetcodeShareUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetRetcodeShareUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRetcodeShareUrlResponseBody) GetUrl() *string {
	return s.Url
}

func (s *GetRetcodeShareUrlResponseBody) SetRequestId(v string) *GetRetcodeShareUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRetcodeShareUrlResponseBody) SetUrl(v string) *GetRetcodeShareUrlResponseBody {
	s.Url = &v
	return s
}

func (s *GetRetcodeShareUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
