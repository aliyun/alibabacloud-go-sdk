// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserHdfsInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddUserHdfsInfoResponseBody
	GetRequestId() *string
}

type AddUserHdfsInfoResponseBody struct {
	// example:
	//
	// FB0B7918-198C-46A8-AB9B-FE15403B1F0A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserHdfsInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUserHdfsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserHdfsInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUserHdfsInfoResponseBody) SetRequestId(v string) *AddUserHdfsInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserHdfsInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
