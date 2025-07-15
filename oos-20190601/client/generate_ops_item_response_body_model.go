// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateOpsItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOpsItemIds(v []*string) *GenerateOpsItemResponseBody
	GetOpsItemIds() []*string
	SetRequestId(v string) *GenerateOpsItemResponseBody
	GetRequestId() *string
}

type GenerateOpsItemResponseBody struct {
	// The O\\&M item list.
	OpsItemIds []*string `json:"OpsItemIds,omitempty" xml:"OpsItemIds,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// DA4F08D4-DA54-5407-84B9-108C71D75B17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateOpsItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateOpsItemResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateOpsItemResponseBody) GetOpsItemIds() []*string {
	return s.OpsItemIds
}

func (s *GenerateOpsItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateOpsItemResponseBody) SetOpsItemIds(v []*string) *GenerateOpsItemResponseBody {
	s.OpsItemIds = v
	return s
}

func (s *GenerateOpsItemResponseBody) SetRequestId(v string) *GenerateOpsItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateOpsItemResponseBody) Validate() error {
	return dara.Validate(s)
}
