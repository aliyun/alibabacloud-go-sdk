// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetNotifyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetNotifyConfigResponseBody
	GetRequestId() *string
}

type SetNotifyConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 771A1414-27BF-53E6-AB73-EFCB*****ACF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetNotifyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetNotifyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetNotifyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetNotifyConfigResponseBody) SetRequestId(v string) *SetNotifyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetNotifyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
