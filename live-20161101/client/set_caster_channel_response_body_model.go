// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCasterChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetCasterChannelResponseBody
	GetRequestId() *string
}

type SetCasterChannelResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCasterChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetCasterChannelResponseBody) GoString() string {
	return s.String()
}

func (s *SetCasterChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetCasterChannelResponseBody) SetRequestId(v string) *SetCasterChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetCasterChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
