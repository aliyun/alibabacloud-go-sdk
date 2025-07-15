// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportDoNotCallNumbersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilePath(v string) *ImportDoNotCallNumbersRequest
	GetFilePath() *string
	SetInstanceId(v string) *ImportDoNotCallNumbersRequest
	GetInstanceId() *string
	SetNumberList(v string) *ImportDoNotCallNumbersRequest
	GetNumberList() *string
	SetRemark(v string) *ImportDoNotCallNumbersRequest
	GetRemark() *string
}

type ImportDoNotCallNumbersRequest struct {
	// example:
	//
	// temp/blacklist/import/15772471154xxxx/ccc-test/20220315100340/blacklist.xlsx
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ["1900000****","1312121****"]
	NumberList *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ImportDoNotCallNumbersRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportDoNotCallNumbersRequest) GoString() string {
	return s.String()
}

func (s *ImportDoNotCallNumbersRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *ImportDoNotCallNumbersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ImportDoNotCallNumbersRequest) GetNumberList() *string {
	return s.NumberList
}

func (s *ImportDoNotCallNumbersRequest) GetRemark() *string {
	return s.Remark
}

func (s *ImportDoNotCallNumbersRequest) SetFilePath(v string) *ImportDoNotCallNumbersRequest {
	s.FilePath = &v
	return s
}

func (s *ImportDoNotCallNumbersRequest) SetInstanceId(v string) *ImportDoNotCallNumbersRequest {
	s.InstanceId = &v
	return s
}

func (s *ImportDoNotCallNumbersRequest) SetNumberList(v string) *ImportDoNotCallNumbersRequest {
	s.NumberList = &v
	return s
}

func (s *ImportDoNotCallNumbersRequest) SetRemark(v string) *ImportDoNotCallNumbersRequest {
	s.Remark = &v
	return s
}

func (s *ImportDoNotCallNumbersRequest) Validate() error {
	return dara.Validate(s)
}
