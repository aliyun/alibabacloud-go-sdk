// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIProjectConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDIProjectConfigResponseBodyData) *ListDIProjectConfigResponseBody
	GetData() *ListDIProjectConfigResponseBodyData
	SetRequestId(v string) *ListDIProjectConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDIProjectConfigResponseBody
	GetSuccess() *bool
}

type ListDIProjectConfigResponseBody struct {
	// The information about the query.
	Data *ListDIProjectConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDIProjectConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDIProjectConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListDIProjectConfigResponseBody) GetData() *ListDIProjectConfigResponseBodyData {
	return s.Data
}

func (s *ListDIProjectConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDIProjectConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDIProjectConfigResponseBody) SetData(v *ListDIProjectConfigResponseBodyData) *ListDIProjectConfigResponseBody {
	s.Data = v
	return s
}

func (s *ListDIProjectConfigResponseBody) SetRequestId(v string) *ListDIProjectConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDIProjectConfigResponseBody) SetSuccess(v bool) *ListDIProjectConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ListDIProjectConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDIProjectConfigResponseBodyData struct {
	// The default global configuration of synchronization solutions. The value indicates the processing rules of different types of DDL messages. Example: {"RENAMECOLUMN":"WARNING","DROPTABLE":"WARNING","CREATETABLE":"WARNING","MODIFYCOLUMN":"WARNING","TRUNCATETABLE":"WARNING","DROPCOLUMN":"WARNING","ADDCOLUMN":"WARNING","RENAMETABLE":"WARNING"}
	//
	// Field description:
	//
	// 	- RENAMECOLUMN: renames a column.
	//
	// 	- DROPTABLE: deletes a table.
	//
	// 	- CREATETABLE: creates a table.
	//
	// 	- MODIFYCOLUMN: changes the data type of a column.
	//
	// 	- TRUNCATETABLE: clears a table.
	//
	// 	- DROPCOLUMN: deletes a column.
	//
	// 	- ADDCOLUMN: creates a column.
	//
	// 	- RENAMETABLE: renames a table.
	//
	// DataWorks processes a DDL message of a specific type based on the following rules:
	//
	// 	- WARNING: ignores the message and records an alert in real-time synchronization logs. The alert contains information about the situation that the message is ignored because of an execution error.
	//
	// 	- IGNORE: discards the message and does not send it to the destination.
	//
	// 	- CRITICAL: terminates the real-time synchronization node and sets the node status to Failed.
	//
	// 	- NORMAL: sends the message to the destination to process the message. Each destination processes DDL messages based on its own business logic. If DataWorks adopts the NORMAL policy, DataWorks only forwards DDL messages.
	//
	// example:
	//
	// {"RENAMECOLUMN":"WARNING","DROPTABLE":"WARNING","CREATETABLE":"WARNING","MODIFYCOLUMN":"WARNING","TRUNCATETABLE":"WARNING","DROPCOLUMN":"WARNING","ADDCOLUMN":"WARNING","RENAMETABLE":"WARNING"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s ListDIProjectConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDIProjectConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDIProjectConfigResponseBodyData) GetConfig() *string {
	return s.Config
}

func (s *ListDIProjectConfigResponseBodyData) SetConfig(v string) *ListDIProjectConfigResponseBodyData {
	s.Config = &v
	return s
}

func (s *ListDIProjectConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
