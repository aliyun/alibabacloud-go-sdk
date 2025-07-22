// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutonomousNotifyEventContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAutonomousNotifyEventContentResponseBody
	GetCode() *string
	SetData(v string) *GetAutonomousNotifyEventContentResponseBody
	GetData() *string
	SetMessage(v string) *GetAutonomousNotifyEventContentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAutonomousNotifyEventContentResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetAutonomousNotifyEventContentResponseBody
	GetSuccess() *string
}

type GetAutonomousNotifyEventContentResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the notification events.
	//
	// example:
	//
	// {\\"taskId\\":\\"7e1ba595-0889-48ff-a6ff-010f54991d****\\",\\"taskType\\":\\"SQL_OPTIMIZE\\",\\"advisorId\\":\\"636dc5f34664dd56ff0****\\",\\"sqlId\\":\\"e2b1d6c1ee1bb29555a828b59f16****\\",\\"indexAdviceCount\\":1,\\"indexAdvices\\":[{\\"schemaName\\":\\"das\\",\\"tableName\\":\\"students\\",\\"indexName\\":\\"idx_name\\",\\"columns\\":[\\"name\\"],\\"unique\\":false,\\"ddlAddIndex\\":\\"ALTER TABLE `das`.`students` ADD INDEX `idx_name` (`name`)\\",\\"priority\\":0,\\"optimizeId\\":\\"96232794517277511\\"}],\\"tuningAdvices\\":[],\\"improvement\\":8127.25,\\"supportLevel\\":3,\\"priority\\":\\"HIGH\\"}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// >  If the request was successful, Successful is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAutonomousNotifyEventContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAutonomousNotifyEventContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutonomousNotifyEventContentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAutonomousNotifyEventContentResponseBody) GetData() *string {
	return s.Data
}

func (s *GetAutonomousNotifyEventContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAutonomousNotifyEventContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAutonomousNotifyEventContentResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetAutonomousNotifyEventContentResponseBody) SetCode(v string) *GetAutonomousNotifyEventContentResponseBody {
	s.Code = &v
	return s
}

func (s *GetAutonomousNotifyEventContentResponseBody) SetData(v string) *GetAutonomousNotifyEventContentResponseBody {
	s.Data = &v
	return s
}

func (s *GetAutonomousNotifyEventContentResponseBody) SetMessage(v string) *GetAutonomousNotifyEventContentResponseBody {
	s.Message = &v
	return s
}

func (s *GetAutonomousNotifyEventContentResponseBody) SetRequestId(v string) *GetAutonomousNotifyEventContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutonomousNotifyEventContentResponseBody) SetSuccess(v string) *GetAutonomousNotifyEventContentResponseBody {
	s.Success = &v
	return s
}

func (s *GetAutonomousNotifyEventContentResponseBody) Validate() error {
	return dara.Validate(s)
}
