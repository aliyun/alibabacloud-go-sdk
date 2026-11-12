// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecutorConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetExecutorConfigResponseBody
	GetCode() *int32
	SetData(v *GetExecutorConfigResponseBodyData) *GetExecutorConfigResponseBody
	GetData() *GetExecutorConfigResponseBodyData
	SetMessage(v string) *GetExecutorConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetExecutorConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetExecutorConfigResponseBody
	GetSuccess() *bool
}

type GetExecutorConfigResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query result.
	Data *GetExecutorConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// not support query script history, please upgrade engine version to 2.2.2+
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 39938688-0BAB-5AD8-BF02-F4910FAC7589
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetExecutorConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExecutorConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetExecutorConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetExecutorConfigResponseBody) GetData() *GetExecutorConfigResponseBodyData {
	return s.Data
}

func (s *GetExecutorConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetExecutorConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExecutorConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetExecutorConfigResponseBody) SetCode(v int32) *GetExecutorConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetExecutorConfigResponseBody) SetData(v *GetExecutorConfigResponseBodyData) *GetExecutorConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetExecutorConfigResponseBody) SetMessage(v string) *GetExecutorConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetExecutorConfigResponseBody) SetRequestId(v string) *GetExecutorConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExecutorConfigResponseBody) SetSuccess(v bool) *GetExecutorConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetExecutorConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetExecutorConfigResponseBodyData struct {
	// The default global configuration content for the Data Integration synchronization Solutions node, which specifies the default processing policy for different DDL message types. Example:
	//
	// {"RENAMECOLUMN":"WARNING","DROPTABLE":"WARNING","CREATETABLE":"WARNING","MODIFYCOLUMN":"WARNING","TRUNCATETABLE":"WARNING","DROPCOLUMN":"WARNING","ADDCOLUMN":"WARNING","RENAMETABLE":"WARNING"}
	//
	// DDL message types:
	//
	// - RENAMECOLUMN: Rename a column.
	//
	// - DROPTABLE: Delete a table.
	//
	// - CREATETABLE: Create a table.
	//
	// - MODIFYCOLUMN: Modify a column type.
	//
	// - TRUNCATETABLE: Truncate a table.
	//
	// - DROPCOLUMN: Delete a column.
	//
	// - ADDCOLUMN: Add a column.
	//
	// - RENAMETABLE: Rename a table.
	//
	// When DataWorks receives a DDL message of the corresponding type, the processing policy is as follows:
	//
	// - WARNING: Discard the message directly and record alerting information in the real-time synchronization log, indicating that the message was discarded due to an execute error.
	//
	// - IGNORE: Discard the message directly without sending it to the destination data source.
	//
	// - CRITICAL: Set the real-time synchronization node to the error state and stop the node.
	//
	// - NORMAL: Forward the message to the destination data source for processing. Because different destination data sources may have a different processing policy for DDL messages, DataWorks only forwards the message.
	//
	// example:
	//
	// [{"cluster":"c2c619b5129e0400fa3df263b249622aa","namespace":"default","service":"xxljob-http-demo1-svc"}]
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The executor type.
	//
	// example:
	//
	// k8s_service
	ExecutorType *string `json:"ExecutorType,omitempty" xml:"ExecutorType,omitempty"`
}

func (s GetExecutorConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetExecutorConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetExecutorConfigResponseBodyData) GetConfig() *string {
	return s.Config
}

func (s *GetExecutorConfigResponseBodyData) GetExecutorType() *string {
	return s.ExecutorType
}

func (s *GetExecutorConfigResponseBodyData) SetConfig(v string) *GetExecutorConfigResponseBodyData {
	s.Config = &v
	return s
}

func (s *GetExecutorConfigResponseBodyData) SetExecutorType(v string) *GetExecutorConfigResponseBodyData {
	s.ExecutorType = &v
	return s
}

func (s *GetExecutorConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
