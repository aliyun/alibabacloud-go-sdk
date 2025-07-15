// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatasetItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDatasetItemResponseBody
	GetRequestId() *string
}

type ModifyDatasetItemResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F7DE77BC-0F7D-5A18-B494-BD2C********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDatasetItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatasetItemResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDatasetItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDatasetItemResponseBody) SetRequestId(v string) *ModifyDatasetItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDatasetItemResponseBody) Validate() error {
	return dara.Validate(s)
}
