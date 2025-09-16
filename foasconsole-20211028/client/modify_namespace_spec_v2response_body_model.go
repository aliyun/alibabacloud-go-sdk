// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNamespaceSpecV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyNamespaceSpecV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyNamespaceSpecV2ResponseBody
	GetSuccess() *bool
}

type ModifyNamespaceSpecV2ResponseBody struct {
	// example:
	//
	// 23A9C718-DDAB-1696-B025-18FBC830F7C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyNamespaceSpecV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNamespaceSpecV2ResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNamespaceSpecV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNamespaceSpecV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyNamespaceSpecV2ResponseBody) SetRequestId(v string) *ModifyNamespaceSpecV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNamespaceSpecV2ResponseBody) SetSuccess(v bool) *ModifyNamespaceSpecV2ResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyNamespaceSpecV2ResponseBody) Validate() error {
	return dara.Validate(s)
}
