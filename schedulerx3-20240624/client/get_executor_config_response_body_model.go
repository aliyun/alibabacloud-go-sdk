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
	// The status code returned for the request.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The configuration data for the Executor.
	Data *GetExecutorConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// not support query script history, please upgrade engine version to 2.2.2+
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
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
	// The default global configuration for Data Integration tasks. This configuration specifies the default handling policies for different types of DDL messages. Example:
	//
	// `{"RENAMECOLUMN":"WARNING","DROPTABLE":"WARNING","CREATETABLE":"WARNING","MODIFYCOLUMN":"WARNING","TRUNCATETABLE":"WARNING","DROPCOLUMN":"WARNING","ADDCOLUMN":"WARNING","RENAMETABLE":"WARNING"}`
	//
	// The DDL message types are as follows:
	//
	// - RENAMECOLUMN: `RENAME COLUMN`
	//
	// - DROPTABLE: `DROP TABLE`
	//
	// - CREATETABLE: `CREATE TABLE`
	//
	// - MODIFYCOLUMN: `MODIFY COLUMN`
	//
	// - TRUNCATETABLE: `TRUNCATE TABLE`
	//
	// - DROPCOLUMN: `DROP COLUMN`
	//
	// - ADDCOLUMN: `ADD COLUMN`
	//
	// - RENAMETABLE: `RENAME TABLE`
	//
	// When DataWorks receives a DDL message, it applies one of the following handling policies:
	//
	// - WARNING: Discards the message and logs a warning in the Real-time Synchronization Task log.
	//
	// - IGNORE: Discards the message without sending it to the Destination Data Source.
	//
	// - CRITICAL: Causes the Real-time Synchronization Task to fail.
	//
	// - NORMAL: Forwards the message to the Destination Data Source. Because handling of DDL messages can vary by Destination Data Source, DataWorks only forwards the message.
	//
	// example:
	//
	// [{"cluster":"c2c619b5129e0400fa3df263b249622aa","namespace":"default","service":"xxljob-http-demo1-svc"}]
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The type of the Executor.
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
