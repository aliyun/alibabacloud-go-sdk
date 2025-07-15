// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachGroupPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachGroupPluginResponseBody
	GetRequestId() *string
}

type DetachGroupPluginResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 3707E6FA-749C-5352-B72A-ABE95D9DEA49
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachGroupPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachGroupPluginResponseBody) GoString() string {
	return s.String()
}

func (s *DetachGroupPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachGroupPluginResponseBody) SetRequestId(v string) *DetachGroupPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachGroupPluginResponseBody) Validate() error {
	return dara.Validate(s)
}
