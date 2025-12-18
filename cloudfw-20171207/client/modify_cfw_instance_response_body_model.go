// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCfwInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCfwInstanceResponseBody
	GetRequestId() *string
}

type ModifyCfwInstanceResponseBody struct {
	// example:
	//
	// F0F82705-CFC7-5F83-86C8-A063892F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCfwInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCfwInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCfwInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCfwInstanceResponseBody) SetRequestId(v string) *ModifyCfwInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCfwInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
