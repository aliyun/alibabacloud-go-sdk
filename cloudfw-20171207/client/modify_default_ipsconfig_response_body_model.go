// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefaultIPSConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDefaultIPSConfigResponseBody
	GetRequestId() *string
}

type ModifyDefaultIPSConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 65885B52-00EC-5728-96******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefaultIPSConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefaultIPSConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefaultIPSConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDefaultIPSConfigResponseBody) SetRequestId(v string) *ModifyDefaultIPSConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDefaultIPSConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
