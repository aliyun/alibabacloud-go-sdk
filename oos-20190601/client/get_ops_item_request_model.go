// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpsItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpsItemId(v string) *GetOpsItemRequest
	GetOpsItemId() *string
	SetRegionId(v string) *GetOpsItemRequest
	GetRegionId() *string
}

type GetOpsItemRequest struct {
	// The O\\&M item ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// oi-d52b08695e2b46ae8413
	OpsItemId *string `json:"OpsItemId,omitempty" xml:"OpsItemId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetOpsItemRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOpsItemRequest) GoString() string {
	return s.String()
}

func (s *GetOpsItemRequest) GetOpsItemId() *string {
	return s.OpsItemId
}

func (s *GetOpsItemRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetOpsItemRequest) SetOpsItemId(v string) *GetOpsItemRequest {
	s.OpsItemId = &v
	return s
}

func (s *GetOpsItemRequest) SetRegionId(v string) *GetOpsItemRequest {
	s.RegionId = &v
	return s
}

func (s *GetOpsItemRequest) Validate() error {
	return dara.Validate(s)
}
