// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetUserConfigsResponseBody
	GetRequestId() *string
}

type SetUserConfigsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// dsjk****dfjksdf
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetUserConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetUserConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *SetUserConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetUserConfigsResponseBody) SetRequestId(v string) *SetUserConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetUserConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}
