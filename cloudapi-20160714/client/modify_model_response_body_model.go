// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyModelResponseBody
	GetRequestId() *string
}

type ModifyModelResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4173F95B-360C-460C-9F6C-4A960B904411
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyModelResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyModelResponseBody) SetRequestId(v string) *ModifyModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyModelResponseBody) Validate() error {
	return dara.Validate(s)
}
