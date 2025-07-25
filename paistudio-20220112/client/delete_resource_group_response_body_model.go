// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteResourceGroupResponseBody
	GetRequestId() *string
	SetResourceGroupID(v string) *DeleteResourceGroupResponseBody
	GetResourceGroupID() *string
}

type DeleteResourceGroupResponseBody struct {
	// example:
	//
	// F082BD0D-21E1-5F9B-81A0-AB07485B03CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rgvl9d6utwcscukh
	ResourceGroupID *string `json:"ResourceGroupID,omitempty" xml:"ResourceGroupID,omitempty"`
}

func (s DeleteResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceGroupResponseBody) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *DeleteResourceGroupResponseBody) SetRequestId(v string) *DeleteResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceGroupResponseBody) SetResourceGroupID(v string) *DeleteResourceGroupResponseBody {
	s.ResourceGroupID = &v
	return s
}

func (s *DeleteResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
