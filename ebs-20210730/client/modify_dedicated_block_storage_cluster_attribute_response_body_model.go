// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedBlockStorageClusterAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDedicatedBlockStorageClusterAttributeResponseBody
	GetRequestId() *string
}

type ModifyDedicatedBlockStorageClusterAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 11B55F58-D3A4-4A9B-9596-342420D0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDedicatedBlockStorageClusterAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedBlockStorageClusterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedBlockStorageClusterAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDedicatedBlockStorageClusterAttributeResponseBody) SetRequestId(v string) *ModifyDedicatedBlockStorageClusterAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDedicatedBlockStorageClusterAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
