// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateAssetGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *CreateOrUpdateAssetGroupResponseBody
	GetGroupId() *int64
	SetRequestId(v string) *CreateOrUpdateAssetGroupResponseBody
	GetRequestId() *string
}

type CreateOrUpdateAssetGroupResponseBody struct {
	// The ID of the server group.
	//
	// example:
	//
	// 9935302
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// E70074C8-DFB4-44C5-96C7-909DD231D68A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrUpdateAssetGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateAssetGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAssetGroupResponseBody) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateOrUpdateAssetGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrUpdateAssetGroupResponseBody) SetGroupId(v int64) *CreateOrUpdateAssetGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateOrUpdateAssetGroupResponseBody) SetRequestId(v string) *CreateOrUpdateAssetGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrUpdateAssetGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
