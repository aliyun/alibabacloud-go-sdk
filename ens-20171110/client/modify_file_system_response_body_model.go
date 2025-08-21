// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFileSystemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyFileSystemResponseBody
	GetRequestId() *string
}

type ModifyFileSystemResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFileSystemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFileSystemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyFileSystemResponseBody) SetRequestId(v string) *ModifyFileSystemResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFileSystemResponseBody) Validate() error {
	return dara.Validate(s)
}
