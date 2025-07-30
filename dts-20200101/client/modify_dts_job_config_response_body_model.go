// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDtsJobConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDtsJobConfigResponseBody
	GetRequestId() *string
}

type ModifyDtsJobConfigResponseBody struct {
	// request ID
	//
	// example:
	//
	// 068FA72F-4800-4A54-90BB-94806068****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDtsJobConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDtsJobConfigResponseBody) SetRequestId(v string) *ModifyDtsJobConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDtsJobConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
