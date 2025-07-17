// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPartitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetPartitionRequest
	GetName() *string
	SetTableId(v string) *GetPartitionRequest
	GetTableId() *string
}

type GetPartitionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ds=20250101
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The table ID. For more details, refer to the response of the ListTables operation and [description of concepts related to metadata entities.](https://help.aliyun.com/document_detail/2880092.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-column:11075xxxx::test_project:test_schema:test_table
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
}

func (s GetPartitionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPartitionRequest) GoString() string {
	return s.String()
}

func (s *GetPartitionRequest) GetName() *string {
	return s.Name
}

func (s *GetPartitionRequest) GetTableId() *string {
	return s.TableId
}

func (s *GetPartitionRequest) SetName(v string) *GetPartitionRequest {
	s.Name = &v
	return s
}

func (s *GetPartitionRequest) SetTableId(v string) *GetPartitionRequest {
	s.TableId = &v
	return s
}

func (s *GetPartitionRequest) Validate() error {
	return dara.Validate(s)
}
