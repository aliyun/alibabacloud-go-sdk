// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMCTableSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetColumns(v []*GetMCTableSchemaResponseBodyColumns) *GetMCTableSchemaResponseBody
	GetColumns() []*GetMCTableSchemaResponseBodyColumns
	SetPartitionColumns(v []*string) *GetMCTableSchemaResponseBody
	GetPartitionColumns() []*string
	SetRequestId(v string) *GetMCTableSchemaResponseBody
	GetRequestId() *string
}

type GetMCTableSchemaResponseBody struct {
	Columns          []*GetMCTableSchemaResponseBodyColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	PartitionColumns []*string                              `json:"PartitionColumns,omitempty" xml:"PartitionColumns,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMCTableSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMCTableSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetMCTableSchemaResponseBody) GetColumns() []*GetMCTableSchemaResponseBodyColumns {
	return s.Columns
}

func (s *GetMCTableSchemaResponseBody) GetPartitionColumns() []*string {
	return s.PartitionColumns
}

func (s *GetMCTableSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMCTableSchemaResponseBody) SetColumns(v []*GetMCTableSchemaResponseBodyColumns) *GetMCTableSchemaResponseBody {
	s.Columns = v
	return s
}

func (s *GetMCTableSchemaResponseBody) SetPartitionColumns(v []*string) *GetMCTableSchemaResponseBody {
	s.PartitionColumns = v
	return s
}

func (s *GetMCTableSchemaResponseBody) SetRequestId(v string) *GetMCTableSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMCTableSchemaResponseBody) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMCTableSchemaResponseBodyColumns struct {
	// example:
	//
	// column1
	Name    *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Preview []*string `json:"Preview,omitempty" xml:"Preview,omitempty" type:"Repeated"`
	// example:
	//
	// STRING
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetMCTableSchemaResponseBodyColumns) String() string {
	return dara.Prettify(s)
}

func (s GetMCTableSchemaResponseBodyColumns) GoString() string {
	return s.String()
}

func (s *GetMCTableSchemaResponseBodyColumns) GetName() *string {
	return s.Name
}

func (s *GetMCTableSchemaResponseBodyColumns) GetPreview() []*string {
	return s.Preview
}

func (s *GetMCTableSchemaResponseBodyColumns) GetType() *string {
	return s.Type
}

func (s *GetMCTableSchemaResponseBodyColumns) SetName(v string) *GetMCTableSchemaResponseBodyColumns {
	s.Name = &v
	return s
}

func (s *GetMCTableSchemaResponseBodyColumns) SetPreview(v []*string) *GetMCTableSchemaResponseBodyColumns {
	s.Preview = v
	return s
}

func (s *GetMCTableSchemaResponseBodyColumns) SetType(v string) *GetMCTableSchemaResponseBodyColumns {
	s.Type = &v
	return s
}

func (s *GetMCTableSchemaResponseBodyColumns) Validate() error {
	return dara.Validate(s)
}
