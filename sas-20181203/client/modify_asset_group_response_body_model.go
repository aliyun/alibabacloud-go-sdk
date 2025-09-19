// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAssetGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAssetGroupResponseBody
	GetRequestId() *string
}

type ModifyAssetGroupResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C2677612-7207-4AEB-BD48-8BA528F86777
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAssetGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAssetGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAssetGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAssetGroupResponseBody) SetRequestId(v string) *ModifyAssetGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAssetGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
