// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportAdminsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ImportAdminsRequest
	GetInstanceId() *string
	SetRamIdList(v string) *ImportAdminsRequest
	GetRamIdList() *string
}

type ImportAdminsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["26972543893791****"]
	RamIdList *string `json:"RamIdList,omitempty" xml:"RamIdList,omitempty"`
}

func (s ImportAdminsRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportAdminsRequest) GoString() string {
	return s.String()
}

func (s *ImportAdminsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ImportAdminsRequest) GetRamIdList() *string {
	return s.RamIdList
}

func (s *ImportAdminsRequest) SetInstanceId(v string) *ImportAdminsRequest {
	s.InstanceId = &v
	return s
}

func (s *ImportAdminsRequest) SetRamIdList(v string) *ImportAdminsRequest {
	s.RamIdList = &v
	return s
}

func (s *ImportAdminsRequest) Validate() error {
	return dara.Validate(s)
}
